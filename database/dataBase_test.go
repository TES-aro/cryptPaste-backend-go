package dataBase

import (
	"os"
	"testing"
)


func TestGetEnvVar(t *testing.T){
	dbURL := os.Getenv("MSQL_URL")
	if len(dbURL) < 5 {
		t.Fatal("environment variable for postgres address not found")
	}
}
func TestCreation(t *testing.T){
	db, err := ConnectToDB("cryptTest")
	if err != nil {
		t.Fatal("couldn't connect to database. error: " + err.Error())
	}
	defer db.Close()
	err = InitTable(db)
	if err != nil {
		t.Fatal("couldn't init a table. error: " + err.Error())
	}
}

func TestPostAndFetch(t *testing.T){
	db, err := ConnectToDB("cryptTest")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer db.Close()

	id := "test1"
	err = PostDB(db, id, "kieli", "lorem ipsum")
	if err != nil {
		t.Fatal(err)
	}

	res , err1:= FetchDB(id, db)
	if err1 != nil {
		t.Fatal(err1)
	}
	if res.Content != "lorem ipsum" {
		t.Error("content doesn't match")
	}
	if res.Language != "kieli" {
		t.Error("language doesn't match")
	}

	err2 := PostDB(db, id, "uusi kieli", "ipsum lorem")
	newRes, err := FetchDB(id, db)
	if err2 == nil {
		t.Error("accepted post with same ID")
	}
	if res.Content != newRes.Content || res.Language != newRes.Language{
		t.Error("mutation")
	}

	wantErr := PostDB(db, "id2", "very long name for a language to test sql","content is fun")
	if wantErr == nil {
		t.Error(err.Error())
	}

	_, wantErr2 := FetchDB("id2", db)
	if wantErr2 == nil {
		t.Error("accepted too long language field nad managed to fetch")
	}
}


func TestDropTable(t *testing.T){
	db, err := ConnectToDB("cryptTest")
	if err != nil {
		t.Fatalf("couldn't connect to database")
	}
	defer db.Close()
	_,  err = bobbyDropTables(db)
	if err != nil {
		t.Error(err.Error())
	}
}
