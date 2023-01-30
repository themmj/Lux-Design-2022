package lux

import "encoding/json"

type Direction int

const (
	DirCenter Direction = 0
	DirUp     Direction = 1
	DirRight  Direction = 2
	DirDown   Direction = 3
	DirLeft   Direction = 4
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
