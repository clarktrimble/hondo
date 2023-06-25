package main

import (
	"fmt"

	"github.com/clarktrimble/hondo"
)

func main() {

	sessionId := hondo.Rand(7)
	fmt.Printf("Your session id will be: %s\n", sessionId)
}
