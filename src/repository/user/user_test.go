  package user

import (
  "db"
  "models"
  "os"
  "testing"

  "gopkg.in/mgo.v2"
  "github.com/stretchr/testify/assert"
)

const (
  MongoDBUrl = "mongodb://localhost:27017/douban_test"
)

var userModelAsset1 = models.User {
  UserId: "2274326",
  Name: "a testing user",
  Avatar: "https://img3.doubanio.com/icon/up74783952-2.jpg",
  Books: make([]interface{}, 3),
}

func setup() *mgo.Database {
  os.Setenv("MONGODB_URL", MongoDBUrl)
  db.Connect()

  return db.GetCon()
}

func tearDown() {
  // Purge test db
  con := db.GetCon()
  err := con.C(models.CollectionUser).DropCollection()
	if err != nil {
		panic(err)
	}
}

func TestSave(t *testing.T) {
  con := setup()
  assert := assert.New(t)

  user, err := Save(userModelAsset1, con)
  assert.Nil(err, "Repository", "Should not be an error")
  assert.Equal(userModelAsset1.UserId, user.UserId, "Repository", "Should be the same user object")
  assert.NotZero(user.CreatedOn, "Repository", "Should have a created_on field set")
  assert.NotZero(user.UpdatedOn, "Repository", "Should have a created_on field set")

  // Double save for duplicate error handling
  user, err = Save(userModelAsset1, con)
  assert.Error(err, "Should be an duplication error")

  tearDown()
}
