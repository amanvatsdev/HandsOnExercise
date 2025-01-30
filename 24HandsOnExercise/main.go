package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for A := 0; A < 101; A++ {
	// 	fmt.Println(A)
	// }
	//#25
	// for Y := 1; Y < 43; Y++ {
	// 	X := rand.Intn(5)
	// 	switch {
	// 	case X == 0:
	// 		fmt.Printf("%d.The Value of X %d \n", Y, X)
	// 	case X == 1:
	// 		fmt.Printf("%d.The Value of X %d \n", Y, X)
	// 	case X == 2:
	// 		fmt.Printf("%d.The Value of X %d \n", Y, X)
	// 	case X == 3:
	// 		fmt.Printf("%d.The Value of X %d \n", Y, X)
	// 	case X == 4:
	// 		fmt.Printf("%d.The Value of X %d \n", Y, X)
	// 	case X == 5:
	// 		fmt.Printf("%d.The Value of X %d \n", Y, X)
	// 	default:
	// 		fmt.Printf("%d.The Number is  %dm\no,Yre than 5")
	// 	}
	// }

	//#26
	// A:=15
	// for A>=0{
	// 	fmt.Printf("Value of A:%d\n",A)
	// 	A--
	// }
	// #27:
	// 	A:=0
	// 	for {
	// 		fmt.Printf("The value of A:=%d\n",A)
	// 		if A>=22{
	// 			break
	// 		}
	// 		A++
	// 	}

	// #28

	// for X := 0; X < 50; X++ {
	// 	if X%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(X)

	// }

	// #29
	// for X := 1; X < 6; X++ {
	// 	fmt.Printf("This is the main loop that will reapeat 5 times and the Sno. is%d\t\n", X)
	// 	for Y := 1; Y < 5; Y++ {
	// 		fmt.Printf("Nested loop Sno: %d\t\n", Y)
	// 	}
	// }
	// #30
	// Xi:=[]int{42,43,44,45,46,47,48}
	// 	for In,Val:=range Xi{
	// 		fmt.Printf("%d.The value is %d\n",In+1,Val)
	// 	}

	//  #31
	// Data := map[string]int{
	// 	"Aman":   22,
	// 	"Vikash": 32,
	// 	"Sinu":   25,
	// }
	// for In, Val := range Data {
	// 	fmt.Printf("Name is :%s    Age is :%d\n", In, Val)
	// }

	// #32
	// Data := map[string]int{
	// 	"Aman":   22,
	// 	"Vikash": 32,
	// 	"Sinu":   25,
	// }
	// for In, Val := range Data {
	// 	fmt.Printf("Name is :%s    Age is :%d\n", In, Val)
	// }

	// DataNew := Data["Aman"]

	// if Value, ok := Data["Aman"]; ok {
	// 	fmt.Println("----The age of Aman is==  ", Value)
	// }
	// fmt.Println("The Age Of Aman is ", DataNew)
	// // fmt.Println("The Age Of Aman is ", DataNew)
	// DataNew2 := Data["Q"]

	// fmt.Println("The Age Of Aman is ", DataNew2)

	// if Value, ok := Data["Q"]; !ok {
	// 	fmt.Println("Thier is no data of Q and the value of Q is ", Value)
	// }

	// 33
	Total := 0
	for Y := 0; Y < 101; Y++ {
		if X := rand.Intn(5); X == 3 {
			fmt.Printf("Value of X is  %d\t anf the total times is %d\n", X, Total)
			Total++
		}
	}

	// //#34
	// fmt.Println("--true || false--", (true || false))
	// fmt.Println("--true && false--", (true && false))
	// fmt.Println("--true || true--", (true || true))
	// fmt.Println("--true && true--", (true && true))
	// fmt.Println(!true)

}
