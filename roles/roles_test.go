package roles_test

import (
	"testing"

	"github.com/shanehowearth/deputy/roles"
	"github.com/stretchr/testify/assert"
)

func TestGetSubOrdinates(t *testing.T) {
	tRoles := `[{
"Id": 1,
"Name": "System Administrator",
"Parent": 0
},
{
"Id": 2,
"Name": "Location Manager",
"Parent": 1
},
{
"Id": 3,
"Name": "Supervisor",
"Parent": 2
},
{
"Id": 4,
"Name": "Employee",
"Parent": 3
},
{
"Id": 5,
"Name": "Trainer",
"Parent": 3
}
]`
	tUsers := `[
{
"Id": 1,
"Name": "Adam Admin",
"Role": 1
},
{
"Id": 2,
"Name": "Emily Employee",
"Role": 4
},
{
"Id": 3,
"Name": "Sam Supervisor",
"Role": 3
},
{
"Id": 4,
"Name": "Mary Manager",
"Role": 2
},
{"Id": 5,
"Name": "Steve Trainer",
"Role": 5
}
]`
	r := roles.NewRoleCollection()
	err := r.SetRoles(tRoles)
	assert.Nil(t, err, "set roles was not expected to generate an error but got %v", err)
	err = r.SetUsers(tUsers)
	assert.Nil(t, err, "set users was not expected to generate an error but got %v", err)
	testcases := map[string]struct {
		output []roles.User
		user   int
	}{
		"First Supplied Example": {
			user: 3,
			output: []roles.User{
				{2, "Emily Employee", 4},
				{5, "Steve Trainer", 5},
			},
		},
		"Second Supplied Example": {
			user: 1,
			output: []roles.User{
				{2, "Emily Employee", 4},
				{3, "Sam Supervisor", 3},
				{4, "Mary Manager", 2},
				{5, "Steve Trainer", 5},
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {

			output := r.GetSubOrdinates(tc.user)
			assert.Equal(t, len(tc.output), len(output), "Length of output does not match expected, wanted: %d, got %d", len(tc.output), len(output))
			for idx := range tc.output {
				assert.Equalf(t, tc.output[idx].ID, output[idx].ID, "IDs did not match, expected %d, got %d", tc.output[idx].ID, output[idx].ID)
				assert.Equalf(t, tc.output[idx].Name, output[idx].Name, "Names did not match, expected %s, got %s", tc.output[idx].Name, output[idx].Name)
				assert.Equalf(t, tc.output[idx].Role, output[idx].Role, "Roles did not match, expected %d, got %d", tc.output[idx].Role, output[idx].Role)
			}
		})
	}
}
