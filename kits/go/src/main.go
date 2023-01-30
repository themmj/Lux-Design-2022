package main

import (
	"encoding/json"
	"go-agent/lux"
	"os"
)

type Info struct {
	EnvCfg lux.EnvConfig `json:"env_cfg"`
}

type Input struct {
	Step                 int    `json:"step"`
	Player               string `json:"player"`
	RemainingOverageTime int    `json:"remainingOverageTime"`
	Info                 Info   `json:"info"`
}

func main() {
	var decoder = json.NewDecoder(os.Stdin)
	var input Input
	err := decoder.Decode(&input)
	if err != nil {
		lux.Log("could not decode json input")
	}
	lux.DumpJsonToFile("input.json", &input)
}
