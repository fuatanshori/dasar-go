package main

import "fmt"

func main() {
	a:=10
	b:=11
	c:=a+b
	fmt.Println(c)
	
	d:=14
	e:=10
	f:=d-e
	fmt.Println(f)

	g:=10
	g+=1
	fmt.Println(g)
	g+=10
	fmt.Println(g)
	g++
	fmt.Println(g)
}