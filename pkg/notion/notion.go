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

	client := notion.NewClient(apiKey)

	var allPages []notion.Page
	var cursor string

	for {
		params := &notion.DatabaseQuery{
			PageSize:    100,
			StartCursor: cursor,
		}

		response, err := client.QueryDatabase(context.Background(), dbID, params)
		if err != nil {
			fmt.Println("Erro ao consultar o banco de dados:", err)
			break
		}

		allPages = append(allPages, response.Results...)

		if !response.HasMore {
			break
		}

		cursor = *response.NextCursor
	}

	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}

	b, err := json.Marshal(allPages)
	if err != nil {
		log.Fatalf("Error loading database: %s", err)
	}
	os.WriteFile("dados.json", []byte(b), 0666)

}
