package main

import "fmt"

func factorial(n int)int{
	if n == 1{
		return n
	}else{
		return n *factorial(n-1)
	}
}

func main() {
	fmt.Println(factorial(10))
}