package main

import (
	"github.com/LukasLimalkl/api-sheets/pkg/notion"
	"github.com/LukasLimalkl/api-sheets/pkg/sheets"
)

func main() {
	notion.ConnectNotion()
	sheets.ConnectSheets()

}
