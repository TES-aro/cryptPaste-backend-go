package connection

import(
	"net/http"
)


func createMux(postURL, getURL string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(getURL, &GEThandler{getURL})
	mux.Handle(postURL, &POSThandler{})
	return mux
}
