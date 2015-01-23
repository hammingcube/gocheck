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
	http.Redirect(w, r, "https://img.shields.io/badge/solution"+last, http.StatusFound)
	return
}
 

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8092", nil)
}

