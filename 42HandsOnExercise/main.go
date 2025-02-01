package main

import (
	"fmt"

)

func main() {
	// 	// #109 Hands on Exercise =42
	// 	A := [5]int{5, 4, 6, 5, 6}
	// 	for V := range A {
	// 		fmt.Printf("The Type is %T\n\t and the value is %v\n", V, V)
	// 	}
	// }

	// // #110 Hands on Exercise =43
	// B:=make([]int,0,11)
	// B=append(B,40,41,42,43,44,45,46,47,48,49)

	// for _,V:=range B{
	// 	fmt.Printf("Value=%v\t And the Type=%T\n",V,V)
	// }

	// //111 Hands On Exercise =44

	// B:=[]int{42,43,44,45,46,47,48,49,50,51}
	// fmt.Println(B[:5])
	// fmt.Println(B[5:])
	// fmt.Println(B[2:7])
	// fmt.Println(B[1:6])

	// //113 Hands On Exercise 46

	// X := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// fmt.Println(X)
	// // output 44,45,46,47 DONT WANT TO PRINT THESE

	// X = append(X[:2], X[6:]...)
	// fmt.Println(X)

	//114 Hands ON Exercise 47

	// S := make([]string, 0, 50)
	// S = append(S, "Andhra Pradesh", "Arunachal Pradesh", "Assam", "Bihar", "Chhattisgarh", "Goa", "Gujarat", "Haryana", "Himachal Pradesh", "Jharkhand", "Karnataka", "Kerala", "Madhya Pradesh", "Maharashtra", "Manipur", "Meghalaya", "Mizoram", "Nagaland", "Odisha", "Punjab", "Rajasthan", "Sikkim", "Tamil Nadu", "Telangana", "Tripura", "Uttar Pradesh", "Uttarakhand", "West Bengal")

	// fmt.Println("Lenth", len(S))
	// fmt.Println("Capacity", cap(S))
	// for A := 0; A < len(S); A++ {
	// 	fmt.Println(A+1, ".", S[A])
	// }






	// #115 Hands ON Exercise 
	X:=[]string{"Aman","Raju"}

	Y:=[]string{"Rama","Bhushan"}

	Rah:=[][]string{X,Y}
	fmt.Println(Rah)

	for _,Value:=range Rah{
		fmt.Println("-----------------------")
		for Index,Data:=range Value{
			fmt.Printf("Index=%d  Data=%s\n",Index,Data)
		}
	}
}
