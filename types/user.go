package types

type User struct {
	ID        string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName`
	LastName  string `bson:"lastName" json:"lastName`
}
