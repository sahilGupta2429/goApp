package main

import "fmt"

func controlSt() {
	if true {
		fmt.Println("Hello")
	}

	//Initializer syntax
	// Allows statement to run i.e. initializer block to calculate boolean condition

	sstPop := map[string]int{
		"Haryana": 56262565,
		"Punjab":  547841254,
		"Delhi":   897264256,
	}
	// sstChildPop, ok := sstPop["Haryana"]   ->  this is initializer statement to cal ok
	// sstChildPop is only defined within the scope of if statement
	if sstChildPop, ok := sstPop["Haryana"]; ok {
		fmt.Println(sstChildPop)
	}

	//Equality operators
	eq1 := 20
	eq2 := 30
	if eq1 < eq2 {
		fmt.Println("First")
	}

	if eq1 > eq2 {
		fmt.Println("second")
	}

	if eq1 == eq2 {
		fmt.Println("third")
	}

	if eq1 <= eq2 {
		fmt.Println("fourth")
	}

	if eq1 >= eq2 {
		fmt.Println("fifth")
	}

	if eq1 != eq2 {
		fmt.Println("sixth")
	}

	// Logical Operators
	//OR
	if eq1 < 1 || eq2 > 100 {
		fmt.Println("OR Operator")
	}

	//AND
	if eq1 < 1 && eq2 > 100 {
		fmt.Println("AND Operator")
	}

	//NOT
	if !false {
		fmt.Println("In NOT oprator")
	}

	// if -else   ->  else if
	if false {
		fmt.Println("If else skip")
	} else if false {
		fmt.Println("else if skip")
	} else {
		fmt.Println("In Else")
	}

	// switch statement
	switch 2 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Nothing")
	}

	// IN go case can compare multiple values
	// But each case must be unique
	// you can not have 5 in both cases
	switch 5 {
	case 1, 5, 10:
		fmt.Println("1,5,10")
	case 2, 4, 6:
		fmt.Println("2,4,6")
	}

	/// switch can also have initializer code
	// here ini generates 5
	switch ini := 2 + 3; ini {
	case 5:
		fmt.Println("Check 5")
	case 2:
		fmt.Println("Check 2")
	case 3:
		fmt.Println("Check 3")
	}

	//Tagless syntax
	// The logic is with case
	//The case condition can overlap here
	// it will just run the code in first case that returns true
	icas := 10
	switch {
	case icas <= 10:
		fmt.Println("Check tagless 1")
		//fallthrough
	case icas <= 20:
		fmt.Println("Check tagless 2")
	default:
		fmt.Println("default tagless")
	}
	// switch statement has break keyword implied
	// Go has 'fallthrough' keyword, which will execute the code in next case weather or not
	// case is true or false

	// Type interface can take any type used in GO
	var typeSwitch interface{} = [3]int{}
	switch typeSwitch.(type) {
	case int:
		fmt.Println("INT")
		//break
		//fmt.Println("Break Check")
	case float64:
		fmt.Println("Float64")
	case string:
		fmt.Println("String")
	case [2]int:
		fmt.Println("Array of 2 size")
	default:
		fmt.Println("something else") //[2]int != [3]int
	}
	//If you want to break in between running any code in case
	// Then u can use break keyword
}
