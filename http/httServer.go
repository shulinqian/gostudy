package http

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorillasss!\n"))
}

func main() {

	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)

}