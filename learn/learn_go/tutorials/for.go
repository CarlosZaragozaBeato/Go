package main 

import "fmt"


func main(){
	
	i := 1 

	for i <= 3 {
		fmt.Println(i)
		i += 1 
	}

	for j:= 7; j <=9; j++ {
		fmt.Println("LOGO: %v", j)
	}

	for {
		fmt.Println("LOOP")
		break
	}

	for n := 0; n<=5; n++ {
		if n % 2 ==0 {
			continue 
		}
		fmt.Println(n)
	}
}