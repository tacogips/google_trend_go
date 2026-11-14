package main

import (
	"time"

	"github.com/tacogips/google_trend_go"
)

func main() {

	stories := google_trend_go.NewStories(google_trend_go.DefaultStoriesConfig)
	loc, _ := time.LoadLocation("America/Chicago")

	resp, err := stories.Fetch(google_trend_go.LangEN, google_trend_go.GeoUS, loc, 10, 10)
	if err != nil {
		panic(err)
	}

	println("# Latest featured stories")
	for _, each := range resp.FeaturedStories {
		println(each.Title)
	}

	println("# Latest trend stories")
	for _, each := range resp.TrendStories {
		println(each.Title)
	}

}
