// Hands on Exercise creating build by 1.go build main.go 2.go build . 3.go build ./...
// all these three work in same work
// 1.(go build main.go)  that will make the file main.go simply executable just using --./main.exe
// 2.go build .) that will make the file like <folder Name>.exe and can execute by ./<folder Name>.exe
// 3.go build ./... that will make all files executable and executed by ./<folder Name>.exe
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
