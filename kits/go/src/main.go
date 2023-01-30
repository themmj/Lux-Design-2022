package main

import (
	"fmt"
	"go-agent/lux"
)

func main() {
	for {
		input, err := lux.ParseGameState()
		lux.DumpJsonToFile("input.json", &input)
		if err != nil {
			break
		}
		fmt.Println("{}")
	}
}
