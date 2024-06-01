package main

import "fmt"

func hello(message string)(response string){
	response = "hello "+ message
	return
}

func apiData(id string)(response  map[string]string, response2 string){
	data :=map[string]string{
		"id" : "1","nama":"Alicia","asal":"Jakarta","status":"200",
	}
	if data["id"]==id {
		response = data
	}else{
		response = map[string]string{"status":"400"}
	}
	return response,"None"
}



func main() {
	fmt.Println(hello("Fuat Anshori"))
	data,_ := apiData("1")
	fmt.Println(data)
}