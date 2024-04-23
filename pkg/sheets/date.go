package sheets

import "time"

func formatDateProperty(dateStr string) string {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return ""
	}
	t = t.AddDate(0, 0, 1)
	loc, _ := time.LoadLocation("America/Araguaina")
	return t.In(loc).Format("02/01/2006")
}

func formatTimeProperty(timeStr string) string {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	loc, _ := time.LoadLocation("America/Araguaina")
	return t.In(loc).Format("02/01/2006 15:04:05")
}
