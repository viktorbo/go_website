package main

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hmm, mb it's easy")
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts")
}

func HandleRequest() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts", ContactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	HandleRequest()
}
