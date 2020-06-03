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

const (
	// BaseURL for urlscan API
	BaseURL = "https://urlscan.io/api/v1"
)

// Client is main structure of the library, a requester to urlscan.io.
type Client struct {
	apiKey string
}

// NewClient is a constructor of Client
func NewClient(apiKey string) Client {
	client := Client{
		apiKey: apiKey,
	}

	return client
}

// req Make a request and fill an output structure
// apiKey is only required for POST requests, otherwise it can be blank
func req(ctx context.Context, method string, values *url.Values, apiName string, input interface{}, output interface{}, apiKey string) (int, error) {
	body := &bytes.Buffer{}
	if input != nil {
		rawData, err := json.Marshal(input)
		if err != nil {
			return 0, errors.Wrap(err, "Fail to marshal urlscan.io submit argument")
		}
		body = bytes.NewBuffer(rawData)
	}

	uri := fmt.Sprintf("%s/%s/", BaseURL, apiName)

	// Add url values
	if values != nil && len(*values) != 0 {
		uri = fmt.Sprintf("%s?%s", uri, values.Encode())
	}

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return 0, errors.Wrap(err, "Fail to create urlscan.io scan POST request")
	}

	// Add headers if we are posting
	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("API-Key", apiKey)
	}

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
