package main

//import "fmt"
import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Hello from Variables section!!!")
	/* Checking the Type of variable
	%T - format specifier
	reflect.TypeOf - function from the reflect pkg.
	*/
	var grades int = 56
	fmt.Printf("Variable grades = %v is of type %T \n", grades, grades)

	/* Data Type conversion - Type Casting

	 */

	var h int = 90
	var fl float64 = float64(h)
	fmt.Printf("%.2f \n", fl)
	var j int = int(fl)
	fmt.Printf("%v \n", j)

	/* string conversion package - strconv
	1. Itoa()
	 - COnverts int to string
	 - returns one value - string formed with the given integrer.
	2. Atoi()
	 - converts string to integer
	 - returns 2 values - the corresponding integer, error (if any).
	*/
	var u int = 23
	var s string = strconv.Itoa(u)
	fmt.Printf("%q \n", s)

	var str string = "200"
	itr, err := strconv.Atoi(str)

	fmt.Printf("%v, %T \n", itr, itr)
	fmt.Printf("%v, %T \n", err, err)

	/* Constants - const
	const <Constant Name> <data type> = <value>
	1. Untyped constant - They are explicitly given a type at declaration.
	 - allow for flexibility
	2. Typed constants
	 - define here
	*/

	const age = 12 // Untyped
	const name string = "Harry"

}
