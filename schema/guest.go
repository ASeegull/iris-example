package schema

// Guest describes user document in MongoDB
type Guest struct {
	ID        string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"first_name" json:"first_name,omitempty"`
	LastName  string `bson:"last_name" json:"last_name,omitempty"`
	Email     string `bson:"email" json:"email,omitempty"`
	Phone     string `bson:"phone" json:"phone,omitempty"`
}
