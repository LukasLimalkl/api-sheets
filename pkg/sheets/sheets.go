package sheets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Item struct {
	ID      string `json:"id"`
	Created string `json:"created_time"`
	Edited  string `json:"last_edited_time"`
	Parent  struct {
		Type       string `json:"type"`
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Archived bool   `json:"archived"`
	URL      string `json:"url"`
	Icon     struct {
		Type  string `json:"type"`
		Emoji string `json:"emoji"`
	} `json:"icon"`
	Cover struct {
		Type string `json:"type"`
		File struct {
			URL        string `json:"url"`
			ExpiryTime string `json:"expiry_time"`
		} `json:"file"`
	} `json:"cover"`
	Properties struct {
		Campanha         Relation    `json:"Campanha"`
		Card             UniqueID    `json:"Card"`
		Bm               Select      `json:"Bm"`
		CardCriativo     Relation    `json:"Card Criativo"`
		Cliente          Relation    `json:"Cliente"`
		CriadoEm         interface{} `json:"Criado em"`
		DataDePostagem   Date        `json:"Data de Postagem"`
		Demanda          Select      `json:"Demanda"`
		Departamento     Select      `json:"Departamento"`
		Fabricacao       Date        `json:"Fabricação"`
		FluxoCinema      Status      `json:"Fluxo - Cinema"`
		FluxoDesigner    Status      `json:"Fluxo - Designer"`
		Post             Title       `json:"Post"`
		Responsavel      Select      `json:"Responsável"`
		StatusDoCard     Status      `json:"Status do Card"`
		TipoDoEntregavel Select      `json:"Tipo do Entregável"`
		UltimaEdicao     interface{} `json:"Última edição"`
	} `json:"properties"`
}

type Title struct {
	Text string `json:"text"`
}

type Relation struct {
	ID string `json:"id"`
}

type UniqueID struct {
	Prefix string `json:"prefix"`
	Number string `json:"number"`
}

type Select struct {
	Name string `json:"name"`
}

type Status struct {
	Name string `json:"name"`
}

type Date struct {
	Start string `json:"start"`
}

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

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

	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	headerToColumn := map[string]int{
		"Card Number":        0,
		"BM":                 1,
		"Post":               2,
		"Responsável":        3,
		"Departamento":       4,
		"Fluxo - Cinema":     5,
		"Fluxo - Designer":   6,
		"Status do Card":     7,
		"Tipo do Entregável": 8,
		"Demanda":            9,
		"Data de Postagem":   10,
		"Fabricação":         11,
		"Última Edição":      12,
		"Criado Em":          13,
	}

	var values [][]interface{}
	for _, item := range items {
		row := make([]interface{}, len(headerToColumn))
		for header, columnIndex := range headerToColumn {
			switch header {
			case "Card Number":
				row[columnIndex] = item.Properties.Card.Prefix + item.Properties.Card.Number
			case "BM":
				row[columnIndex] = item.Properties.Bm.Name
			case "Post":
				row[columnIndex] = item.Properties.Post.Text
			case "Responsável":
				row[columnIndex] = item.Properties.Responsavel.Name
			case "Departamento":
				row[columnIndex] = item.Properties.Departamento.Name
			case "Fluxo - Cinema":
				row[columnIndex] = item.Properties.FluxoCinema.Name
			case "Fluxo - Designer":
				row[columnIndex] = item.Properties.FluxoDesigner.Name
			case "Status do Card":
				row[columnIndex] = item.Properties.StatusDoCard.Name
			case "Tipo do Entregável":
				row[columnIndex] = item.Properties.TipoDoEntregavel.Name
			case "Demanda":
				row[columnIndex] = item.Properties.Demanda.Name
			case "Data de Postagem":
				row[columnIndex] = formatDate(item.Properties.DataDePostagem.Start)
			case "Fabricação":
				row[columnIndex] = formatDate(item.Properties.Fabricacao.Start)
			case "Última Edição":
				row[columnIndex] = formatInterface(item.Properties.UltimaEdicao)
			case "Criado Em":
				row[columnIndex] = item.Created
			}
		}
		values = append(values, row)
	}

	spreadsheetId := "1lDmBM9Aap2l1MGMjyFgilSwk5lzrbC7TjR3DvZPLH48"
	writeRange := "teste!A1"

	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, &sheets.ValueRange{
		Values: values,
	}).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("Não foi possível escrever na planilha: %v", err)
	}

	fmt.Println("Dados inseridos com sucesso na planilha!")
}

func formatDate(dateString string) string {
	if dateString == "" {
		return ""
	}
	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return ""
	}
	return date.Format("2006-01-02")
}

func formatInterface(value interface{}) string {
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case string:
		return v
	case time.Time:
		return v.Format("02/01/2006 15:04:05")
	case map[string]interface{}:
		if stringValue, ok := v["string_value"].(string); ok {
			return stringValue
		}
	default:
		return fmt.Sprintf("%v", v)
	}
	return ""
}
