package main

import (
	"context"

	"google.golang.org/api/sheets/v4"
)

func main() {
	ctx := context.Background()
	sheetsService, err := sheets.NewService(ctx)
}
