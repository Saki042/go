package main

import "fmt"

// This is comment

func main() {

	// Declaring the Variables
	name := "Sakii"
	fmt.Println(name)
	// Declaring the String varible with proper syntax
	var greeting string = "Hi from Go"
	fmt.Println(greeting)

	// Printing Vraibles
	fmt.Print("Hello fmt pkg")
	var city string = "Karnataka"
	fmt.Println(city)
	var nameuser string = "Sakii"
	var user string = "LearningwithSakii"
	fmt.Println("Welcome to ", user, ", ", nameuser)

	// Newline character \n
	var firstname string = "Sakii"
	var lastname string = "Mishra"

	fmt.Print(firstname, "\n")
	fmt.Print(lastname)

	//Println
	fmt.Println(firstname)
	fmt.Println(lastname)

	// Printf - format specifier
	// %v - format the value in default format
	fmt.Println("%v - format the value in default format")

	var formatdefault string = "Day1withGo"
	fmt.Printf("Nice to see you here, at %v", formatdefault)
	fmt.Print("\n")
	fmt.Println("%d - format decimal integers.")
	// %d - format decimal integers.

	var grades int = 42
	fmt.Printf("Marks: %d", grades)
	fmt.Print("\n")

	// Combining all
	var example string = "Joe"
	var i int = 32
	fmt.Printf("Hey %v! You have scored %d/100 in xyz", example, i)
}
