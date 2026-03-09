package connection

import(
	"net/http"
	"database/sql"
)


func CreateMux(postURL, getURL string, sqlPointer *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(getURL, &GEThandler{getURL, sqlPointer})
	mux.Handle(postURL, &POSThandler{sqlPointer})
	return mux
}
