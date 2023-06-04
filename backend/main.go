package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/menasehk13/ResturnatSystem/backend/config"
	"github.com/menasehk13/ResturnatSystem/backend/routes"
)


const (
	author  = "menasehk"
	version = "1.0.0"
)

func main() {

	err := config.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	
	router := routes.SetUpRoutes()
	fmt.Printf("Connected to the database\n")
	fmt.Println("Server is running on port 6000")
	fmt.Printf("Author: %s\n", author)
	fmt.Printf("Version: %s\n", version)
    log.Fatal(http.ListenAndServe(":6000", router))
}




