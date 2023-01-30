package lux

import "encoding/json"

type Direction int

const (
	Center Direction = 0
	Up     Direction = 1
	Right  Direction = 2
	Down   Direction = 3
	Left   Direction = 4
)

type Position struct {
	X int
	Y int
}

func (p *Position) UnmarshalJSON(data []byte) error {
	var v [2]int
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	p.X = v[0]
	p.Y = v[1]
	return nil
}

func (p Position) MarshalJSON() ([]byte, error) {
	var v [2]int
	v[0] = p.X
	v[1] = p.Y
	return json.Marshal(&v)
}
