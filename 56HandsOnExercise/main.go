package main

import "fmt"

func main() {
	P1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Aman",
		friends: map[string]int{
			"Ronny": 26,
			"Ravi":  25,
			"Pawan": 22,
		},
		favDrinks: []string{"Coka Cola", "Thumsup", "Sprite"},
	}

	fmt.Println(P1)
}