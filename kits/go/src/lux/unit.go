package lux

type Unit struct {
	TeamId      int          `json:"team_id"`
	UnitId      string       `json:"unit_id"`
	Power       int          `json:"power"`
	UnitType    string       `json:"unit_type"`
	Pos         Position     `json:"pos"`
	Cargo       Cargo        `json:"cargo"`
	ActionQueue []UnitAction `json:"action_queue"`
}
