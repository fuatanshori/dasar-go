package main

import "fmt"

func main() {
	nilai := 50
	switch nilai{
	case 80:
		fmt.Println("nila 80")
	case 50:
		fmt.Println("nilai 50")
	default:
		fmt.Println("bukan 80 dan 50")
	}
}