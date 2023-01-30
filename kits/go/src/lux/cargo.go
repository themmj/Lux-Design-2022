package lux

type Resource int

const (
	ResourceIce   Resource = 0
	ResourceOre   Resource = 1
	ResourceWater Resource = 2
	ResourceMetal Resource = 3
	ResourcePower Resource = 4
)

type Cargo struct {
	Ice   int `json:"ice"`
	Ore   int `json:"ore"`
	Water int `json:"water"`
	Metal int `json:"metal"`
}
