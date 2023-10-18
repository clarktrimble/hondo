package main

import (
	"fmt"

	"github.com/clarktrimble/hondo"
)

func main() {

	// call Seed for deterministic hondo as one might want for unit testing:
	// rand.Seed(1) //nolint:staticcheck // ok for simple unit

	sessionId := hondo.Rand(7)
	fmt.Printf("Your session id will be: %s\n", sessionId)
}
