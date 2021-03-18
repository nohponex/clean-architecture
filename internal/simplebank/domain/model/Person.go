package model

type (
	PersonID string
	 Person struct {
		 Username string
	 }
)

func NewPerson(username string) *Person {
	return &Person{Username: username}
}
