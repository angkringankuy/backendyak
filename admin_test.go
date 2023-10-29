package backendyak

import (
	"fmt"
	"testing"

	model "github.com/angkringankuy/backendyak/model"
	module "github.com/angkringankuy/backendyak/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user
func TestInsertUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	var userdata User
	userdata.Email = "masyahida4@gmail.com"
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