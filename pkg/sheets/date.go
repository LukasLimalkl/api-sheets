package sheets

import "time"

func formatDate(dateString string) string {
	// Converter a string de data para o formato de data e hora do Go
	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return "" // Retornar uma string vazia em caso de erro
	}

	// Formatar a data no formato desejado
	return date.Format("2006-01-02 15:04:05")
}
