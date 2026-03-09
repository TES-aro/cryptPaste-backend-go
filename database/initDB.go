package dataBase

import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"os"
	"errors"
)

func ConnectToDB(DBname string) (*sql.DB, error) {
	dbURL := os.Getenv("MSQL_URL")
	if len(dbURL) < 5 {
		return nil, errors.New("no environmental variable for database address")
	}
	dbURL = dbURL + "/" + DBname
	db ,err:= sql.Open("mysql", dbURL)
	return db, err
}

func InitTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS pasteCrypt (
		id varchar(16) NOT NULL,
		language varchar(32),
		content text NOT NULL,
		created timestamp DEFAULT NOW(),
		PRIMARY KEY (id)
	)`
	_, err :=db.Exec(query)
	return err
}
