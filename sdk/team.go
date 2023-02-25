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

// CreateTeam - Creates a new Team
func (c *Client) CreateTeam(team Team) (*Team, error) {

	rb, err := json.Marshal(team)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/team", c.ApiURL),
		strings.NewReader(string(rb)),
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	team = Team{}

	err = json.Unmarshal(body, &team)

	if err != nil {
		return nil, err
	}

	return &team, nil
}

// UpdateTeam - Updates a new Team
func (c *Client) UpdateTeam(team Team) (*Team, error) {

	rb, err := json.Marshal(team)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/team/%s", c.ApiURL, team.ID),
		strings.NewReader(string(rb)),
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	team = Team{}

	err = json.Unmarshal(body, &team)

	if err != nil {
		return nil, err
	}

	return &team, nil
}

// GetTeam - Gets a new Team
func (c *Client) GetTeam(teamId string) (*Team, error) {

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/team/%s", c.ApiURL, teamId),
		nil,
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	team := Team{}

	err = json.Unmarshal(body, &team)

	if err != nil {
		return nil, err
	}

	return &team, nil
}

// DeleteTeam - Deletes a Team
func (c *Client) DeleteTeam(teamId string) error {

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/team/%s", c.ApiURL, teamId),
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
