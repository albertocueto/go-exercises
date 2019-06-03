package main

import "fmt"

/*
Use var to DECLARE three VARIABLES. The variables should have package level scope.
Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
identifier “x” type int
identifier “y” type string
identifier “z” type bool
in func main
print out the values for each identifier
The compiler assigned values to the variables. What are these values called?
*/

/*
Using the code from the previous exercise,
At the package level scope, assign the following values to the three variables
for x assign 42
for y assign “James Bond”
for z assign true
in func main
use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
print out the value stored by variable “s”

*/

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%T: %v\t%T: %v\t%T: %v\n", x, x, y, y, z, z)
	fmt.Println(s)
}
