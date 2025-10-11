// Package main demonstrates Go struct usage with embedded types and pointer receivers.
//
// The code shows:
// - Struct embedding (contactInfo embedded in person)
// - Pointer receivers vs value receivers
// - Method definitions on custom types
//
// Key concepts:
// * (asterisk) - declares a pointer type or dereferences a pointer
// & (ampersand) - gets the address of a variable (address-of operator)
//
// Go shortcut: When calling methods with pointer receivers, Go automatically
// takes the address of the value. So jim.updateName() works even though
// updateName has a pointer receiver - Go converts it to (&jim).updateName()
// behind the scenes.
//
// Similarly, when you have a pointer and call a method with a value receiver,
// Go automatically dereferences it.
package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zip:   94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
