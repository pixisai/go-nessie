package main

import (
	"fmt"
	"log"
	"os"
	// "time"

	"github.com/pixisai/go-nessie/api"
	"github.com/pixisai/go-nessie/models"
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

	// Create a new branch from main
	sourceBranch := "main"
	newBranchName := "feature-branch-v1.3"

	fmt.Printf("\nCreating new branch '%s' from '%s'...\n", newBranchName, sourceBranch)
	newBranch, err := client.CreateBranch(newBranchName, sourceBranch)
	if err != nil {
		log.Fatalf("Failed to create branch: %v", err)
	}
	fmt.Printf("Created new branch: %s (hash: %s)\n", newBranch.Name, newBranch.Hash)

	// List branches again to see our new branch
	fmt.Println("\nListing branches after creation:")
	if err := listBranches(client); err != nil {
		log.Fatalf("Failed to list branches: %v", err)
	}

	// Get specific branch info
	fmt.Printf("\nGetting details for branch '%s':\n", newBranchName)
	branch, err := client.GetBranch(newBranchName)
	if err != nil {
		log.Fatalf("Failed to get branch details: %v", err)
	}
	fmt.Printf("Branch details: %s (type: %s, hash: %s)\n", branch.Name, branch.Type, branch.Hash)

	// Delete the branch we created
	fmt.Printf("\nDeleting branch '%s'...\n", newBranchName)
	if err := deleteBranch(client, newBranchName, newBranch.Hash); err != nil {
		log.Fatalf("Failed to delete branch: %v", err)
	}
	fmt.Printf("Successfully deleted branch: %s\n", newBranchName)

	// List branches one final time to confirm deletion
	fmt.Println("\nListing branches after deletion:")
	if err := listBranches(client); err != nil {
		log.Fatalf("Failed to list branches: %v", err)
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

func deleteBranch(client *api.Client, name, hash string) error {
	deleteReq := &models.DeleteReferenceRequest{
		Name: name,
		Hash: hash, // Optional: specify hash to ensure we're deleting the right version
	}
	return client.DeleteBranch(deleteReq)
}
