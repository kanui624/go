package main

import "fmt"

func main() {
	// Since we are specifying the type manually for go, go will infer the type at compile time
	var a = 4
	var b = 5.2



	fmt.Println(a, b)
	// Outputs:
	// 4 5.2

    // The line below will result in a type error since by this time in the code the type of these 
	// variables has already been infered
	// a = b
	// Outputs:
	// cannot use b (type float64) as type int in assignment
	
	// However, if we convert the type when re-assigning the variables go will execute
	a = int(b)
	// Outputs:
	// 5 5.2
	fmt.Println(a, b)

	// go is a strongly type language meaning once the type has been declared we will have to convert it
	// The below code will come out to an error
	// var x int 
	// x = "5"

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
	// Outputs:
	// 0 0  false
}

// Notes:
// Go zero values:
// numeric types: 0
// bool type: false
// string type: "" (empty string)
// pointer type: nil

