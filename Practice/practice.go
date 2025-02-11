package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println("Error Occured", err)
		if os.IsNotExist(err) {
			fmt.Println("The File does not exist")
		}
		return
	}
	fmt.Println("File opened succesfully")

	defer file.Close()

	Data, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("error occured while reading the file", err)
		return
	}
	fmt.Println("--------------------------")
	fmt.Println("FILE CONTENT:=")
	fmt.Println(string(Data))
}
