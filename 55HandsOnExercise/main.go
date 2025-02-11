package main

import "fmt"

func main() {
	type Engine struct {
		Electric bool
	}

	type Vehicle struct {
		Engine Engine
		Model  string
		Doors  int
		Colour string
	}

	V1 := Vehicle{
		Engine: Engine{Electric: true},
		Model:  "2019",
		Doors:  5,
		Colour: "Black",
	}
	V2 := Vehicle{
		Engine: Engine{Electric: false},
		Model:  "2024",
		Doors:  5,
		Colour: "White",
	}
	AllVeh := []Vehicle{V1, V2}

	fmt.Println("Allveh.", AllVeh)
	fmt.Println("V1", V1)
	fmt.Println("V2", V2)

	fmt.Println(V1.Engine.Electric, V1.Model, V1.Colour)
	fmt.Println(V2.Engine.Electric, V2.Model, V2.Colour)

}
