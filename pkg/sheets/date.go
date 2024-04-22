package sheets

import "time"

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
