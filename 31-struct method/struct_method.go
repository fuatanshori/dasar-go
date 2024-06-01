package main

import "fmt"

type Number int
type Mahasiswa struct{
	Name, Jurusan string
	Nim, Umur Number
}

func (m Mahasiswa) getName(){
	fmt.Println(m.Name)
}

func main() {
	fuat := Mahasiswa{Name: "Fuat Anshori", 
	Jurusan: "Teknik Informatika", Nim: 20041066,
	Umur: 24,
	}
	fuat.getName()
}

