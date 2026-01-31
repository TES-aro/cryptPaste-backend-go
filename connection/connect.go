package connection

import (
	"dataBase"
	"fmt"
	"log"
	"net/http"
)

func Connect() {

    // API routes
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world")
    })
    http.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
	   	if req.Method == "POST" {
		   	fmt.Println("someone tried to post at /api")
		   	w.WriteHeader(200)
		   	fmt.Fprintf(w,"all good")
		   	return
	   	}
	   	if req.Method == "GET"{
	   		JsonResponse(w, req)
	   		return
	   	}
	   	w.WriteHeader(500)
	   	fmt.Fprint(w,"no suppor for operation")
    })

    port := ":6969"
    // Start server on port specified above
    log.Fatal(http.ListenAndServe(port, nil))
    fmt.Println("Server is running on port" + port)
}

func respondPost(content string) string {
	ID := dataBase.GenID()
	dataBase.WriteFile(content, "./data/"+ID)
	return ID
}

func respondGet(ID string) (string, error) {
	a, err := dataBase.ReadFile("./data/"+ID)
	if err != nil {
		return "", err
	}
	return a, nil

}
