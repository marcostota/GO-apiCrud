package models

import "github.com/marcostota/apicrud/db"

func Insert(todo Todo) (id int64, err error) {
	con, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer con.Close()

	sql := `INSERT INTO todos (title, description, done)VALUES ($1,$2,$3) RETURNING id`

	err = con.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)
	return
}
