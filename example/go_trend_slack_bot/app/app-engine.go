package app

import (
	"net/http"
	"os"

	"github.com/nlopes/slack"
	"github.com/tacogips/google_trend_go/example/go_trend_slack_bot"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/job/hot_trends", regeisterTrendsHandler)
	http.HandleFunc("/job/stories", regeisterStories)
}

func regeisterTrendsHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	if r.Method == "GET" {
		task := taskqueue.NewPOSTTask("/job/hot_trends", nil)
		if _, err := taskqueue.Add(c, task, "default"); err != nil {
			log.Errorf(c, "%s", err.Error())
			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	} else if r.Method == "POST" {

		slackToken := os.Getenv("SLACK_TOKEN")
		if slackToken == "" {
			log.Errorf(c, "%s", "no slack api key")
			w.Write([]byte("error"))
			return
		}

		hotTrendChannel := os.Getenv("HOT_TREND_CHANNEL")
		if hotTrendChannel == "" {
			log.Errorf(c, "%s", "no hot trend channel")
			w.Write([]byte("error"))
			return
		}

		httpClient := urlfetch.Client(c)
		slack.SetHTTPClient(httpClient)

		if err := go_trend_slack_bot.PushHotTrends(httpClient, slackToken, hotTrendChannel); err != nil {
			log.Errorf(c, "%s", err.Error())
			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	}
}

func regeisterStories(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	if r.Method == "GET" {
		task := taskqueue.NewPOSTTask("/job/stories", nil)
		if _, err := taskqueue.Add(c, task, "default"); err != nil {
			log.Errorf(c, "%s", err.Error())
			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	} else if r.Method == "POST" {

		slackToken := os.Getenv("SLACK_TOKEN")
		if slackToken == "" {
			log.Errorf(c, "%s", "no slack api key")
			w.Write([]byte("error"))
			return
		}

		storiesChannel := os.Getenv("STORIES_CHANNEL")
		if storiesChannel == "" {
			log.Errorf(c, "%s", "no stories  channel")
			w.Write([]byte("error"))
			return
		}

		httpClient := urlfetch.Client(c)
		slack.SetHTTPClient(httpClient)

		if err := go_trend_slack_bot.PushStories(httpClient, slackToken, storiesChannel); err != nil {
			log.Errorf(c, "%s", err.Error())
			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	}
}
