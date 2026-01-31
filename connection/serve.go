package connection

import(
	"net/http"
	"encoding/json"
	"time"
)

type jsonResponse struct{
	Time string `json:"time"`
	Stuff string `json:"stuff"`
}

func JsonResponse(w http.ResponseWriter, r *http.Request) {
	time := time.Now().String()
	response := jsonResponse{time, "lorem ipsum" }
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response)
	err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



