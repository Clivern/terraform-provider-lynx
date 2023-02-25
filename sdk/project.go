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

// DeleteProject - Deletes a Project
func (c *Client) DeleteProject(projectId string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/project/%s", c.ApiURL, projectId), nil)

	if err != nil {
		return err
	}

	_, err = c.doRequest(req)

	if err != nil {
		return err
	}

	return nil
}
