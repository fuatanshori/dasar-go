package main

import "fmt"

func main() {
	var name string
	var nim int32

	name = "Fuat"
	fmt.Println(name)
	name ="Fuat Anshori satu"
	fmt.Println(name)

	nim = 20041066
	fmt.Println(nim)

	var name2 string = "Fuat Anshori Dua"
	fmt.Println(name2)
	
	name3 := "Fuat Anshori Tiga"
	fmt.Println(name3) 

	var (
		name4 string= "Fuat empat"  
		name5 string= "Anshori lima"  
	)
	fmt.Println(name4,name5)

}