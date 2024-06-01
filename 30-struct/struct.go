package main

import "fmt"

type Number int
type Mahasiswa struct{
	Nama, Jurusan string
	Nim, Umur Number
}


func main() {
	fuat := Mahasiswa{Nama: "Fuat Anshori", 
	Jurusan: "Teknik Informatika", Nim: 20041066,
	Umur: 24,
	}
	fmt.Println(fuat)
}

