package sheets

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
