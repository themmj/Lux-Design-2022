package lux

type Observation struct {
	Board        Board                         `json:"board"`
	Units        map[string]map[string]Unit    `json:"units"`
	Teams        map[string]Team               `json:"teams"`
	Factories    map[string]map[string]Factory `json:"factories"`
	RealEnvSteps int                           `json:"real_env_steps"`
}
