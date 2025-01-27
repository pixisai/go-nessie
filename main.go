package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


// go-nessie/
// │── main.go
// ├── pkg/
// │   ├── api/
// │   │   ├── branch.go
// │   │   ├── delete.go
// │   │   └── fetch.go
// │   └── utils/
// │   │   └── http_client.go
// │   └── models(types)/
// │       └── request/
// │       		└── branch.go
// │       		└── branch.go
// │       └── post.go
// │       └── delete.go
// ├── config/
// │   └── config.go
// ├── go.mod
// ├── go.sum
// └── README.md

func main() {
	fmt.Println("Nessie Started")

	// Create a new Nessie client
	client := nessie.NewClient("http://0.0.0.0:19120")

	// Get all trees
	trees, err := client.GetAllTrees()
	if err != nil {
		log.Fatalf("Failed to get trees: %v", err)
	}

	fmt.Println("All trees:")
	for _, ref := range trees.References {
		fmt.Printf("- %s (%s): %s\n", ref.Name, ref.Type, ref.Hash)
	}

}
