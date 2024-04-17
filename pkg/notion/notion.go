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

	filter := &notion.DatabaseQueryFilter{
		Property: "campanha",
	}

	f := &notion.DatabaseQuery{
		Filter: filter,
	}
	client := notion.NewClient(apiKey)

	pages, err := client.QueryDatabase(context.Background(), dbID, f)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}

	b, err := json.Marshal(pages)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}
