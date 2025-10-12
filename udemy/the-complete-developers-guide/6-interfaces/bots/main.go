package main

// Interfaces in Go:
// - Define a contract (set of methods) that types must implement
// - Implementation is IMPLICIT - no "implements" keyword needed
// - Any type that has the required methods automatically satisfies the interface
// - Enable polymorphism: functions can accept any type that satisfies the interface
// - Are satisfied by concrete types, not by embedding or inheritance
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// Ver custom logic for generating an spanish greeting
	return "Hola!"
}
