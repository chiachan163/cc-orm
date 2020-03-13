package __TPL__

import "github.com/chiachan163/cc-orm/model/mongo"

// __MYSQL_MODEL__ create mysql model
type __MYSQL_MODEL__ struct {
	User
}

// __MONGO_MODEL__ create mongodb model
type __MONGO_MODEL__ struct {
	Meta
}

// User user info
type User struct {
	Id        int64  `key:"pri"`
	Name      string `key:"uni"`
	Age       int32  `json:"age"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedTs int64  `json:"deleted_ts"`
}

type Meta struct {
	Id        mongo.ObjectId `key:"pri" bson:"_id"`
	Uid       int64          `key:"uni" bson:"uid"`
	Hobby     []string       `bson:"hobby"`
	Tags      []string       `bson:"tags"`
	CreatedAt int64          `bson:"created_at"`
	UpdatedAt int64          `bson:"updated_at"`
	DeletedTs int64          `bson:"deleted_ts"`
}
