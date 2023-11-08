package main

import "fmt"

func main() {
	//second_function()
	third_function()

}

// func second_function() {
// 	fmt.Println("Second Function call")
// }

func third_function() {

	// strings
	var name_one string = "carlos"
	var name_two = "fenomeno"
	var name_three string

	fmt.Println(name_one, name_two, name_three)

	name_two = "Gaucho"

	fmt.Println(name_one, name_two, name_three)

	name_four := "EL FENOMENO || GAUCHO"
	fmt.Println(name_four)

	// ints
	var age_one int = 18
	var age_two = 22
	age_three := 44

	fmt.Println(age_one, age_two, age_three)

	// int bytes
	var one_bytes int8 = 8
	var two_bytes int16 = 44
	var three_bytes int32 = 4751
	var fourth_bytes uint64 = 4751

	fmt.Println(one_bytes, two_bytes, three_bytes, fourth_bytes)
}
