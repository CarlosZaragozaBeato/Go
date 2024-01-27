package main 

import "fmt"

const englishHelloPrefix = "Hello, "


func Hello(value string) string {
	return englishHelloPrefix + value
}



func main() { 
	fmt.Println("Hello, World")
	fmt.Println(Hello("GACHA"))

}