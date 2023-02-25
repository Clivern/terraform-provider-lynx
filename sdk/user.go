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

// CreateUser - Creates a new User
func (c *Client) CreateUser(user User) (*User, error) {
	rb, err := json.Marshal(user)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/user", c.ApiURL), strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	user = User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser - Deletes a User
func (c *Client) DeleteUser(userId string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/user/%s", c.ApiURL, userId), nil)

	if err != nil {
		return err
	}

	_, err = c.doRequest(req)

	if err != nil {
		return err
	}

	return nil
}
