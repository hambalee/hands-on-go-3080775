// types/maps/lookups/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	var authors = map[string]author{
		"tm": {name: "Toni Morrison"},
		"ma": {name: "Marcus Aurelius"},
	}

	// read the value for a non-existent key
	fmt.Println("JR: ", authors["jr"])

	// check when a key is present in the map
	// fmt.Println(authors["jr"])
	// a, ok := authors["tm"] // ok => true
	a, ok := authors["jr"] // ok => false
	fmt.Printf("a = %v, ok = %v\n", a, ok)
}
