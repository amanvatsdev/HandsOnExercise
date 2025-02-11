package main

import "fmt"

func main() {

	//121,122,123 --49,50,51 Hands on exercise

	Info := map[string][]string{
		"Kumar_Ravi":  {"Shooting", "Singing"},
		"Verma_Raj":   {"Swimming", "Dancing"},
		"Gupta_Pawan": {"Playing Cricket", "Mimicry"},
	}

	// for Index, Value := range Info {
	// 	fmt.Printf("%s.\n", Index)
	// 	for I, V := range Value {
	// 		fmt.Printf("\t%d.%s\n", I+1, V)
	// 	}
	// }
	//Adding data of Sunny Sharma in Data
	Info["Sharma_Sunny"] = []string{"Literature", "Singing", "Swimming"}
	// fmt.Println("-----------------")

	// for Index, Value := range Info {
	// 	fmt.Printf("%s.\n", Index)
	// 	for I, V := range Value {
	// 		fmt.Printf("\t%d.%s\n", I+1, V)
	// 	}
	// }

	//Deleting the name of Ravi Kumar

	// Info["Sharma_Sunny"] = append(Info["Sharma_Sunny"][1:])
	// Info["Sharma_Sunny"]=Info["Sharma_Sunny"][2:]

	Info["Sharma_Sunny"] = append(Info["Sharma_Sunny"][:1], Info["Sharma_Sunny"][2:]...)

	fmt.Println("================================")

	for Index, Value := range Info {
		fmt.Printf("%s\n", Index)
		for I, V := range Value {
			fmt.Printf("\t%d.%s\n", I+1, V)
		}
	}

}
