package user

import (
	"context"
	"encoding/base64"
	"github.com/ChatFalcon/ChatFalcon/mongo"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Token defines the user token.
type Token struct {
	Username string `json:"username" bson:"username"`
	Token    string `json:"token" bson:"_id"`
}

// Sets the password. PLEASE MAKE SURE THE PASSWORD ISN'T BLANK BEFORE RUNNNING THIS!
func (u *User) setPassword(Password string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(Password), 10)
	if err != nil {
		return err
	}
	u.HashedSaltedPassword = base64.StdEncoding.EncodeToString(b)
	return nil
}

// Insert is used to insert the user in the database. PLEASE RUN ADDITIONAL CHECKS BEFORE RUNNING THIS!
func (u *User) Create(Password string) (token string, err error) {
	err = u.setPassword(Password)
	if err != nil {
		return
	}
	_, err = mongo.DB.Collection("users").InsertOne(context.TODO(), u)
	if err != nil {
		return
	}
	token = uuid.Must(uuid.NewUUID()).String()
	_, err = mongo.DB.Collection("tokens").InsertOne(context.TODO(), &Token{
		Username: u.Username,
		Token:    token,
	})
	return
}
