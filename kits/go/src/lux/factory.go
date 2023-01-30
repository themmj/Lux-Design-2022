package lux

type Factory struct {
	Pos         Position        `json:"pos"`
	Power       int             `json:"power"`
	Cargo       Cargo           `json:"cargo"`
	UnitId      string          `json:"unit_id"`
	StrainId    int             `json:"strain_id"`
	TeamId      int             `json:"team_id"`
	ActionQueue []FactoryAction `json:"action_queue"`
}
