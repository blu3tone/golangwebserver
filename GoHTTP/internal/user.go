package model

import "time"

type User struct{

	Base
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Email     string     `json:"email"`
	Mobile    string     `json:"mobile,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	Address   string     `json:"address,omitempty"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	Active    bool       `json:"active"`
	Token     string     `json:"-"`
	Role *Role `json:"role,omitempty"`

	RoleID     int `json:"-"`
	Books 	[]Book `json:"book, omitempty"`


}