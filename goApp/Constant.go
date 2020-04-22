package main

import (
	"fmt"
)

/// these are called enumarated constants
const (
	con1 = iota
	con2 = iota
	con3 = iota
)

const (
	acon = iota
	bcon
	ccon
)

const (
	con11 = iota
	con22
)

const (
	//with iota we have to start with 0 value, but if don't care about 0 then we don't have
	//to asign memory to it , thus '_'
	_ = iota //Underscore is write only value
	myvar1
	myvar2
	myvar3
)

//Arithmatic,bitwise,bitshift operations are allowed with constants

const (
	_  = iota             //  _ = 0
	KB = 1 << (10 * iota) // 1 << (10*1) => 1 << (10) => 2^0 * 2^10 => 2*10 => 1024bytes
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func test() {
	const myConst int = 3
	// myConst = 4    changing value of constant is not allowed
	fmt.Printf("%v %T\n", myConst, myConst)

	//const myConst2 int = math.Sin(1.57) // this will not work
	//to evaluate myConst2 we require to run the function math.Sin(1.57) , which
	// does not happen at compile time thus we can not assignt the constant
	//fmt.Printf("%v \n", myConst2)

	//There can be integer, string, float and boolean constants
	//they can also be shattered

	const var1 int = 23
	var var2 int = 12
	//Since constant and variable are of same type they can be added and result will be variable
	// Typed constants can work only with same type
	fmt.Printf("%v %T\n", var1+var2, var1+var2)

	const avar = 23
	var bvar int16 = 12
	//you would think this will not work coz avar is const of int8 and bvar is int16
	// but how comiler sees it is (23+bvar) =>  the compiler interprets 23 as int16 it doesnot
	// interpret constant
	//untyped constant can work with similar types (little more flexible)
	fmt.Printf("%v %T\n", avar+bvar, avar+bvar)

	//constant initialized to iota(special variable) works as a counter and another iota value is incremented accordingly
	fmt.Printf("%v\n", con1)
	fmt.Printf("%v\n", con2)
	fmt.Printf("%v\n", con3)

	//In this constant block it does not matter if constant is not initialized ,
	// any constant or variable followed by iota will be incremented as above constant block
	// because it sees a pattern of iota
	fmt.Printf("%v\n", acon)
	fmt.Printf("%v\n", bcon)
	fmt.Printf("%v\n", ccon)

	//another constant block when reinitialized with iota starts the new counter
	// thus iota constant scope is within the constant block
	fmt.Printf("%v\n", con11)
	fmt.Printf("%v\n", con22)

	fileSize := 40000000000.0
	fmt.Printf("%0.2fGB \n", fileSize/GB)

}
