package main

import (
	"fmt"
	"os"

	"github.com/isidroevc/blist_api/facing"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file %s\n", err)
	}
	port := os.Getenv("PORT")
	app := facing.BuildApp()

	portText := fmt.Sprintf(":%s", port)
	fmt.Println(portText)
	app.Listen(portText)
}
