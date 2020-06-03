package urlscan

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// String converts string variable and literal to pointer
func String(s string) *string {
	return &s
}

// Uint64 converts uint64 variable and literal to pointer
func Uint64(u uint64) *uint64 {
	return &u
}

// Client is main structure of the library, a requester to urlscan.io.
type Client struct {
	apiKey  string
	BaseURL string
}

// NewClient is a constructor of Client
func NewClient(apiKey string) Client {
	client := Client{
		apiKey:  apiKey,
		BaseURL: "https://urlscan.io/api/v1",
	}

	return client
}

func (x Client) post(ctx context.Context, apiName string, input interface{}, output interface{}) (int, error) {
	rawData, err := json.Marshal(input)
	if err != nil {
		return 0, errors.Wrap(err, "Fail to marshal urlscan.io submit argument")
	}

	uri := fmt.Sprintf("%s/%s/", x.BaseURL, apiName)

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "POST", uri, bytes.NewReader(rawData))
	if err != nil {
		return 0, errors.Wrap(err, "Fail to create urlscan.io scan POST request")
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("API-Key", x.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, errors.Wrap(err, "Fail to send urlscan.io POST request")
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, errors.Wrap(err, "Fail to read urlscan.io POST result")
	}

	err = json.Unmarshal(buf, &output)
	if err != nil {
		return resp.StatusCode, errors.Wrap(err, "Fail to unmarshal urlscan.io POST result")
	}

	return resp.StatusCode, nil
}

func (x Client) get(ctx context.Context, apiName string, values url.Values, output interface{}) (int, error) {
	var qs string
	if values != nil {
		qs = "?" + values.Encode()
	}

	uri := fmt.Sprintf("%s/%s/%s", x.BaseURL, apiName, qs)

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return 0, errors.Wrap(err, "Fail to create urlscan.io get request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, errors.Wrap(err, "Fail to send urlscan.io get request")
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, errors.Wrap(err, "Fail to read urlscan.io get result")
	}

	err = json.Unmarshal(buf, &output)
	if err != nil {
		return resp.StatusCode, errors.Wrap(err, "Fail to unmarshal urlscan.io get result")
	}

	return resp.StatusCode, nil
}
