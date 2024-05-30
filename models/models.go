// package models

// type User struct {
// 	Username string `bson:"username" json:"username"`
// 	Password string `bson:"password" json:"password"`
// }

package models

type User struct {
	Password       string `json:"password" bson:"password"`
	Name           string `json:"name" bson:"name"`
	Email          string `json:"email" bson:"email"`
	SessionExpires string `json:"sessionExpires" bson:"sessionExpires"`
	SessionToken   string `json:"sessionToken" bson:"sessionToken"`
	IsAdmin        bool   `json:"isAdmin" bson:"isAdmin"`
}

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
