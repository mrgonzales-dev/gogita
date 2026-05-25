package main

import (
	"fmt"
)

func checkage(age int) string {
	if age > 18 {
		return "You are an adult"
	} else {
		return "You are a child"
	}
}

// var variablename type = value
func main() {
	f_name := "Kenneth"
	s_name := "Gonzales"
	var age = 22

	fmt.Println("Hello, I am " + f_name + " " + s_name)
	fmt.Println("I am " + fmt.Sprint(age) + " years old")
	fmt.Println(checkage(age))
}
