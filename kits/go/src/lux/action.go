package lux

import "encoding/json"

type UnitActionType int

const (
	UnitActionMove         UnitActionType = 0
	UnitActionTransfer     UnitActionType = 1
	UnitActionPickup       UnitActionType = 2
	UnitActionDig          UnitActionType = 3
	UnitActionSelfDestruct UnitActionType = 4
	UnitActionRecharge     UnitActionType = 5
)

type UnitAction struct {
	Type      UnitActionType
	Direction Direction
	Distance  int
	Resource  Resource
	Amount    int
	Repeat    int
	N         int
}

func (a *UnitAction) UnmarshalJSON(data []byte) error {
	var v [6]int
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	a.Type = UnitActionType(v[0])
	a.Direction = Direction(v[1])
	if a.Type == UnitActionTransfer || a.Type == UnitActionPickup {
		a.Resource = Resource(v[2])
	}
	a.Distance = v[2]
	a.Amount = v[3]
	a.Repeat = v[4]
	a.N = v[5]
	return nil
}

func (a UnitAction) MarshalJSON() ([]byte, error) {
	var v [6]int
	v[0] = int(a.Type)
	v[1] = int(a.Direction)
	v[2] = a.Distance
	if a.Type == UnitActionTransfer || a.Type == UnitActionPickup {
		v[2] = int(a.Resource)
	}
	v[3] = a.Amount
	v[4] = a.Repeat
	v[5] = a.N
	return json.Marshal(&v)
}

type FactoryActionType int

const (
	FactoryActionBuildLight FactoryActionType = 0
	FactoryActionBuildHeavy FactoryActionType = 1
	FactoryActionWater      FactoryActionType = 2
)

type FactoryAction struct {
	Type FactoryActionType
}

func (a *FactoryAction) UnmarshalJSON(data []byte) error {
	var v int
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	a.Type = FactoryActionType(v)
	return nil
}

func (a FactoryAction) MarshalJSON() ([]byte, error) {
	var v int
	v = int(a.Type)
	return json.Marshal(&v)
}
