package lux

type Team struct {
	TeamId           int    `json:"team_id"`
	Faction          string `json:"faction"`
	Water            int    `json:"water"`
	Metal            int    `json:"metal"`
	FactoriesToPlace int    `json:"factories_to_place"`
	FactoryStrains   []int  `json:"factory_strains"`
	PlaceFirst       bool   `json:"place_first"`
}
