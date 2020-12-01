package main

import "fmt"

func main() {
	// int age = 10; => c++
	// var age int = 30; // => go
	// In go we can omit the type and not get an error
	var ageTwo = 30; 
	fmt.Println("Age:", ageTwo)

	var name = "Marley"
	fmt.Println("Your name is:", name)
	// The blank operator
	// _ = name

	s := "learning golang"
	fmt.Println(s)

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