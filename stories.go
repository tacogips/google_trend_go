package google_trend_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

const (
	latestStoriesURL = "https://trends.google.com/trends/api/stories/latest"
	storySummaryURL  = "https://trends.google.com/trends/api/stories/summary"
)

func NewStories(config StoriesConfig) *Stories {
	return &Stories{
		config: config,
	}
}

type LatestStorieResponse struct {
	Date             string   `json:"date"`
	FeaturedStoryIDs []string `json:"featuredStoryIds"`
	TrendStoryIDs    []string `json:"trendingStoryIds"`
}

type StoriesSummaryResponse struct {
	Date            string         `json:"date"`
	FeaturedStories []StorySummary `json:"featuredStories"`
	TrendStories    []StorySummary `json:"trendingStories"`
}

func (s *StoriesSummaryResponse) Add(ss *StoriesSummaryResponse) {
	s.FeaturedStories = append(s.FeaturedStories, ss.FeaturedStories...)
	s.TrendStories = append(s.TrendStories, ss.TrendStories...)
}

type StorySummary struct {
	ID            string         `json:"id"`
	Title         string         `json:"title"`
	EntityNames   []string       `json:"entityNames"`
	StoryImage    StoryImage     `json:"image"`
	StoryArticles []StoryArticle `json:"articles"`
	IDsForDedup   []string       `json:"idsForDedup"`
}

type StoryImage struct {
	NewURL   string `json:"newsUrl"`
	Source   string `json:"source"`
	ImageURL string `json:"imgUrl"`
}

type StoryArticle struct {
	ArticleTitle string `json:"articleTitle"`
	URL          string `json:"url"`
	Source       string `json:"source"`
	Time         string `json:"time"`
}

type Stories struct {
	config StoriesConfig
}

type StoriesConfig struct {
	Client        *http.Client
	RetryNum      int
	RetryInterval time.Duration
}

var DefaultStoriesConfig = StoriesConfig{
	Client:        &http.Client{},
	RetryNum:      5,
	RetryInterval: 500 * time.Millisecond,
}

// FetchSummaryIDs
func (stories *Stories) FetchSummaryIDs(lang Lang, geo Geo, loc *time.Location) (*LatestStorieResponse, error) {
	geoCode, ok := storiesGeo[geo]
	if !ok {
		return nil, errors.New("no geo for stories")
	}

	timezone := toStorieAPITimeZone(loc)

	params := url.Values{}
	params.Add("hl", lang.String())
	params.Add("tz", strconv.Itoa(timezone))
	params.Add("cat", "all")
	params.Add("fi", "15") //TODO(tacogips) ?
	params.Add("fs", "15") //TODO(tacogips) ?
	params.Add("geo", geoCode)
	params.Add("ri", "300") //TODO(tacogips) ?
	params.Add("rs", "15")  //TODO(tacogips) ?
	params.Add("sort", "0") //TODO(tacogips) ?

	tryNum := 0

	for {
		response, err := get(stories.config.Client, latestStoriesURL, params)
		tryNum += 1

		if err != nil {
			if tryNum <= stories.config.RetryNum {
				time.Sleep(stories.config.RetryInterval)
				continue
			} else {
				return nil, err
			}
		}

		// sanitize
		response, err = sanitizeJS(response)
		if err != nil {
			return nil, err
		}

		latesStories := new(LatestStorieResponse)
		err = json.Unmarshal(response, latesStories)
		if err != nil {
			return nil, err
		}

		return latesStories, nil
	}
}

