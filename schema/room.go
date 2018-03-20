package schema

import "gopkg.in/mgo.v2/bson"

// Room describes user document in MongoDB
type Room struct {
	ID         bson.ObjectId `bson:"_id"`
	Hotel      interface{}   `bson:"hotel"`
	RoomNumber int           `bson:"room_number"`
	Category   string        `bson:"category"`
	PriceMod   []string      `bson:"price_mod"`
	Beds       int           `bson:"beds"`
	Shower     bool          `bson:"shower"`
	Photos     []string      `bson:"photos"`
	IsVacant   bool          `bson:"vacant"`
	Issues     []string      `bson:"issues"`
	Guests     []string      `bson:"giests"`
}
