package main

import (
	"fmt"
)

func main() {
	type Person struct {
		FirstName         string
		LastName          string
		FavouriteIceCream []string
	}

	Person1 := Person{
		FirstName:         "Aman",
		LastName:          "Vats",
		FavouriteIceCream: []string{"vanilla", "Choclate"},
	}
	Person2 := Person{
		FirstName:         "Ravi",
		LastName:          "Parashar",
		FavouriteIceCream: []string{"Butterscotch", "Cornato"},
	}

	//start from here upper one copied from 53 hands on exercise
	v := map[string]Person{
		Person1.LastName: Person1,
		Person2.LastName: Person2,
	}
	fmt.Println(v)
	for I, Value := range v {
		fmt.Printf("Index :%s  Value:%v\n", I, Value)
		for SNo,Data:= range Value.FavouriteIceCream{
			fmt.Printf("Sno:%d   Data:%s\n",SNo,Data)
		}
	}
}
