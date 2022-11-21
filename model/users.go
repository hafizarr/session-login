package model

type User struct {
	ID        string `form:"id" json:"id"`
	Username  string `form:"username" json:"username"`
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
	Password  string `form:"password" json:"passwoed"`
}
