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

	// First, list existing branches
	fmt.Println("\nListing existing branches:")
	if err := listBranches(client); err != nil {
		log.Fatalf("Failed to list branches: %v", err)
	}

	// Get main branch hash
	// mainBranch := "main"
	// main, err := client.GetBranch(mainBranch)
	// if err != nil {
	// 	log.Fatalf("Failed to get main branch: %v", err)
	// }

	// // Define namespace and table names
	// namespace := "test-namespace-1"
	// tableName := "iceberg-table-2"

	// Create a namespace
	// fmt.Printf("\nCreating namespace '%s'...\n", namespace)
	// namespaceOps := []models.Operation{
	// 	{
	// 		Type: "PUT",
	// 		Key: models.ContentKey{
	// 			Elements: []string{namespace},
	// 		},
	// 		Content: &models.Content{
	// 			Type: "NAMESPACE",
	// 		},
	// 	},
	// }

	// // Commit namespace creation to main
	// namespaceResponse, err := client.CommitChange(mainBranch, main.Hash, namespaceOps, fmt.Sprintf("Create namespace %s", namespace))
	// if err != nil {
	// 	log.Fatalf("Failed to create namespace: %v", err)
	// }
	// main = &namespaceResponse.TargetBranch
	// fmt.Printf("Namespace created in main (hash: %s)\n", main.Hash)

	// Create an Iceberg table in the namespace
	// fmt.Printf("\nCreating Iceberg table '%s.%s'...\n", namespace, tableName)
	// tableOps := []models.Operation{
	// 	{
	// 		Type: "PUT",
	// 		Key: models.ContentKey{
	// 			Elements: []string{namespace, tableName},
	// 		},
	// 		Content: &models.Content{
	// 			Type:             "ICEBERG_TABLE",
	// 			MetadataLocation: "/path/to/metadata/location",
	// 			SnapshotId:       1,
	// 			SchemaId:         1,
	// 			SpecId:           1,
	// 			SortOrderId:      1,
	// 		},
	// 	},
	// }

	// // Commit table creation to main
	// tableResponse, err := client.CommitChange(mainBranch, main.Hash, tableOps, fmt.Sprintf("Create Iceberg table %s.%s", namespace, tableName))
	// if err != nil {
	// 	log.Fatalf("Failed to create table: %v", err)
	// }
	// main = &tableResponse.TargetBranch
	// fmt.Printf("Table created in main (hash: %s)\n", main.Hash)

	// Get the content ID from the commit response
	// var tableContentId string
	// for _, content := range tableResponse.AddedContents {
	// 	if len(content.Key.Elements) == 2 &&
	// 		content.Key.Elements[0] == namespace &&
	// 		content.Key.Elements[1] == tableName {
	// 		tableContentId = content.ContentId
	// 		break
	// 	}
	// }
	// fmt.Printf("Table content ID: %s\n", tableContentId)

	// Create a feature branch
	// featureBranch := "feature-branch-1"
	// fmt.Printf("\nCreating feature branch '%s' from '%s'...\n", featureBranch, mainBranch)
	// newBranch, err := client.CreateBranch(featureBranch, mainBranch)
	// if err != nil {
	// 	log.Fatalf("Failed to create branch: %v", err)
	// }
	// fmt.Printf("Created feature branch: %s (hash: %s)\n", newBranch.Name, newBranch.Hash)

	// Update the table in the feature branch
// 	fmt.Printf("\nUpdating table ...\n")
// 	updateOps := []models.Operation{
// 		{
// 			Type: "PUT",
// 			Key: models.ContentKey{
// 				Elements: []string{namespace, tableName},
// 			},
// 			Content: &models.Content{
// 				Type:             "ICEBERG_TABLE",
// 				ContentId:        tableContentId,
// 				MetadataLocation: "/path/to/metadata/location/v2",
// 				SnapshotId:       2,
// 				SchemaId:         1,
// 				SpecId:           1,
// 				SortOrderId:      1,
// 			},
// 		},
// 	}

// 	// Commit table update to feature branch
// 	updateResponse, err := client.CommitChange(featureBranch, newBranch.Hash, updateOps, fmt.Sprintf("Update Iceberg table %s.%s", namespace, tableName))
// 	if err != nil {
// 		log.Fatalf("Failed to update table: %v", err)
// 	}
// 	fmt.Printf("Table updated in feature branch (hash: %s)\n", updateResponse.TargetBranch.Hash)

// 	// List final branch state
	// fmt.Println("\nFinal branch state:")
	// if err := listBranches(client); err != nil {
	// 	log.Fatalf("Failed to list branches: %v", err)
	// }
	// res := api.Dummy()
	// fmt.Println(res)
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
