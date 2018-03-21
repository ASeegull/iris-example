package schema

// Guest describes user document in MongoDB
type Guest struct {
	ID        string `bson:"_id" json:"_id"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Email     string `bson:"email" json:"email"`
	Phone     string `bson:"phone" json:"phone"`
}
