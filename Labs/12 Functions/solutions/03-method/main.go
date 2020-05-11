package main

import "fmt"

type User struct {
	First string
	Last  string
	Email string
}

func (u User) FullName() string {
	return u.First + " " + u.Last
}

func (u User) FormattedEmail() string {
	return fmt.Sprintf("%s %s <%s>", u.First, u.Last, u.Email)
}

func (u *User) Reset() {
	u.First = ""
	u.Last = ""
	u.Email = ""
}
func main() {
	u := User{"Rob", "Pike", "commander@google.com"}
	fmt.Printf("%+v\n", u)
	fmt.Println("Full Name", u.FullName())
	fmt.Println("Formatted Email", u.FormattedEmail())
	u.Reset()
	fmt.Printf("%+v\n", u)
}
