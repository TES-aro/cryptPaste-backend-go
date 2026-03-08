package connection

import (
	"dataBase"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
)

const maxSize int = 655360

func pow10(exponent int) int {
	res := 1
	for i := 0; i < exponent; i++ {
		res = res * 10
	}
	return res
}
func genId() string {
	id := 0
	for i:= 0; i<16; i++{
		id = id + (rand.IntN(10) * pow10(i))
	}
	return strconv.Itoa(id)
}


func fetchContent(ID string, dbPointer *sql.DB) ([]byte, error) {
  content , err:= dataBase.FetchDB(ID, dbPointer)

	json , err := json.Marshal(JSONpaste{content.Language, content.Content })
	if err != nil {
		return nil, err
	}
	return json, nil
}

func saveContent(content JSONpaste, dbPointer *sql.DB) ([]byte, error) {
	id := genId()
	err := dataBase.PostDB(dbPointer, id, content.Language, content.Text)
	json , err:= json.Marshal(JSONid{id})
	if err != nil {
		return nil, err
	}
	return json, nil
}

type JSONpaste struct {
	Language string `json:"language"`
	Text string `json:"text"`
}

type JSONid struct {
	ID string `json:"id"`
}

//###

type GEThandler struct{
	url string
	dbPointer *sql.DB
}

func (h *GEThandler) ServeHTTP (w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}
	ID := strings.TrimPrefix(req.URL.Path, h.url)
	if len(ID) != 16 {
		w.WriteHeader(404)
		return
	}
	content , err:= fetchContent(ID, h.dbPointer)
	if err != nil {
		fmt.Println(err) 
	}
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, content)
}


type POSThandler struct{
	dbPointer *sql.DB
}

func (h *POSThandler) ServeHTTP (w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Error: method not supported", 405)
		return
	}
	typeHeader := req.Header.Get("Content-Type")
	fmt.Println("header type: " + typeHeader)
	if typeHeader != "application/json" {
		http.Error(w,"Error: improper Content-Type header", 400)
		return
	}
	contentLen := req.ContentLength
	fmt.Print("content lenght: ")
	fmt.Println(contentLen)
	if contentLen > int64(maxSize) {
		http.Error(w, "Error: content too large", 400)
		return
	}

	// reading the content
	defer req.Body.Close()
	req.Body = http.MaxBytesReader(w,req.Body,int64(maxSize))

	var paste JSONpaste

	err := json.NewDecoder(req.Body).Decode(&paste)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
	}

	if len(paste.Language) > 50 {
		http.Error(w, "Error: misshapen JSON",  400)
		return
	}

	ID , err:= saveContent(paste, h.dbPointer)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprint(w, ID)
}
