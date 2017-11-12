package go_trend_slack_bot

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/nlopes/slack"

	googletrend "github.com/tacogips/google_trend_go"
)

func PushHotTrends(httpClient *http.Client, apiKey string, channel string) error {
	slackApi := slack.New(apiKey)

	params := slack.PostMessageParameters{}

	config := googletrend.HotTrendConfig{
		Client:        httpClient,
		RetryNum:      10,
		RetryInterval: 1000 * time.Millisecond,
	}

	hottrend := googletrend.NewHotTrend(config)
	resp, err := hottrend.Fetch(googletrend.GeoJP)
	if err != nil {
		return err
	}
	attachments := []slack.Attachment{}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	updatedTime := resp.DataUpdateTime().In(jst).Format(time.RFC3339)

	// sort by hotness
	var trends []googletrend.Trend
	for _, eachDt := range resp.TrendsByDate {
		trends = append(trends, eachDt.Trends...)
	}
	sort.SliceStable(trends, func(i, j int) bool {
		return trends[i].HotnessLevel > trends[j].HotnessLevel
	})

	for _, trend := range trends {
		related := []slack.AttachmentField{}
		for _, rel := range trend.RelatedSearchs {
			related = append(related,
				slack.AttachmentField{
					Title: "query",
					Value: rel.Query,
					Short: true,
				})
		}

		articles := []slack.AttachmentField{}
		for _, news := range trend.NewsArticle {
			articles = append(articles,
				slack.AttachmentField{
					Title: "news",
					Value: fmt.Sprintf("%s <%s>", news.Title, news.Link),
					Short: true,
				})
		}

		var fields []slack.AttachmentField
		fields = append(fields,
			slack.AttachmentField{
				Title: "hotness",
				Value: fmt.Sprintf("%d", trend.HotnessLevel),

				Short: true,
			})
		fields = append(fields,
			slack.AttachmentField{
				Title: "traffic",
				Value: trend.FormattedTraffic,
				Short: true,
			})

		fields = append(fields, related...)
		fields = append(fields, articles...)

		color := "#B5B691" // glay
		if trend.HotnessLevel >= 5 {
			color = "#FF5733" //red
		} else if trend.HotnessLevel >= 3 {
			color = "#FFAC33" //orange
		} else if trend.HotnessLevel >= 2 {
			color = "#F6FF33" // yellow
		}

		imgURL := fmt.Sprintf("https:%s", trend.ImageURL)
		attachment := slack.Attachment{
			Title:    fmt.Sprintf("<%s> - %s", trend.Title, updatedTime),
			ThumbURL: imgURL,
			Fields:   fields,
			Color:    color,
		}
		attachments = append(attachments, attachment)
	}

	params.Attachments = attachments
	_, _, err = slackApi.PostMessage(channel, fmt.Sprintf("HotTrend "), params)
	if err != nil {
		return err
	}

	return nil
}
