package main

// Structs vs Maps in Go:
//
// Maps (used here):
//   - Dynamic collection of key-value pairs (both keys and values must be same type)
//   - Keys can be added/removed at runtime
//   - Use when: you don't know all keys ahead of time, need flexible key-value storage,
//     or require fast lookups by key
//
// Structs (seen in previous examples):
//   - Fixed set of fields defined at compile time, each field can be different type
//   - Fields are known and accessed by name (not by key lookup)
//   - Use when: representing a specific entity with known properties, need type safety,
//     or want to define methods on the type
//
// Example: Use struct for a Person{name string, age int}, use map for color codes where
// colors are dynamic.

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		println("Hex code for", color, "is", hex)
	}
}
