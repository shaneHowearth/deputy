// Package main -
package main

import (
	"fmt"
	"log"

	"github.com/shanehowearth/deputy/roles"
)

func main() {
	rolesInput := ""
	r := roles.NewRoleCollection()
	err := r.SetRoles(rolesInput)
	if err != nil {
		log.Fatalf("unable to set roles with error %v", err)
	}

	usersInput := ""
	err = r.SetUsers(usersInput)
	if err != nil {
		log.Fatalf("unable to set uses with error %v", err)
	}
	user := 1
	fmt.Println(r.GetSubOrdinates(user))
}
