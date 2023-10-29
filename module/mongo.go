package module

import (
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	model "github.com/angkringankuy/backendyak/model"
)

func SetConnection(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		// DBString: "mongodb+srv://admin:admin@projectexp.pa7k8.gcp.mongodb.net", //os.Getenv(MONGOCONNSTRINGENV),
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func IsPasswordValid(mongoconn *mongo.Database, collection string, userdata model.User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[model.User](mongoconn, collection, filter)
	return CheckPasswordHash(userdata.Password, res.Password)
}