package notion

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dstotijn/go-notion"
	"github.com/joho/godotenv"
)

func ConnectNotion() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	API_NOTION := os.Getenv("API_NOTION")
	DB_NOTION := os.Getenv("DB_NOTION")

	client := notion.NewClient(API_NOTION)

	page, err := client.FindDatabaseByID(context.Background(), DB_NOTION)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(page)
}
