// Copyright 2024 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateEnvironment - Creates a new Environment
func (c *Client) CreateEnvironment(environment Environment) (*Environment, error) {
	rb, err := json.Marshal(environment)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/project/%s/environment", c.ApiURL, environment.Project.ID),
		strings.NewReader(string(rb)),
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	environment = Environment{}

	err = json.Unmarshal(body, &environment)

	if err != nil {
		return nil, err
	}

	return &environment, nil
}

// UpdateEnvironment - Updates an Environment
func (c *Client) UpdateEnvironment(environment Environment) (*Environment, error) {
	rb, err := json.Marshal(environment)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/project/%s/environment/%s", c.ApiURL, environment.Project.ID, environment.ID),
		strings.NewReader(string(rb)),
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	environment = Environment{}

	err = json.Unmarshal(body, &environment)

	if err != nil {
		return nil, err
	}

	return &environment, nil
}

// GetEnvironment - Gets a new Environment
func (c *Client) GetEnvironment(projectId, environmentId string) (*Environment, error) {

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/project/%s/environment/%s", c.ApiURL, projectId, environmentId),
		nil,
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	environment = Environment{}

	err = json.Unmarshal(body, &environment)

	if err != nil {
		return nil, err
	}

	return &environment, nil
}

// DeleteEnvironment - Deletes an Environment
func (c *Client) DeleteEnvironment(projectId, environmentId string) error {

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/project/%s/environment/%s", c.ApiURL, projectId, environmentId),
		nil,
	)

	if err != nil {
		return err
	}

	_, err = c.doRequest(req)

	if err != nil {
		return err
	}

	return nil
}
