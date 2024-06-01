package main

import "fmt"

func getGoodbye(name string) string{
	return "hello "+name
}

func main() {
	getgoodbye := getGoodbye
	fmt.Println(getgoodbye("Fuat anshori"))

}