package usermanagement

type User struct {
	ID       string  `json:"id" bson:"user_id"`
	Name     string  `json:"name" bson:"name"`
	Email    string  `json:"email" bson:"email"`
	Password string  `json:"password" bson:"password"`
	Profile  Profile `json:"profile" bson:"profile"`
}

type Profile struct {
	ID      string `json:"id" bson:"profile_id"`
	Surname string `json:"surname" bson:"surname"`
	Name    string `json:"name" bson:"name"`
}
