// Copyright 2024 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// LocalApiServer -
const LocalApiServer = "http://localhost:4000/api/v1"

// Client -
type Client struct {
	ApiURL     string
	ApiKey     string
	HTTPClient *http.Client
}

// NewClient -
func NewClient(apiURL, apiKey string) *Client {
	client := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		ApiURL:     apiURL,
		ApiKey:     apiKey,
	}

	return &client
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	req.Header.Set("X-API-Key", c.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
