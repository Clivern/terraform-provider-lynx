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

// DeleteSnapshot - Deletes a Snapshot
func (c *Client) DeleteSnapshot(snapshotId string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/snapshot/%s", c.ApiURL, snapshotId), nil)

	if err != nil {
		return err
	}

	_, err = c.doRequest(req)

	if err != nil {
		return err
	}

	return nil
}
