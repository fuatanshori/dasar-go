package main

import "fmt"

func main() {
	var array1 [2]string
	array1[0]="Fuat"
	array1[1]="Anshori"
	fmt.Println(array1)

	var array2[2] int
	array2[0] = 20041066
	array2[1]= 20
	fmt.Println(array2)

	var array4 = [2]int{
		10,20,
	}
	fmt.Println(array4)

	array5 := [4]int{
		20041066,10,14,11,
	}
	fmt.Println(array5)
	fmt.Println(len(array5))
	fmt.Println(array5[0])
	array5[0]=21
	fmt.Println(array5[0])

	// [...] can be use if declared directly 
	array6:=[...]int{
		10,21,11,12,19,
	}
	fmt.Println(array6)
	fmt.Println(len(array6))
}