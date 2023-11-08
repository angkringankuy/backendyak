package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Role     string `bson:"role,omitempty" json:"role,omitempty"`
}

type Credential struct {
	Status  bool   `bson:"status" json:"status"`
	Token   string `bson:"token,omitempty" json:"token,omitempty"`
	Message string `bson:"message,omitempty" json:"message,omitempty"`
}

type Helper	struct {
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Message  string				`bson:"message,omitempty" json:"message,omitempty"`
}

type InfoTransaksi struct {
	Email   	 string      		`bson:"email,omitempty" json:"email,omitempty"`
	InfoBeli     string				`bson:"infobeli,omitempty" json:"infobeli,omitempty"`
	OrderID  	 string    			`bson:"orderid,omitempty" json:"orderid,omitempty"`
	NoHP	 	 string				`bson:"nohp,omitempty" json:"nohp,omitempty"`
	
}

