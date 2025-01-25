// Using of Variables at package and block level
package main

import "fmt"

var Name string

const Pi float64 = 3.14

func main() {
	var Block string = "Block is at func level or called block level"

	var Name = "this Name is at package level and its value can be changed"
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("Block=%s\nName:%s\nPi is a constant and the value is %.2f\n", Block, Name, Pi)
	fmt.Println("------------------------------------------------------------------")
}
