package main

import (
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.google.com", http.StatusFound)
	return
}
 

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8092", nil)
}

