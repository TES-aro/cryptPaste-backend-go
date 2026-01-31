package dataBase

import(
	"database/sql"
)

const username string = "sp00n"
const password string = "salasana12"
const dbname string = "fanttiDB"

func connect() {
	dsn := username +":" + password + "@tcp(localhost:5432)/" + dbname

}

