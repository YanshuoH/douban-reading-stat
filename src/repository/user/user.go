package user

import(
  "time"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"

  "models"
)

func Get(userId string, con *mgo.Database) (models.User, error) {
  user := models.User{}

  // Try to get user by id
  err := con.C(models.CollectionUser).Find(bson.M{"userId": userId}).One(&user)
  // Try to get user by real name
  if err != nil {
    err = con.C(models.CollectionUser).Find(bson.M{"name": userId}).One(&user)
  }

  return user, err
}

func Save(user models.User, con *mgo.Database) (models.User, error) {
  if time.Time.IsZero(user.CreatedOn) {
    user.CreatedOn = time.Now()
  }
  user.UpdatedOn = time.Now()

  err := con.C(models.CollectionUser).Insert(user)

  return user, err
}
