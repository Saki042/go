package main

import "fmt"

// Global variable
var nameuser string = "Lisaa"

func main() {

	/*
		BLOCK
		Inner block - can access var declaraed within outer blocks.
		Outer block - can't access var declared within inner blocks.
	*/
	city := "London"
	{
		country := "UK"
		fmt.Println(country)
		fmt.Println(city)

	}

	fmt.Println(city)

	/*
		Local Variables = declared inside function or a block
		Global variables = Declared outside of a function or a block
	*/
	// local variable

	var nameex string = "IamNameVar"
	fmt.Println(nameex)

	// Accessing the Global variable
	fmt.Println(nameuser)

	/*	 Zero Values
	bool - false
	int - 0
	float64 - 0.0
	string - ""
	*/

	var o int
	fmt.Printf("%d", o)
	fmt.Print("\n")
	var f1 float64
	fmt.Printf("%.2f", f1)

	// USER INPUT
	fmt.Scanf("%<format specifier> (s)", object_arguments)

	//var userip string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &userip)
	fmt.Println("Hey,", userip)

	// Multiple Input
	var is_muggle bool

	fmt.Print("Enter your name and are you a muggle: ")
	fmt.Scanf("%s %t", &userip, &is_muggle)
	fmt.Println(userip, is_muggle)

	/* Scanf return values
	count - the no. of arguments that the fucntion writes to
	err - any error thrown during the execution of the fucntion
	*/
	var a string
	var b int

	fmt.Println("Enter a string & Integrer value: ")

	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count: ", count)
	fmt.Println("error: ", err)
	fmt.Println("a: ", a, "%t", "b: ", b)

}
