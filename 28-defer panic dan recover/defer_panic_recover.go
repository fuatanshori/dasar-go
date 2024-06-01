package main

import "fmt"

func logging(){
	fmt.Println("Selesai Memanggil Fungsi")
	r:=recover()
	fmt.Println("pesan",r)
}

func runApp(error bool){
	defer logging()
	fmt.Println("Run Application")
	if error{
		panic("ups Error")
	}
}

func main() {
	runApp(true)
}