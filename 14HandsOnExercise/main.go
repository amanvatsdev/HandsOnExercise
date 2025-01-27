package main

import (
	"fmt"

	"github.com/amanvatsdev/Learning/23012025/dependecy/Sound"
)

func Wolf() {
	fmt.Println("Wooh !!")
}
func AdultWolf() {
	fmt.Println("Wohh! Wohh! Wohh!!")
}
func main() {
	Wolf()
	AdultWolf()
	Sound.Cat()
}
