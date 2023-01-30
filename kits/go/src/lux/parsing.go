package lux

import (
	"encoding/json"
	"os"
)

type Info struct {
	EnvCfg EnvConfig `json:"env_cfg"`
}

type Input struct {
	Step                 int         `json:"step"`
	Player               string      `json:"player"`
	RemainingOverageTime int         `json:"remainingOverageTime"`
	Info                 Info        `json:"info"`
	Obs                  Observation `json:"obs"`
}

var decoder = json.NewDecoder(os.Stdin)

func ParseGameState() (*Input, error) {
	var input Input
	err := decoder.Decode(&input)
	if err != nil {
		Log("could not decode json input: %s", err.Error())
		return nil, err
	}
	return &input, nil
}
