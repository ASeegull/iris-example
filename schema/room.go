package schema

import "gopkg.in/mgo.v2/bson"

// Room describes user document in MongoDB
type Room struct {
	ID         bson.ObjectId `bson:"_id" json:"_id"`
	Hotel      interface{}   `bson:"hotel" json:"hotel"`
	RoomNumber int           `bson:"room_number" json:"room_number"`
	Category   string        `bson:"category" json:"category"`
	PriceMod   []string      `bson:"price_mod" json:"price_mod"`
	Beds       int           `bson:"beds" json:"beds"`
	Shower     bool          `bson:"shower" json:"shower"`
	Photos     []string      `bson:"photos" json:"photos"`
	IsVacant   bool          `bson:"vacant" json:"vacant"`
	Issues     []string      `bson:"issues" json:"issues"`
	Guests     []string      `bson:"guests" json:"guests"`
}
