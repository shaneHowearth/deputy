package roles

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetRoles(t *testing.T) {
	testcases := map[string]struct {
		tRoles []Role
		err    error
		state  RoleCollection
	}{
		"Supplied Example": {
			tRoles: []Role{
				{
					ID:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				{
					ID:     2,
					Name:   "Location Manager",
					Parent: 1,
				},
				{
					ID:     3,
					Name:   "Supervisor",
					Parent: 2,
				},
				{
					ID:     4,
					Name:   "Employee",
					Parent: 3,
				},
				{
					ID:     5,
					Name:   "Trainer",
					Parent: 3,
				},
			},
			state: RoleCollection{Roles: map[int]*Role{
				1: {
					ID:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				2: {
					ID:     2,
					Name:   "Location Manager",
					Parent: 1,
				},
				3: {
					ID:     3,
					Name:   "Supervisor",
					Parent: 2,
				},
				4: {
					ID:     4,
					Name:   "Employee",
					Parent: 3,
				},
				5: {
					ID:     5,
					Name:   "Trainer",
					Parent: 3,
				},
			},
			},
		},
		"No roles": {},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			r := NewRoleCollection()
			err := r.setRoles(tc.tRoles)
			if tc.err == nil {
				assert.Nil(t, err, "set roles was not expected to generate an error but got %v", err)
				// Roles have the same number of roles
				assert.Equal(t, len(tc.state.Roles), len(r.Roles), "different number of roles in state, expected %d got %d", len(tc.state.Roles), len(r.Roles))
				for key := range tc.state.Roles {
					// each Role exists in each roles.Roles
					if role, ok := r.Roles[key]; ok {
						assert.Equal(t, tc.state.Roles[key].ID, role.ID, "role IDs did not match, expecting %d, got %d", tc.state.Roles[key].ID, role.ID)
						assert.Equal(t, tc.state.Roles[key].Name, role.Name, "role Names did not match, expecting %s, got %s", tc.state.Roles[key].Name, role.Name)
						assert.Equal(t, tc.state.Roles[key].Parent, role.Parent, "role Parents did not match, expecting %d, got %d", tc.state.Roles[key].Parent, role.Parent)
						assert.Equal(t, tc.state.Roles[key].Users, role.Users, "role Users did not match, expecting %v, got %v", tc.state.Roles[key].Users, role.Users)
					} else {
						t.Fatalf("key missing from r %d", key)
					}
				}
			} else {
				assert.NotNil(t, err, "set roles was expected to generate an error but got %v", err)
			}
		})
	}
}

func TestSetUsers(t *testing.T) {
	testcases := map[string]struct {
		tRoles []Role
		tUsers []User
		err    error
		state  RoleCollection
	}{
		"Supplied Example": {
			tRoles: []Role{
				{
					ID:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				{
					ID:     2,
					Name:   "Location Manager",
					Parent: 1,
				},
				{
					ID:     3,
					Name:   "Supervisor",
					Parent: 2,
				},
				{
					ID:     4,
					Name:   "Employee",
					Parent: 3,
				},
				{
					ID:     5,
					Name:   "Trainer",
					Parent: 3,
				},
			},
			state: RoleCollection{
				Roles: map[int]*Role{
					1: {
						ID:     1,
						Name:   "System Administrator",
						Parent: 0,
						Users: []User{
							{
								ID:   1,
								Name: "Adam Admin",
								Role: 1,
							},
						},
					},
					2: {
						ID:     2,
						Name:   "Location Manager",
						Parent: 1,
						Users: []User{
							{
								ID:   4,
								Name: "Mary Manager",
								Role: 2,
							},
						},
					},
					3: {
						ID:     3,
						Name:   "Supervisor",
						Parent: 2,
						Users: []User{
							{
								ID:   3,
								Name: "Sam Supervisor",
								Role: 3,
							},
						},
					},
					4: {
						ID:     4,
						Name:   "Employee",
						Parent: 3,
						Users: []User{
							{
								ID:   2,
								Name: "Emily Employee",
								Role: 4,
							},
						},
					},
					5: {
						ID:     5,
						Name:   "Trainer",
						Parent: 3,
						Users: []User{
							{
								ID:   5,
								Name: "Steve Trainer",
								Role: 5,
							},
						},
					},
				},
			},
			tUsers: []User{
				{
					ID:   1,
					Name: "Adam Admin",
					Role: 1,
				},
				{
					ID:   2,
					Name: "Emily Employee",
					Role: 4,
				},
				{
					ID:   3,
					Name: "Sam Supervisor",
					Role: 3,
				},
				{
					ID:   4,
					Name: "Mary Manager",
					Role: 2,
				},
				{
					ID:   5,
					Name: "Steve Trainer",
					Role: 5,
				},
			},
		},
		"No roles": {
			tRoles: []Role{},
			state:  RoleCollection{},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			r := NewRoleCollection()
			err := r.setRoles(tc.tRoles)
			assert.Nil(t, err, "set roles was not expected to generate an error but got %v", err)
			err = r.setUsers(tc.tUsers)
			if tc.err == nil {
				assert.Nil(t, err, "set users was not expected to generate an error but got %v", err)
				// Roles have the same number of roles
				assert.Equal(t, len(tc.state.Roles), len(r.Roles), "different number of roles in state, expected %d got %d", len(tc.state.Roles), len(r.Roles))
				for key := range tc.state.Roles {
					// each Role exists in each Roles
					if role, ok := r.Roles[key]; ok {
						assert.Equal(t, tc.state.Roles[key].ID, role.ID, "role IDs did not match, expecting %d, got %d", tc.state.Roles[key].ID, role.ID)
						assert.Equal(t, tc.state.Roles[key].Name, role.Name, "role Names did not match, expecting %s, got %s", tc.state.Roles[key].Name, role.Name)
						assert.Equal(t, tc.state.Roles[key].Parent, role.Parent, "role Parents did not match, expecting %d, got %d", tc.state.Roles[key].Parent, role.Parent)
						assert.Equal(t, tc.state.Roles[key].Users, role.Users, "role Users did not match, expecting %v, got %v", tc.state.Roles[key].Users, role.Users)
					} else {
						t.Fatalf("key missing from r %d", key)
					}
				}
			} else {
				assert.NotNil(t, err, "set users was expected to generate an error but got %v", err)
			}
		})
	}
}

func TestSetRolesJsonUnMarshal(t *testing.T) {

	jsonUnmarshal = func([]byte, interface{}) error {
		return fmt.Errorf("test error")
	}

	// Reset jsonUnmarshal variable when the test exits
	defer func() { jsonUnmarshal = json.Unmarshal }()

	r := NewRoleCollection()
	err := r.SetRoles("Nothing needed")
	assert.NotNil(t, err, "expected an error but got nil")
	assert.EqualError(t, fmt.Errorf("set roles was unable to unmarshal json with error test error"), err.Error(), "did not get the expected set roles error")
}

func TestSetUsersJsonUnMarshal(t *testing.T) {

	jsonUnmarshal = func([]byte, interface{}) error {
		return fmt.Errorf("test error")
	}

	// Reset jsonUnmarshal variable when the test exits
	defer func() { jsonUnmarshal = json.Unmarshal }()

	r := NewRoleCollection()
	err := r.SetUsers("Nothing needed")
	assert.NotNil(t, err, "expected an error but got nil")
	assert.EqualError(t, fmt.Errorf("set users was unable to unmarshal json with error test error"), err.Error(), "did not get the expected set users error")
}
