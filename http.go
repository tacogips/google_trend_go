package google_trend_go

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

const (
	generalURL          = "https://trends.google.com/trends/api/explore"
	interestovertimeURL = "https://trends.google.com/trends/api/widgetdata/multiline"
	interestbyregionURL = "https://trends.google.com/trends/api/widgetdata/comparedgeo"
	relatedqueriesURL   = "https://trends.google.com/trends/api/widgetdata/relatedsearches"
	suggestionsURL      = "https://trends.google.com/trends/api/autocomplete/"
)

func get(cli *http.Client, url string, queryStrings url.Values) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, vs := range queryStrings {
		for _, v := range vs {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid response: code[%d] body[%s] ", resp.StatusCode, string(bytes))
	}

	return bytes, nil
}

func postForm(cli *http.Client, url string, form url.Values) ([]byte, error) {
	bodyBuffer := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuffer)

	for k, v := range form {
		bodyWriter.WriteField(k, v[0])
	}

	req, err := http.NewRequest("POST", url, bodyBuffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())
	req.Header.Add("Content-Length", strconv.Itoa(len(bodyBuffer.Bytes())))
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Expect", "100-continue")
	req.Header.Del("Accept-Encoding")

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid response: code[%d] body[%s] ", resp.StatusCode, string(bytes))
	}

	return bytes, nil
}
