package main


type Genre struct{
	ID		int
	Name 	string
	Rate 	int
}


func InsertGenre( genre *Genre) error{
	var id int
	err := db.QueryRow(`
		INSERT INTO genres(name, rate)
		VALUES ($1, $2)
		RETURNING id
	`, genre.Name, genre.Rate).Scan(&id)
	if err != nil {
		return err
	}
	genre.ID = id
	return nil
}