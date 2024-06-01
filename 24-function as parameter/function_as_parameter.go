package main

import "fmt"

// 2
func getName(name string,filter func(name string) string){
	// 3
	filteredName := filter(name)
	// 6
	fmt.Println(filteredName)
}

// 4
func badWordName(name string)string{	
	// 5.1
	if name == "Anjing"{
		return "****"
	// 5.2
	}else{
		return name
	}
}

func main() {
	// 1
	getName("Fuat Anshori",badWordName) 
}