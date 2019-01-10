package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("I am the parent")

	data, err := Asset("child")
	if err != nil {
		log.Fatalf("Asset child was not embeded: %s", err)
	}
	log.Printf("Asset child found with size %d", len(data))
}
