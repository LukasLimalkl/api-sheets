package sheets

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func ConnectSheets() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	apiKey := os.Getenv("API_SHEETS")
	ctx := context.Background()
	sheetsService, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sheetsService)
}
