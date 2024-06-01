package main

import "fmt"

func main() {
	arr1 := [...]string{"Fuat","Anshori","John","Doe"}
	slicearr1 := arr1[1:3]
	slicearr2 := arr1[1:]
	var slicearr3 []string= arr1[:3]
	slicearr4 := arr1[:]
	fmt.Println(slicearr1)
	fmt.Println(slicearr2)
	fmt.Println(slicearr3)
	fmt.Println(slicearr4)
	
	slicearr1[0]="Alicia"
	fmt.Println(arr1)

	days := [...]string{"Senin","selasa","rabu"}
	fromSlice := days[:]
	toSlice := make([]string,len(fromSlice),cap(fromSlice))
	copy(toSlice,fromSlice)
	fmt.Println(toSlice) 
}