package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Email string `json:"-"`
}

type authUser User

func (a authUser) MarshalJSON() ([]byte, error) {
	// declare a private type to write our data to
	type usr struct {
		ID    int    `json:"id"`
		First string `json:"first,omitempty"`
		Last  string `json:"last,omitempty"`
		Email string `json:"email,omitempty"`
	}

	// Populate our new instance with just the data we need
	us := usr{
		ID:    a.ID,
		First: a.First,
		Last:  a.Last,
		Email: a.Email,
	}

	return json.Marshal(us)

}

func handler(w http.ResponseWriter, r *http.Request) {
	u := User{
		ID:    1,
		First: "Cory",
		Last:  "LaNou",
		Email: "cory@gopherguides.com",
	}
	w.Header().Set("Content-Type", "application/text")

	username, password, ok := r.BasicAuth()
	if ok && username == "admin" && password == "password" {
		// Type convert user to an authUser to send all protected information to be serialized in the response.
		json.NewEncoder(w).Encode(authUser(u))
		return
	}

	json.NewEncoder(w).Encode(u)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*

// section: command-anon
$ curl localhost:8080
{"id":1,"first":"Cory","last":"LaNou"}
// section: command-anon

// section: command-auth
$ curl --user admin:password localhost:8080
{"id":1,"first":"Cory","last":"LaNou","email":"cory@gopherguides.com"}
// section: command-auth

*/
