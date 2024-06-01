package main

import (
	"fmt"
)

func main() {
	dataMahasiswa := map[string]string{
		"nama":"Fuat Anshori",
		"nim":"20041066",
		"asal":"Palembang",
	}
	delete(dataMahasiswa,"nim")
	fmt.Println(dataMahasiswa["nama"])
	fmt.Println(dataMahasiswa)

	newMap := make(map[string]string)
	newMap["message"] = "success"
	newMap["code"]="201"
	fmt.Println(newMap)

}