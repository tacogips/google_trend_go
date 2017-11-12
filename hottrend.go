package google_trend_go

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	hottrendSearchURL = "https://trends.google.com/trends/hottrends/hotItems"
)

// NewHotTrend returns new hot_trend api
func NewHotTrend(config HotTrendConfig) *HotTrend {
	return &HotTrend{
		config: config,
	}
}

type TrendRelatedSearch struct {
	Query string `json:"query"`
}

type HotTrendResponse struct {
	DataUpdateTimeUnitTime float64       `json:"dataUpdateTime"`
	TrendsByDate           []TrendByDate `json:"trendsByDateList"`
}

func (resp HotTrendResponse) DataUpdateTime() time.Time {
	return time.Unix(int64(resp.DataUpdateTimeUnitTime), 0)
}

type TrendByDate struct {
	Date   string  `json:"date"`
	Trends []Trend `json:"trendsList"`
}

type Trend struct {
	Date                    string               `json:"date"` // format: "20060102"
	Title                   string               `json:"title"`
	TitleLinkURL            string               `json:"titleLinkUrl"`
	HotnessLevel            int                  `json:"hotnessLevel"`
	FormattedTraffic        string               `json:"formattedTraffic"`
	StartTime               float64              `json:"startTime"`
	ImageURL                string               `json:"imgUrl"`
	ImageSource             string               `json:"imgSource"`
	Safe                    bool                 `json:"safe"`
	NewsArticle             []TrendArticle       `json:"newsArticlesList"`
	RelatedSearchs          []TrendRelatedSearch `json:"relatedSearchesList"`
	TrafficBucketLowerBound int                  `json:"trafficBucketLowerBound"`
}

type TrendArticle struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
	Source  string `json:"source"`
}

type HotTrend struct {
	config HotTrendConfig
}

type HotTrendConfig struct {
	Client        *http.Client
	RetryNum      int
	RetryInterval time.Duration
}

var DefaultHotTrendConfig = HotTrendConfig{
	Client:        &http.Client{},
	RetryNum:      10,
	RetryInterval: 1000 * time.Millisecond,
}

// Fetch return latest hot trend.
func (hottrend *HotTrend) Fetch(geo Geo) (*HotTrendResponse, error) {
	params := url.Values{}
	params.Add("ajax", "1")
	params.Add("pn", geoPtMap.getOrDefualt(geo))
	params.Add("htd", "") //TODO pagination?
	params.Add("htv", "l")

	tryNum := 0
	for {
		response, err := postForm(hottrend.config.Client, hottrendSearchURL, params)
		tryNum += 1

		if err != nil {
			if tryNum <= hottrend.config.RetryNum {
				interval := hottrend.config.RetryInterval
				//sleep at least 1 sec
				if interval < 1000*time.Millisecond {
					interval = 1000 * time.Millisecond
				}
				time.Sleep(hottrend.config.RetryInterval)
				continue
			} else {
				return nil, err
			}
		}

		trendResp := new(HotTrendResponse)
		err = json.Unmarshal(response, trendResp)
		if err != nil {
			return nil, err
		}

		return trendResp, nil
	}

}
