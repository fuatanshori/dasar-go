package main

import "fmt"

func sumAll(numbers ...int)(total int){
	total = 0
	for i := 0; i < len(numbers); i++ {
		total+=numbers[i]
	}
	return
}

func main() {
	fmt.Println(sumAll(1,2,3,46,7,8,9,10,11,12,13))
	num := []int{11,12,13,14}
	fmt.Println(sumAll(num...))

}
