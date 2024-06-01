package main

import "fmt"

type HashName interface{
	getName() string
}

type Animal struct{
	Name string
}

func (a Animal) getName()string{
	return a.Name
}

func main() {
	burung := Animal{"Beo"}
	fmt.Println(burung.getName())
}