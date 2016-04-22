package models

import (
  "gopkg.in/mgo.v2/bson"
  "time"
)

const (
  CollectionUser = "users"
)

// User model
type User struct {
  Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
  UserId    string `json:"id" bson:"userId"`
  Name      string `json:"name" bson:"name"`
  Avatar    string `json:"avatar" bson:"large_avatar"`
  Books     []interface{} `json:"collections" binding:"required" bson:"books"`
  CreatedOn time.Time `json:"created_on" bson:"created_on" binding:"required"`
	UpdatedOn time.Time `json:"updated_on" bson:"updated_on" binding:"required"`
}

func CastApiToModel(userInfo map[string]interface{}, bookCollection []interface{}) *User {
  user := &User {
    UserId: userInfo["id"].(string),
    Name: userInfo["name"].(string),
    Avatar: userInfo["large_avatar"].(string),
    Books: bookCollection,
  }

  return user
}
