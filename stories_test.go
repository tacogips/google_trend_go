package google_trend_go

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStories(t *testing.T) {
	st := NewStories(DefaultStoriesConfig)

	jst, _ := time.LoadLocation("Asia/Tokyo")
	resp, err := st.Fetch(LangJA, GeoJP, jst, 10, 10)
	if err != nil {
		t.Error(err)
	} else {
		if len(resp.TrendStories) == 0 {
			t.Errorf("story summary empty : %#v", len(resp.TrendStories))
		}
	}
}

func TestStoriesIDs(t *testing.T) {

	st := NewStories(DefaultStoriesConfig)

	jst, _ := time.LoadLocation("Asia/Tokyo")

	resp, err := st.FetchSummaryIDs(LangJA, GeoJP, jst)

	if err != nil {
		t.Error(err)
	} else {
		if len(resp.TrendStoryIDs) == 0 {
			t.Errorf("story summary ids empty : %#v", resp)
		}
	}

}

func TestStrChunk(t *testing.T) {

	chunkSize := 3
	{
		testDatas := []string{}

		for i := 0; i < 1; i++ {
			testDatas = append(testDatas, strconv.Itoa(i))
		}

		r := strChunk(testDatas)
		assert.Equal(t, len(r), 1)
		assert.Equal(t, r[0], testDatas)
	}

	{
		testDatas := []string{}

		for i := 0; i < chunkSize; i++ {
			testDatas = append(testDatas, strconv.Itoa(i))
		}

		r := strChunk(testDatas)
		assert.Equal(t, len(r), 1)
		assert.Equal(t, r[0], testDatas)
	}

	{
		testDatas := []string{}

		for i := 0; i < chunkSize+1; i++ {
			testDatas = append(testDatas, strconv.Itoa(i))
		}

		r := strChunk(testDatas)
		assert.Equal(t, len(r), 2)
		assert.Equal(t, r[0], testDatas[0:chunkSize])
		assert.Equal(t, r[1], testDatas[chunkSize:chunkSize+1])
	}
}

func TestToStorieAPITimeZone(t *testing.T) {

	loc, _ := time.LoadLocation("Asia/Tokyo")

	tz := toStorieAPITimeZone(loc)
	if tz != -540 {
		t.Errorf("%d", tz)
	}
}
