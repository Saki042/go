package main

import "fmt"

// This is comment

func main() {

	fmt.Println("Hello World!!")
	// Declaring the Variables
	name := "Sakshi"
	fmt.Println(name)
	// Declaring the String varible with proper syntax
	var greeting string = "Hi from Go"
	fmt.Println(greeting)

	// Printing Vraibles
	fmt.Print("Hello fmt pkg")
	var city string = "U.P"
	fmt.Println(city)
	var nameuser string = "Sakshi"
	var user string = "kodekloud"
	fmt.Println("Welcome to ", user, ", ", nameuser)

	// Newline character \n
	var firstname string = "Sakshi"
	var lastname string = "Mishra"

	fmt.Print(firstname, "\n")
	fmt.Print(lastname)

	//Println
	fmt.Println(firstname)
	fmt.Println(lastname)

	// Printf - format specifier
	// %v - format the value in default format

	var formatdefault string = "Day1withGo"
	fmt.Printf("Nice to see you here, at %v", formatdefault)

}
