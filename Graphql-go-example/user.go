package main


type User struct{
	ID			int
	Username 	string
	Password	string
}


func GetUserByID(id int)	(*User, error){
	var username string

	err := db.QueryRow("SELECT username FROM users WHERE id=$1", id).Scan(&username)
	if err != nil {
		return nil, err
	}

	return &User{
		ID: id,
		Username: username,
		Password: "",
	}, nil
}

func InsertUser(user *User) error{
	var id int
	err := db.QueryRow(`
		INSERT INTO users(username, password)
		VALUES ($1, $2)
		RETURNING id
	`, user.Username, user.Password).Scan(&id)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}


func RemoveUserByID(id int) error{
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}