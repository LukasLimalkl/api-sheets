package sheets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func ConnectSheets() {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	data, err := os.ReadFile("./dados.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %s", err)
	}

	var items Items
	if err := json.Unmarshal(data, &items); err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	spreadsheetId := "1lDmBM9Aap2l1MGMjyFgilSwk5lzrbC7TjR3DvZPLH48"

	var values [][]interface{}
	for _, item := range items {
		cripto := item.Properties.Card.ID
		decoded, err := url.QueryUnescape(cripto)
		if err != nil {
			fmt.Println("Erro ao decodificar:", err)
			return
		}

		row := []interface{}{
			decoded,
			item.Properties.Bm.Select.Name,
			item.Properties.DataDePostagem.Date.Start,
		}
		values = append(values, row)
	}

	sheetid := 0

	response1, err := srv.Spreadsheets.Get(spreadsheetId).Fields("sheets(properties(sheetId,title))").Do()
	if err != nil || response1.HTTPStatusCode != 200 {
		log.Fatalf("Unable to retrieve sheet: %v", err)
	}
	sheetName := ""
	for _, v := range response1.Sheets {
		prop := v.Properties
		sheetID := prop.SheetId
		if sheetID == int64(sheetid) {
			sheetName = prop.Title
			break
		}
	}

	valueInputOption := "USER_ENTERED"
	insertDataOption := "INSERT_ROWS"
	rb := &sheets.ValueRange{
		Values: values,
	}
	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, sheetName, rb).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to append values: %v", err)
	}

	fmt.Println("Dados inseridos com sucesso na planilha!")
}
