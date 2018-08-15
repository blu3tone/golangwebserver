package model

import "time"

type Author struct {
	Base
	Name  string  `json: "name"`
	Phone string  `json: "phone"`
	Address string `json: "address"`
	BirthDay *time.Time `json: "birthday"`
	Books []*Book `json: "books, omitempty"`
}