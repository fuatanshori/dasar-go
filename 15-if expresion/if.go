package main
  
import "fmt"

func main() {
	
	nilai := 20
	kkm :=70

	if nilai<=kkm {
		fmt.Println("tidak lulus")
	}else if nilai >=kkm{
		fmt.Println("lulus")
	}else{
		fmt.Println("tidak diketahui")
	}

}