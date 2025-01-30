package main

import (
	"fmt"
	"math/rand"
)

func main() {
	X := rand.Intn(250)
	fmt.Printf("Name of Var =%d\t Value of Var =%d\n", X, X)
	if X<100{
		fmt.Println("The No is between 0 to 100")
	}else if  X <200 && X >=100{
		fmt.Println("The Number is between 100 to 200")
	}else {
		fmt.Println("The Number is between 200 to 250")
	}

	Y :=rand.Intn(3)
	fmt.Println(Y)
}
