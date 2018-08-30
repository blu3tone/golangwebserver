package main

import "time"

type Author struct{
	ID		int
	Name 	string
	Birthday 	time.Time
}


func InsertAuthor( author *Author) error{
	var id int
	err := db.QueryRow(`
		INSERT INTO authors	(name)
		VALUES ($1)
		RETURNING id
	`, author.Name).Scan(&id)
	if err != nil {
		return err
	}
	author.ID = id
	return nil
}

func RemoveAuthorByID(id int) error{
	_, err := db.Exec("DELETE FROM authors WHERE id=$1", id)
	return err
}