package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct for mongodb models
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Lastname  string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	UserID    string             `bson:"userId,omitempty" json:"userId,omitempty"`
	CodeAuth  string             `bson:"code_auth,omitempty" json:"code_auth,omitempty"`
	Admin     int                `bson:"admin,omitempty" json:"admin,omitempty"`
	Gender    string             `bson:"gender,omitempty" json:"gender,omitempty"`
	UrlPhoto  string             `bson:"url_photo,omitempty" json:"url_photo,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

// Request for getMyUser and createUser
type UserRequest struct {
	Name     string    `bson:"name,omitempty" json:"name,omitempty"`
	Lastname string    `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Email    string    `bson:"email,omitempty" json:"email,omitempty"`
	Password string    `bson:"password,omitempty" json:"password,omitempty"`
	BornOn   time.Time `bson:"bornOn,omitempty" json:"bornOn,omitempty"`
}

// Response for getMyUser and createUser
type MyUserResponse struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Lastname  string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	BornOn    time.Time          `bson:"bornOn,omitempty" json:"bornOn,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Logged    bool               `bson:"logged,omitempty" json:"logged,omitempty"`
	UserID    string             `bson:"userId,omitempty" json:"userId,omitempty"`
}

type ResponseUser struct {
	User   MyUserResponse `json:"user,omitempty"`
	Sucess bool           `json:"sucess,omitempty"`
	Errors []Error        `json:"errors,omitempty"`
}
