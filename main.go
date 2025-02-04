package main

import (
	"fmt"
	"log"
	"os"

	// "time"
	"github.com/pixisai/go-nessie/api"
	// "github.com/pixisai/go-nessie/models"
)

func main() {
	// Get Nessie server URL from environment variable or use default
	serverURL := os.Getenv("NESSIE_SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:19120"
	}

	fmt.Printf("Connecting to Nessie server at %s\n", serverURL)

	// Create a new Nessie client
	client := api.NewClient(serverURL)

	// List existing branches
	fmt.Println("\nListing existing branches:")
	if err := listBranches(client); err != nil {
		log.Fatalf("Failed to list branches: %v", err)
	}

	tableDetails, err := client.GetTableDetails("main", "my_namespace_1", "table_1")
	if err != nil {
		log.Fatalf("Failed to get table details: %v", err)
	}

	// Access table metadata
	fmt.Printf("Table Schema ID: %d\n", tableDetails.Table.SchemaId)
	fmt.Printf("Table Spec ID: %d\n", tableDetails.Table.SpecId)
	fmt.Printf("Table Schema Fields:\n")
	for _, field := range tableDetails.Table.Schema.Fields {
		fmt.Printf("  - %s (%s)\n", field.Name, field.Type)
	}

}

func listBranches(client *api.Client) error {
	branches, err := client.GetAllBranches()
	if err != nil {
		return err
	}
	for _, ref := range branches.References {
		fmt.Printf("- %s (%s): %s\n", ref.Name, ref.Type, ref.Hash)
	}
	return nil
}
