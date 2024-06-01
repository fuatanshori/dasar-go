package main

import "fmt"

func main() {
	// count := 1
	// for count < 10 {
	// 	fmt.Println("count ===> ", count)
	// 	count++
	// }

	// for i :=0 ;i<10;i++{
	// 	fmt.Println(i)
	// }

	names:=[]string{
		"Fuat","Anshori","Asal","Banjarmasin",
	}
	for i:= 0;i<len(names);i++{
		fmt.Println(names[i])
	}
	for index, name := range names {
		fmt.Println("index : ",index,",name : ",name)
	}
	for _, name := range names {
		fmt.Println("name : ",name)
	}
}
