package main

import "fmt"

// 3
type Blacklist func(string)bool

// 2
func registerUser(name string, blaklist Blacklist){
	// 5.1
	if blaklist(name){
		fmt.Println("You are blocked",name)
	// 5.2
	}else{
		fmt.Println("Welcome",name)
	}
}

func main() {
	// 1
	registerUser("Anjing",
	// 4
	func(name string) bool {
		return name == "Anjing"
	})
}