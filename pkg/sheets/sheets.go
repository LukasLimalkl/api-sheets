package sheets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func ConnectSheets() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	ctx := context.Background()

	config := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	CLIENT_TOKEN := os.Getenv("GOOGLE_CLIENT_SECRET")

	token, err := config.Exchange(ctx, CLIENT_TOKEN)

	sheetKey := os.Getenv("KEY_SHEETS")
	sheetsService, err := sheets.NewService(ctx, token)
	if err != nil {
		fmt.Println(err)
	}

	data, err := os.ReadFile("./dados.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %s", err)
	}

	var records []map[string]interface{}
	if err := json.Unmarshal(data, &records); err != nil {
		log.Fatalf("Error decoding JSON: %s", err)
	}

	var vr sheets.ValueRange
	for _, record := range records {
		var row []interface{}
		for _, value := range record {
			row = append(row, value)
		}
		vr.Values = append(vr.Values, row)
	}

	dataRange := "A1"

	_, err = sheetsService.Spreadsheets.Values.Append(sheetKey, dataRange, &vr).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("Unable to write data to sheet: %v", err)
	}

}
