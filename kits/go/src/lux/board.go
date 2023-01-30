package lux

type Board struct {
	Ice              [][]int  `json:"ice"`
	Lichen           [][]int  `json:"ice"`
	LichenStrains    [][]int  `json:"ice"`
	Ore              [][]int  `json:"ice"`
	Rubble           [][]int  `json:"ice"`
	ValidSpawnsMask  [][]bool `json:"ice"`
	FactoryOccupancy [][]int  `json:"ice"`
	FactoriesPerTeam int      `json:"factories_per_team"`
}
