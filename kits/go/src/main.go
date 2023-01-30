package main

import (
	"encoding/json"
	"os"
    "go-agent/lux"
)

type Info struct {
    Step int64 `json:"step"`;
    Player string `json:"player"`;
    RemainingOverageTime int64 `json:"remainingOverageTime"`;
}

func main() {
    var decoder = json.NewDecoder(os.Stdin)
    var info Info
    err := decoder.Decode(&info)
    if err != nil {
        lux.Log("could not decode json input")
    }
    lux.DumpJsonToFile("input.json", &info);
    lux.Log("read %d, %s, %d", info.Step, info.Player, info.RemainingOverageTime);
}

