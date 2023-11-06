package backendyak

import (
	"fmt"
	"testing"
	module "github.com/angkringankuy/backendyak/module"

)

// user
func TestInsertUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	var userdata User
	userdata.Username = "syahid"
	userdata.Role = "admin"
	userdata.Password = "kepoah"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}

func TestGetAllUserFromUsername(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	anu := module.GetUserFromUsername(mconn, "user", "syahid")
	fmt.Println(anu)
}

func TestGetAllUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	anu := module.GetAllUser(mconn, "user")
	fmt.Println(anu)
}