package go_trend_slack_bot

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nlopes/slack"
	googletrend "github.com/tacogips/google_trend_go"
)

func PushStories(httpClient *http.Client, apiKey string, channel string) error {
	slackApi := slack.New(apiKey)

	params := slack.PostMessageParameters{}

	config := googletrend.StoriesConfig{
		Client:        httpClient,
		RetryNum:      10,
		RetryInterval: 1000 * time.Millisecond,
	}

	stories := googletrend.NewStories(config)

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	resp, err := stories.Fetch(googletrend.LangJA, googletrend.GeoJP, jst, 0, 10)
	if err != nil {
		return err
	}
	attachments := []slack.Attachment{}

	for idx, story := range resp.TrendStories {
		fields := []slack.AttachmentField{}
		for _, article := range story.StoryArticles {
			fields = append(fields,
				slack.AttachmentField{
					Title: "article",
					Value: fmt.Sprintf("%s <%s>", article.ArticleTitle, article.URL),
					Short: true,
				})
		}

		attachment := slack.Attachment{
			Title:    fmt.Sprintf("%d: <%s>", idx+1, story.Title),
			ThumbURL: story.StoryImage.ImageURL,
			Fields:   fields,
		}
		attachments = append(attachments, attachment)
	}

	params.Attachments = attachments
	_, _, err = slackApi.PostMessage(channel, fmt.Sprintf("Stories"), params)
	if err != nil {
		return err
	}

	return nil
}
