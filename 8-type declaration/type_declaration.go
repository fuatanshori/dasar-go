package main

import "fmt"

func main() {
	// type declaration is alias 
	// ex : type umur as int

	type number int8

	var umur number = 13
	fmt.Println(umur)

	var umur2 int16 = 22
	var umur3 number = number(umur2)
	fmt.Println(umur3) 
}