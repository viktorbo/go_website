package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name       string
	age        uint16
	money      int16
	avg_grades float64
	happiness  float32
}

func (u User) GetUserInfo() string {
	return fmt.Sprintf("User name: %s\n"+
		"Age: %d \n"+
		"Money: %d \n", u.name, u.age, u.money)
}

func (u *User) SetUserName(name string) {
	u.name = name
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hmm, mb it's easy")
	// var Bob User = ...
	// Bob := User{name: "Bob", age: 25, money: -50, happiness: 0.8, avg_grades: 4.2}
	Bob := User{"Bob", 25, -50, 4.2, 0.8}
	fmt.Fprintf(w, "User name is: "+Bob.name+"\n")
	Bob.name = "Not Bob"
	fmt.Fprintf(w, "User name is: "+Bob.name)
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts \n\n")

	Bob := User{"Bob", 25, -50, 4.2, 0.8}
	fmt.Fprintln(w, Bob.GetUserInfo())

	Bob.SetUserName("Alex")
	fmt.Fprintln(w, Bob.name)
}

func HandleRequest() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts", ContactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	HandleRequest()
}
