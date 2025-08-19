package main

import (
	"log"

	"github.com/mechiko/utility"
)

func main() {
	code, err := utility.Sscc("46164463019900001")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(code)
}
