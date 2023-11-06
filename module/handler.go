package module

import (
	"encoding/json"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/whatsauth/watoken"

	model "github.com/angkringankuy/backendyak/model"
)

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	data := GetAllUser(mconn, collectionname)
	return GCFReturnStruct(data)
}

func GCFPostHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response model.Credential
	Response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser model.User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
			if err != nil {
				Response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				Response.Message = "Selamat Datang"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Password Salah"
		}
	}

	return GCFReturnStruct(Response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func Register(Mongoenv, dbname string, r *http.Request) string {
    resp := new(model.Credential)
    userdata := new(model.User)
    resp.Status = false
    conn := SetConnection(Mongoenv, dbname)
    err := json.NewDecoder(r.Body).Decode(&userdata)
    if err != nil {
        resp.Message = "error parsing application/json: " + err.Error()
    } else {
        resp.Status = true
        insertedID, err := InsertUser(conn, "user", *userdata)
        if err != nil {
            resp.Message = "Gagal memasukkan data ke database: " + err.Error()
        } else {
            resp.Message = "Berhasil Input data dengan ID: " + insertedID.Hex()
        }
    }
    return GCFReturnStruct(resp)
}

func DeleteUsers(Mongoenv, dbname string, r *http.Request) string {
    resp := new(model.Credential)
    userdata := new(model.User)
    resp.Status = false
    conn := SetConnection(Mongoenv, dbname)
    err := json.NewDecoder(r.Body).Decode(&userdata)
    if err != nil {
        resp.Message = "error parsing application/json: " + err.Error()
    } else {
        resp.Status = true
        err := DeleteUser(conn, "user", userdata.Username) // Menggunakan DeleteUser untuk menghapus data berdasarkan username
        if err != nil {
            resp.Message = "Gagal menghapus data dari database: " + err.Error()
        } else {
            resp.Message = "Berhasil menghapus data dengan username: " + userdata.Username
        }
    }
    return GCFReturnStruct(resp)
}

func ResetPassword(Mongoenv, dbname string, r *http.Request) string {
    resp := new(model.Credential)
    userdata := new(model.User)
    resp.Status = false
    conn := SetConnection(Mongoenv, dbname)
    err := json.NewDecoder(r.Body).Decode(&userdata)
    if err != nil {
        resp.Message = "error parsing application/json: " + err.Error()
    } else {
        resp.Status = true
        // Reset password logic
        pass, _ := HashPassword(userdata.Password)
        update := bson.M{"password": pass}
        filter := bson.M{"username": userdata.Username}
        err := UpdateOneDoc(conn, "user", filter, update)
        if err != nil {
            resp.Message = "Gagal mereset password: " + err.Error()
        } else {
            resp.Message = "Berhasil mereset password: " + userdata.Username
        }
    }
    return GCFReturnStruct(resp)
}

func CSMessage(Mongoenv, dbname string, r *http.Request) string {
    resp := new(model.Credential)
    data := new(model.Helper)
    resp.Status = false
    conn := SetConnection(Mongoenv, dbname)
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        resp.Message = "error parsing application/json: " + err.Error()
    } else {
        resp.Status = true
        insertedID, err := CustomerService(conn, "help", data.Name, data.Email, data.Message)
        if err != nil {
            resp.Message = "Gagal memasukkan data ke database: " + err.Error()
        } else {
            resp.Message = "Berhasil Input data dengan ID: " + insertedID.Hex()
        }
    }
    return GCFReturnStruct(resp)
}

func Transaksi(Mongoenv, dbname string, r *http.Request) string {
    resp := new(model.Credential)
    data := new(model.Transaksi)
    resp.Status = false
    conn := SetConnection(Mongoenv, dbname)
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        resp.Message = "error parsing application/json: " + err.Error()
    } else {
        resp.Status = true
        insertedID, err := InsertPayment(conn, "transaksi", data.Email, data.OrderID)
        if err != nil {
            resp.Message = "Gagal memasukkan data ke database: " + err.Error()
        } else {
            resp.Message = "Berhasil Input data dengan ID: " + insertedID.Hex()
        }
    }
    return GCFReturnStruct(resp)
}







