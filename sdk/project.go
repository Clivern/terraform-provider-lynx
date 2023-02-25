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

// CreateProject - Creates a new Project
func (c *Client) CreateProject(project Project) (*Project, error) {

	rb, err := json.Marshal(project)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/project", c.ApiURL),
		strings.NewReader(string(rb)),
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	project = Project{}

	err = json.Unmarshal(body, &project)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

// UpdateProject - Updates a new Project
func (c *Client) UpdateProject(project Project) (*Project, error) {

	rb, err := json.Marshal(project)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/project/%s", c.ApiURL, project.ID),
		strings.NewReader(string(rb)),
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	project = Project{}

	err = json.Unmarshal(body, &project)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

// GetProject - Gets a new Project
func (c *Client) GetProject(projectId string) (*Project, error) {

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/project/%s", c.ApiURL, projectId),
		nil,
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	project = Project{}

	err = json.Unmarshal(body, &project)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

// DeleteProject - Deletes a Project
func (c *Client) DeleteProject(projectId string) error {

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/project/%s", c.ApiURL, projectId),
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
