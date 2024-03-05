package main

import (
	"fmt"
	applicationbase "nalanda_backend/applicationBase"
)

// The main entrypoint function for our application.
func main() {
	fmt.Println("Nalanda: An Open Society for Academics")

	// Start the entrypoint for GIN API server.
	applicationbase.ServerEntrypoint()
}
