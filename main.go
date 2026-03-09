package main

import (
	"connection"
	"dataBase"
	"fmt"
	"net/http"
	"os"
)
func main(){
	port := os.Getenv("CRYP_PORT")
	if port == ""{
		port = "1337"
	}
	db, err := dataBase.ConnectToDB("pasteCrypt")
	if db == nil {
		fmt.Println("couldn't connect to database.")
		return
	}
	dataBase.InitTable(db) 
	if err != nil {
		fmt.Println("Error connecting to database: " +err.Error())
	}
	defer db.Close()

	mux := connection.CreateMux("/api","/api/", db)
	fmt.Printf("running server on port %s\n", port)
	err = http.ListenAndServe(":"+port, mux)
	fmt.Println(err) 
}
