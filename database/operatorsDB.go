package dataBase

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)



func PostDB(db *sql.DB, id, language, content string) error{
	query := `INSERT INTO pasteCrypt (id, language, content)
		VALUES (?, ?, ?)`
	_, err := db.ExecContext(context.Background(), query, id, language, content)
	return err
}

func FetchDB(id string, db *sql.DB) (Entry, error){
	query := "SELECT language, content FROM pasteCrypt WHERE id = ?"
	var language string
	var content string
	err := db.QueryRow(query, id).Scan(&language, &content)
	if err != nil {
		return Entry{}, err
	}
	return Entry{language, content }, err
}

func bobbyDropTables(db *sql.DB) (sql.Result, error) {
	query := "DROP TABLE pasteCrypt"
	msg , err := db.Exec(query)
	return msg, err
}
