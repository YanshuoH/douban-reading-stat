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
  Stat      map[string]*StatEntity `json:"stat" bson:"stat"`
}

type StatEntity struct {
  Year int `json:"year"`
  Count int `json:"count"`
  Price float64 `json:"price"`
  Rating map[string]int `json:"rating"`
  Month map[string]int `json:"month"`
  Author map[string]int `json:"author"`
  Translator map[string]int `json:"translator"`
  Publisher map[string]int `json:"publisher"`
  Tags map[string]int `json:"tags"`
  Posters []map[string]string `json:"posters"`
}

type BookEntity struct {
  Rating map[string]interface{} `json:"rating"`
  Author []string `json:"author"`
  Translator []string `json:"translator"`
  Tags []map[string]interface{} `json:"tags"`
  Images map[string]string `json:"images"`
  Publiser string `json:"publisher"`
  Price string `json:"price"`
  Title string `json:"title"`
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
