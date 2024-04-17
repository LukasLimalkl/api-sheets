package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dstotijn/go-notion"
	"github.com/joho/godotenv"
)

func ConnectNotion() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	apiKey := os.Getenv("API_NOTION")
	dbID := os.Getenv("DB_NOTION")

	f := &notion.DatabaseQuery{
		PageSize: 15,
	}

	client := notion.NewClient(apiKey)

	pages, err := client.QueryDatabase(context.Background(), dbID, f)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}

	b, err := json.Marshal(pages)
	if err != nil {
		log.Fatalf("Error loading database: %s", err)
	}
	os.Stdout.Write(b)

}
