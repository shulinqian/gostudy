package demoNet

import (
"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorillasss helloss hahhaa!\n"))
}

func Httpserver() {

	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.1:9999", nil)

}