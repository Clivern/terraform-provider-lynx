// Copyright 2024 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

const (
	RegularUser = "regular"
	SuperUser   = "super"
)

// User Model
type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
	Password string `json:"password,omitempty"`
	Verified bool   `json:"verified,omitempty"`
}

// Team Model
type Team struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
}

// Project Model
type Project struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	Team        Team   `json:"team,omitempty"`
}

// Environment Model
type Environment struct {
	ID       string  `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Slug     string  `json:"slug,omitempty"`
	Username string  `json:"username,omitempty"`
	Secret   string  `json:"secret,omitempty"`
	Project  Project `json:"project,omitempty"`
}

// Snapshot Model
type Snapshot struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	RecordType  string `json:"record_type,omitempty"`
	RecordID    string `json:"record_uuid,omitempty"`
	Team        Team   `json:"team,omitempty"`
}
