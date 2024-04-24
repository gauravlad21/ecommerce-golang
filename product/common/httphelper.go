package common

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpCall(ctx context.Context, method, url string, headerMap, queryMap map[string]string, body []byte) ([]byte, error) {

	var err error
	var request *http.Request
	if body == nil {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	}
	if err != nil {
		return nil, err
	}

	request = addHeaders(ctx, request, headerMap)
	request = addQuery(ctx, request, queryMap)

	client := &http.Client{Timeout: 1 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func addHeaders(ctx context.Context, request *http.Request, headerMap map[string]string) *http.Request {
	for key, val := range headerMap {
		request.Header.Add(key, val)
	}
	return request
}

func addQuery(ctx context.Context, request *http.Request, queryMap map[string]string) *http.Request {
	if queryMap != nil {
		q := request.URL.Query()
		for key, val := range queryMap {
			q.Add(key, val)
		}
		request.URL.RawQuery = q.Encode()
	}
	return request
}
