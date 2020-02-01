package __TPL__

import "github.com/chiachan163/cc-orm/v1/model/mongo"

// __MYSQL_MODEL__ create mysql model
type __MYSQL_MODEL__ struct {
	RoomFurnituresLog
}

// __MONGO_MODEL__ create mongodb model
type __MONGO_MODEL__ struct {
	Furnitures
}

// mongo
type (
	// 位置坐标
	Position struct {
		X int32 `json:"x"`
		Y int32 `json:"y"`
		Z int32 `json:"z"`
	}
	// 家具ID+数量
	FurnitureHas struct {
		Id  int32 `json:"id"`  // 家具ID
		Num int32 `json:"num"` // 已购数量
	}
	// 家具
	Furniture struct {
		Id       int32     `json:"id"`       // 家具ID
		IsFront  int32     `json:"is_front"` // 方向是否朝前[1表示前 -1表示后]
		IsLeft   int32     `json:"is_left"`  // 方向是否朝左[1表示右 -1表示左]
		Position *Position `json:"position"` // 位置坐标
	}
	// 收藏套装
	Favorite struct {
		Fid   int32        `json:"fid"`   // 套装位ID
		Name  string       `json:"name"`  // 套装名称
		Image string       `json:"image"` // 套装截图
		Use   []*Furniture `json:"use"`   // 套装使用的家具
	}
	Furnitures struct {
		Id        mongo.ObjectId  `key:"pri" bson:"_id"`
		CoupleId  string          `key:"uni" bson:"couple_id"`
		Has       []*FurnitureHas `bson:"has"`
		Use       []*Furniture    `bson:"use"`
		Favorites []*Favorite     `bson:"favorites"`
		UpdateUid int64           `bson:"update_uid"`
		CreatedAt int64           `bson:"created_at"`
		UpdatedAt int64           `bson:"updated_at"`
		DeletedTs int64           `bson:"deleted_ts"`
	}
)

type RoomFurnituresLog struct {
	Id        int32  `key:"pri" json:"id"`        // 自增ID
	CoupleId  string `key:"uni" json:"couple_id"` // 情侣ID
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedTs int64  `json:"deleted_ts"`
}
