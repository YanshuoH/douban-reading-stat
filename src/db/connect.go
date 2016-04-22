package db

import (
  "fmt"
  "os"
  "models"

  "gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	MongoDBUrl = "mongodb://localhost:27017/douban"
)

// Connect connects to mongodb
func Connect() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo

  RegistrateModels()
}

func RegistrateModels() {
  collection := Session.DB(Mongo.Database).C(models.CollectionUser)

  // Ensure index
  index := mgo.Index{
    Key:        []string{"userId",},
    Unique:     true,
    DropDups:   true,
    Background: true,
    Sparse:     true,
  }

  err := collection.EnsureIndex(index)
  if err != nil {
    panic(err)
  }
}

func GetCon() *mgo.Database {
  return Session.DB(Mongo.Database)
}
