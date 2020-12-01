package main

import (
	"fmt"
)

func main() {
	// int age = 10; => c++
	// var age int = 30; // => go
	// In go we can omit the type and not get an error
	var ageTwo = 30
	fmt.Println("Age:", ageTwo)

	var name = "Marley"
	fmt.Println("Your name is:", name)
	// The blank operator
	// _ = name

	s := "learning golang"
	fmt.Println(s)

	// with this syntax we can declare more than one variable in a single line
	car, cost := "Honda", 20000
	fmt.Println(car, cost)

	// We can declare muliple variables at once with the syntax below as well
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int

	fmt.Println(a, b, c)

	// After declaring variables and their types we can simple reference their
	// variable name and set it equal to an actual value. This is multiple
	// assignments
	var i, j int
	i, j = 5, 8

	fmt.Println(i, j)
	// Outputs:
	// 5 8

	// The shorthand syntax to swap variables and thier values is:
	i, j = j, i

	fmt.Println(i, j)
	// Outputs:
	// 8 5

	// The below code is a short expression the a variable declaration
	sum := 5 + 2.3 
	fmt.Println(sum)
	// Outputs:
	// 7.3
}

// Notes:
// - When declaring variables we can omit the type declaration and it will still
// 	 work
// - All declared variables must be used. When trying to run or complie go code when
// 	 there are unused variables we get a complilation error
// - if you want to keep the variable declaration and not used it at first you can
//   utilize the "blank" operator which is implemented by setting an underscore (_)
//   equal to the variable name you want go to ignore at compilation. The blank
//   operator should be used with caution
// - Another way of declaring variables is (:=). This is the short declaration
//   version and only works in block scope. We use this when declaring new
//   variables or at least one new variable. We cannot use := for already defined
//   variables
