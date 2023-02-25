// Copyright 2024 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

// DeleteEnvironment - Deletes an Environment
func (c *Client) DeleteEnvironment(projectId, environmentId string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/project/%s/environment/%s", c.ApiURL, projectId, environmentId), nil)

	if err != nil {
		return err
	}

	_, err = c.doRequest(req)

	if err != nil {
		return err
	}

	return nil
}
