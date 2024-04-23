package sheets

import "time"

type Items []struct {
	ID          string    `json:"id"`
	CreatedTime time.Time `json:"created_time"`
	CreatedBy   struct {
		ID string `json:"id"`
	} `json:"created_by"`
	LastEditedTime time.Time `json:"last_edited_time"`
	LastEditedBy   struct {
		ID string `json:"id"`
	} `json:"last_edited_by"`
	Parent struct {
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
			URL        string    `json:"url"`
			ExpiryTime time.Time `json:"expiry_time"`
		} `json:"file"`
	} `json:"cover,omitempty"`
	Properties struct {
		BM struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"BM"`
		Campanha struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			Relation []struct {
				ID string `json:"id"`
			} `json:"relation"`
		} `json:"Campanha"`
		Card struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"Card"`
		CardCriativo struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"Card Criativo"`
		Cliente struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"Cliente"`
		CriadoEm struct {
			ID          string    `json:"id"`
			Type        string    `json:"type"`
			CreatedTime time.Time `json:"created_time"`
		} `json:"Criado em"`
		DataDePostagem struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Date struct {
				Start string `json:"start"`
			} `json:"date"`
		} `json:"Data de Postagem"`
		Demanda struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Demanda"`
		Departamento struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Departamento"`
		FabricaO struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"Fabricação"`
		FluxoCinema struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Status struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"status"`
		} `json:"Fluxo - Cinema"`
		FluxoDesigner struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Status struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"status"`
		} `json:"Fluxo - Designer"`
		Post struct {
			ID    string `json:"id"`
			Type  string `json:"type"`
			Title []struct {
				Type        string `json:"type"`
				Annotations struct {
					Bold  bool   `json:"bold"`
					Color string `json:"color"`
				} `json:"annotations"`
				PlainText string `json:"plain_text"`
				Text      struct {
					Content string `json:"content"`
				} `json:"text"`
			} `json:"title"`
		} `json:"Post"`
		ResponsVel struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Responsável"`
		StatusDoCard struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Status struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"status"`
		} `json:"Status do Card"`
		TipoDoEntregVel struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Tipo do Entregável"`
		LtimaEdiO struct {
			ID             string    `json:"id"`
			Type           string    `json:"type"`
			LastEditedTime time.Time `json:"last_edited_time"`
		} `json:"Última edição"`
	} `json:"properties"`
}
