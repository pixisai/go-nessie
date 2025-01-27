package main

import (
	"fmt"
	"log"
	"github.com/pixisai/go-nessie/api"
)

func main() {
	fmt.Println("Nessie Started")

	// Create a new Nessie client
	client := api.NewClient("http://0.0.0.0:19120")

	// Get all branches
	branches, err := client.GetAllBranches()
	if err != nil {
		log.Fatalf("Failed to get branches: %v", err)
	}

	fmt.Println("All branches:")
	for _, ref := range branches.References {
		fmt.Printf("- %s (%s): %s\n", ref.Name, ref.Type, ref.Hash)
	}
}
