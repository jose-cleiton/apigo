package models

import "Api/db"



func Insert(this Todo) (id int64, err error ){

	conn, err:=db.OpenConnection()

	if err!= nil {
		return
	}

	defer conn.Close()

	sqlStatement := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sqlStatement, this.Title, this.Description, this.Done).Scan(&id)

	return
}