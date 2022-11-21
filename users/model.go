package users

import "html/template"

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

type UserRegister struct {
	Username  string `form:"username" json:"username"`
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
	Password  string `form:"password" json:"passwoed"`
}

type UserLogin struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"passwoed"`
}
