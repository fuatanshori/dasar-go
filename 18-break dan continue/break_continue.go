package main

import "fmt"

func main() {
	names := []string{"Fuat","Anshori","Alicia","Alexandra","John","Doe"}

	for _, name := range names {
		if name == "Anshori"{
			continue
		}else if name == "Alicia"{
			break
		}
		fmt.Println(name)
	}

	count :=10
	for i := 0; i < count; i++ {
		if i % 2 == 0{
			continue
		}
		fmt.Println(i)
	}
}