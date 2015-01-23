package main

import (
	"net/http"
)
var last = "-correct-brightgreen.svg"

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if last == "-correct-brightgreen.svg" {
		last = "-wrong-red.svg"
	} else {
		last = "-correct-brightgreen.svg"
	}
	w.Header().Set("Cache-Control", "no-cache")
	http.Redirect(w, r, "https://img.shields.io/badge/solution"+last, 302)
	return
}
 

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}

