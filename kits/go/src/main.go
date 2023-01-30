package main

import (
	"encoding/json"
	"fmt"
	"go-agent/lux"
)

func main() {
	var agent = new(Agent)
	for {
		input, err := lux.ParseGameState()
		if err != nil {
			lux.Log(err.Error())
			break
		}
		lux.DumpJsonToFile("input.json", &input)

		agent.Step = input.Step
		agent.Player = input.Player
		agent.RemainingOverageTime = input.RemainingOverageTime
		agent.Obs = &input.Obs
		agent.Cfg = &input.Info.EnvCfg

		var res interface{}
		if agent.Step < 0 {
			res = agent.Setup()
		} else {
			agent.Step = agent.Obs.RealEnvSteps
			res = agent.Act()
		}

		b, err := json.Marshal(res)
		if err != nil {
			lux.Log(err.Error())
			break
		}
		fmt.Println(string(b))
	}
}
