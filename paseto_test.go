package testbackendp3

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

// PASETO
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("syahid", privateKey)
	fmt.Println(hasil, err)
}
func TestValidateToken(t *testing.T) {
	tokenstring := "v4.public.eyJleHAiOiIyMDIzLTEwLTIwVDAwOjM3OjQ1KzA3OjAwIiwiaWF0IjoiMjAyMy0xMC0xOVQyMjozNzo0NSswNzowMCIsImlkIjoic3lhaGlkIiwibmJmIjoiMjAyMy0xMC0xOVQyMjozNzo0NSswNzowMCJ91N0jN5aGDQlBLw33aCX75HiVm0hMvYXepxiJvNh6WTcYKO8j9LyeF555AD_vJDZSvL4lb8YHv10r-PX0FTf1BA" // Gantilah dengan token PASETO yang sesuai
	publicKey := "505404e369acbcb337dda5e44cc637df44036d7f031f773f69810e9f0f82e773"
	payload, _err := watoken.Decode(publicKey, tokenstring)
	if _err != nil {
		fmt.Println("expired token", _err)
	} else {
		fmt.Println("ID: ", payload.Id)
		fmt.Println("Di mulai: ", payload.Nbf)
		fmt.Println("Di buat: ", payload.Iat)
		fmt.Println("Expired: ", payload.Exp)
	}
}

// Hash Pass
func TestGeneratePasswordHash(t *testing.T) {
	password := "kepo"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	var userdata User
	userdata.Username = "syahid"
	userdata.Password = "kepo"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	var userdata User
	userdata.Username = "syahid"
	userdata.Password = "kepo"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

// User
func TestInsertUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	var userdata User
	userdata.Username = "syahid"
	userdata.Password = "kepo"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}
