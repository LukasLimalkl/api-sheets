package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
)

func ConnectNotion() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	apiKey := os.Getenv("API_NOTION")
	dbID := os.Getenv("DB_NOTION")

	client := notionapi.NewClient(notionapi.Token(apiKey))

	pages, err := client.Database.Get(context.Background(), notionapi.DatabaseID(dbID))
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
