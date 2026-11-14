package main

import "github.com/tacogips/google_trend_go"

func main() {

	hottrend := google_trend_go.NewHotTrend(google_trend_go.DefaultHotTrendConfig)
	resp, err := hottrend.Fetch(google_trend_go.GeoUS)
	if err != nil {
		panic(err)
	}

	println("# Latest hot trend item")
	for _, eachDt := range resp.TrendsByDate {
		for _, trend := range eachDt.Trends {
			println(trend.Title)
		}
	}

}
