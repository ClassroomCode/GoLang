package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID    int
	First string
	Last  string
	Email string
}

func (u User) String() string {
	return fmt.Sprintf(`{"id":%d","first":"%s","last":"%s"}`, u.ID, u.First, u.Last)
}

type authUser User

func (u authUser) String() string {
	return fmt.Sprintf(`{"id":%d","first":"%s","last":"%s","email":"%s"}`, u.ID, u.First, u.Last, u.Email)
}

func handler(w http.ResponseWriter, r *http.Request) {
	u := User{
		ID:    1,
		First: "Cory",
		Last:  "LaNou",
		Email: "cory@gopherguides.com",
	}
	w.Header().Set("Content-Type", "application/text")

	data := u.String()

	username, password, ok := r.BasicAuth()
	if ok && username == "admin" && password == "password" {
		// Type convert user to an authUser to send all protected information to be serialized in the response.
		data = authUser(u).String()
	}

	w.Write([]byte(data))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*

// section: command-anon
$ curl localhost:8080
{"id":1","first":"Cory","last":"LaNou"}
// section: command-anon

// section: command-auth
$ curl --user admin:password localhost:8080
{"id":1","first":"Cory","last":"LaNou","email":"cory@gopherguides.com"}
// section: command-auth

*/
