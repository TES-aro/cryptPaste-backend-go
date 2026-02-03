package connection

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const maxSize int = 655360


func fetchContent(ID string) Paste {
	json := Paste{"Some language", "lorem ipsum" }
	return json
}

func saveContent(content Paste) string {
	return "NewIDisNotThis"
}

type Paste struct {
	Language string `json:"language"`
	Text string `json:"text"`
}

type GEThandler struct{
	url string
}

func (h *GEThandler) ServeHTTP (w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}
	ID := strings.TrimPrefix(req.URL.Path, h.url)
	if len(ID) != 8 {
		w.WriteHeader(404)
		return
	}
	content := fetchContent(ID)
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, content)
}


type POSThandler struct{}

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

	var paste Paste

	err := json.NewDecoder(req.Body).Decode(&paste)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
	}

	if len(paste.Language) > 50 {
		http.Error(w, "Error: mishapen JSON",  400)
		return
	}

	ID := saveContent(paste)

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprint(w, ID)
}
