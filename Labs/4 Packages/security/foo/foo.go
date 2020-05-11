// Located in $GOPATH/src/github.com/gopherguides/foo
package foo

type User struct {
	First    string
	Last     string
	password string
}

// NewUser returns a user with the intialized values provided
func NewUser(first, last, password string) User {
	return User{
		First:    first,
		Last:     last,
		password: password,
	}
}
