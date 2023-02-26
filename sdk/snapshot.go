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

// CreateSnapshot - Creates a new Snapshot
func (c *Client) CreateSnapshot(snapshot Snapshot) (*Snapshot, error) {

	snapshot.TeamId = snapshot.Team.ID

	rb, err := json.Marshal(snapshot)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/snapshot", c.ApiURL),
		strings.NewReader(string(rb)),
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	snapshot = Snapshot{}

	err = json.Unmarshal(body, &snapshot)

	if err != nil {
		return nil, err
	}

	snapshot.TeamId = snapshot.Team.ID

	return &snapshot, nil
}

// GetSnapshot - Gets a new Snapshot
func (c *Client) GetSnapshot(snapshotId string) (*Snapshot, error) {

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/snapshot/%s", c.ApiURL, snapshotId),
		nil,
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	snapshot := Snapshot{}

	err = json.Unmarshal(body, &snapshot)

	if err != nil {
		return nil, err
	}

	snapshot.TeamId = snapshot.Team.ID

	return &snapshot, nil
}

// DeleteSnapshot - Deletes a Snapshot
func (c *Client) DeleteSnapshot(snapshotId string) error {

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/snapshot/%s", c.ApiURL, snapshotId),
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