//Fetch return latest stories language and geolocation and location.
//The stories contains feature stories and trend stories. XXXstoryNum Parameterr stands for return stories number.
func (stories *Stories) Fetch(lang Lang, geo Geo, loc *time.Location, fetchFeatureStoryNum int, fetchTrendStoryNum int) (*StoriesSummaryResponse, error) {

	latesStories, err := stories.FetchSummaryIDs(lang, geo, loc)
	if err != nil {
		return nil, err
	}

	allSummaries := new(StoriesSummaryResponse)

	//TODO(tacogips) ugly
	{
		ids := latesStories.FeaturedStoryIDs
		if fetchFeatureStoryNum < len(ids) {
			ids = ids[0:fetchFeatureStoryNum]
		}
		idsChunk := strChunk(ids)
		for _, ids := range idsChunk {
			summaries, err := stories.FetchSummaryByIDs(lang, loc, ids)
			if err != nil {
				return nil, err
			}
			allSummaries.Add(summaries)
		}
	}

	//TODO(tacogips) ugly
	{
		ids := latesStories.TrendStoryIDs
		if fetchTrendStoryNum < len(ids) {
			ids = ids[0:fetchTrendStoryNum]
		}
		idsChunk := strChunk(ids)

		for _, ids := range idsChunk {
			summaries, err := stories.FetchSummaryByIDs(lang, loc, ids)
			if err != nil {
				return nil, err
			}
			allSummaries.Add(summaries)
		}
	}

	return allSummaries, nil
}

func sanitizeJS(s []byte) ([]byte, error) {
	pattern := regexp.MustCompile("{")
	// sanitize
	tobeSanitizeIdex := pattern.FindIndex(s)
	if len(tobeSanitizeIdex) != 2 || tobeSanitizeIdex[1] < 1 {
		return s, fmt.Errorf("invalid story response %s", string(s))
	}
	return s[tobeSanitizeIdex[1]-1:], nil
}

func toStorieAPITimeZone(loc *time.Location) int {
	_, offset := time.Now().In(loc).Local().Zone()
	return -1 * offset / 3600 * 60 // in minutes
}

func (stories *Stories) FetchSummaryByIDs(lang Lang, loc *time.Location, ids []string) (*StoriesSummaryResponse, error) {

	timezone := toStorieAPITimeZone(loc)
	params := url.Values{}
	params.Add("hl", lang.String())
	params.Add("tz", strconv.Itoa(timezone))

	for _, id := range ids {
		params.Add("id", id)
	}

	tryNum := 0
	for {
		response, err := get(stories.config.Client, storySummaryURL, params)
		tryNum += 1

		if err != nil {
			if tryNum <= stories.config.RetryNum {
				time.Sleep(stories.config.RetryInterval)
				continue
			} else {
				return nil, err
			}
		}

		// sanitize
		response, err = sanitizeJS(response)
		if err != nil {
			return nil, err
		}

		storiesSummary := new(StoriesSummaryResponse)
		err = json.Unmarshal(response, storiesSummary)
		if err != nil {
			return nil, err
		}

		return storiesSummary, nil
	}

}

var storiesGeo = map[Geo]string{
	GeoIE: "IE",
	GeoUS: "US",
	GeoAR: "AR",
	GeoGB: "GB",
	GeoIT: "IT",
	GeoIN: "IN",
	GeoAU: "AU",
	GeoAT: "AT",
	GeoNL: "NL",
	GeoCA: "CA",
	GeoCO: "CO",
	GeoCH: "CH",
	GeoSE: "SE",
	GeoCL: "CL",
	GeoDE: "DE",
	GeoTR: "TR",
	GeoNZ: "NZ",
	GeoNO: "NO",
	GeoPH: "PH",
	GeoBR: "BR",
	GeoFR: "FR",
	GeoVN: "VN",
	GeoPE: "PE",
	GeoBE: "BE",
	GeoPL: "PL",
	GeoPT: "PT",
	GeoMY: "MY",
	GeoMX: "MX",
	GeoRU: "RU",
	GeoJP: "JP",
}

func strChunk(ids []string) [][]string {
	// TODO(tacogips) for some reason "/api/stories/summary" returns respone less than id num in query string, so set chunk size small val
	chunkSize := 3

	q := len(ids) / chunkSize
	r := len(ids) % chunkSize
	var size = q
	if r != 0 {
		size += 1
	}
	result := make([][]string, size)
	for i := 0; i < size; i++ {
		tailidx := i*chunkSize + chunkSize
		if tailidx > len(ids) {
			tailidx = i*chunkSize + r
		}

		result[i] = ids[i*chunkSize : tailidx]
	}

	return result
}
