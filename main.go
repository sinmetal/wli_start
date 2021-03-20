package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/compute/metadata"
)

func main() {
	fmt.Println("Start WLI")

	ctx := context.Background()

	saEmail, err := metadata.Email("")
	if err != nil {
		log.Fatal(ctx, err.Error())
	}
	fmt.Printf("I am %s\n", saEmail)
}
