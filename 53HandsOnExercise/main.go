package main

import "fmt"

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
	fmt.Println(Person1)
	fmt.Println(Person2)
	Data:=[]Person{Person1,Person2}
	fmt.Println(Data)
	fmt.Println("------------------------")
	for _,Values:= range Data{
	fmt.Printf("Type =%T\t Value:%v\t  Total conclusion:%#v\n",Values,Values,Values)
	}
}