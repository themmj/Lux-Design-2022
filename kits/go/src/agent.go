package main

import "go-agent/lux"

type Agent struct {
	Step                 int
	Player               string
	RemainingOverageTime int
	Obs                  *lux.Observation
	Cfg                  *lux.EnvConfig
}

func (a *Agent) Setup() interface{} {
	return lux.NewBidAction("AlphaStrike", 100)
}

func (a *Agent) Act() interface{} {
	return lux.NewBidAction("AlphaStrike", 100)
}
