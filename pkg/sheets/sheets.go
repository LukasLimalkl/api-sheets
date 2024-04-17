package sheets

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func ConnectSheets() {
	ctx := context.Background()
	sheetsService, err := sheets.NewService(ctx, option.WithAPIKey("client"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sheetsService)
}
