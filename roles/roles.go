// Package roles -
package roles

import (
	"encoding/json"
	"fmt"
	"sort"
)

// RoleCollection -
type RoleCollection struct {
	Roles map[int]*Role
	Users map[int]*User
	max   int // highest role ID
}

// Role -
type Role struct {
	ID     int    `json:"Id"`
	Name   string `json:"Name"`
	Parent int    `json:"Parent"`
	Users  []User
}

// User -
type User struct {
	ID   int    `json:"Id"`
	Name string `json:"Name"`
	Role int    `json:"Role"`
}

// NewRoleCollection -
func NewRoleCollection() *RoleCollection {
	r := RoleCollection{
		Roles: map[int]*Role{},
		Users: map[int]*User{},
	}
	return &r
}

// Make json.Unmarshal easy to test
var jsonUnmarshal = json.Unmarshal

// SetRoles - input is expected to be valid json
func (r *RoleCollection) SetRoles(input string) error {
	roles := []Role{}
	if err := jsonUnmarshal([]byte(input), &roles); err != nil {
		return fmt.Errorf("set roles was unable to unmarshal json with error %w", err)
	}
	return r.setRoles(roles)
}

func (r *RoleCollection) setRoles(roles []Role) error {
	for idx := range roles {
		if roles[idx].ID > r.max {
			r.max = roles[idx].ID
		}
		// Insert
		r.Roles[roles[idx].ID] = &roles[idx]
	}
	return nil
}

// SetUsers - input is expected to be valid json
func (r *RoleCollection) SetUsers(input string) error {
	users := []User{}
	if err := jsonUnmarshal([]byte(input), &users); err != nil {
		return fmt.Errorf("set users was unable to unmarshal json with error %w", err)
	}
	return r.setUsers(users)
}

func (r *RoleCollection) setUsers(users []User) error {
	for idx := range users {
		r.Users[users[idx].ID] = &users[idx]
		// Append to the []Users in the correct Role
		r.Roles[users[idx].Role].Users = append(r.Roles[users[idx].Role].Users, users[idx])
	}
	return nil
}

// GetSubOrdinates -
func (r *RoleCollection) GetSubOrdinates(userID int) []User {
	users := []User{}
	for i := r.Users[userID].Role + 1; i <= r.max; i++ {
		if r.Roles[i] != nil {
			users = append(users, r.Roles[i].Users...)
		}
	}
	// Sort the user slice to make it pretty.
	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})
	return users
}
