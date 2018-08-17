package demo

import (
"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorillasss helloss!\n"))
}

func Httpserver() {

	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.1:8099", nil)

}