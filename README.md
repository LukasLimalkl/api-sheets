# Notion to Google Sheets Integration 
## Description

This project provides an automated solution for exporting data from a Notion database to Google Sheets, ideal for teams and individuals who need to synchronize information between these two popular platforms for enhanced data visualization, manipulation, and analysis. The integration is carried out via a Go application that utilizes the Notion and Google Sheets APIs.

## Features

- Data extraction from Notion databases using the Notion API.
- Automatic data uploading to Google Sheets.
- Flexible configuration for varying database schemas.
- Optional scheduling for automatic updates.

## Prerequisites

Before you begin, ensure you have:
- Go 1.15 or higher.
- Access to Notion with an API integration set up.
- Access to Google Sheets with API credentials configured.

## Setup

### Environment Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/lukaslimalkl/api-sheets
   cd api-sheets

2. Download the dependencies:
   ```bash
   go mod tidy

### Setting Up Notion Credentials

1. Obtain the integration token from Notion by following the instructions here.
2. Create a .env file in the root of the project and add your token:
   ```bash
    API_NOTION="YOU_NOTION_API"
    DB_NOTION="YOU_NOTION_DB"
    GOOGLE_CLIENT_ID="YOU_SHEETS_ID"
    GOOGLE_CLIENT_SECRET="YOU_SHEET_SECRET"

### Usage

1. To run the application and sync the data, use the following command:
    ```bash
     go run cmd/main.go

### Contributing

  Contributions are welcome! If you have suggestions for improving this integration, feel free to fork the repository and submit a pull request, or open an issue with the tags "enhancement" or "bug".




