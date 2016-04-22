package user

import(
  "time"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"

  "models"
)

func Save(user models.User, con *mgo.Database) (models.User, error) {
  if time.Time.IsZero(user.CreatedOn) {
    user.CreatedOn = time.Now()
  }
  user.UpdatedOn = time.Now()

  err := con.C(models.CollectionUser).Insert(user)

  return user, err
}
