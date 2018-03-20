package schema

// User describes user document in MongoDB
type User struct {
	ID        string `bson:"_id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`
}
