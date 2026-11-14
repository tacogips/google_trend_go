package google_trend_go

import "testing"

func TestHotTrend(t *testing.T) {
	ht := NewHotTrend(DefaultHotTrendConfig)
	trendResponse, err := ht.Fetch(GeoJP)
	if err != nil {
		t.Error(err)
	}

	if len(trendResponse.TrendsByDate) == 0 {
		t.Errorf("hottrend empty : %+v", trendResponse)
	}

}
