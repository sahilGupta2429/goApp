package main

import (
	"fmt"
	"strconv"
)

// Block variable declaration at package level , can not use ':='
//  This can only be used in a function

/* var (
	hero   string = "thor"
	number int    = 2
) */

//If variable is declared at the package level with lower case letter then its package level scoped
//else if the vairable is declared with capital letter its scoped OUTSIDE package level

func main() {

	// var i int
	// i = 42
	// var i int = 42
	//i := 44
	// variable declared should be used
	//fmt.Printf("%v,%T", hero, hero)

	var i int = 42
	fmt.Printf("%v %T\n", i, i)

	var j float32
	j = float32(i)

	//float32() , int() are casting function
	// but in go we can not asign one type variable to another , we actually have to cast it
	// j = i (this is not allowed)
	fmt.Printf("%v %T\n", j, j)

	var k string
	k = string(i)
	fmt.Printf("%v %T\n", k, k) //here string(i) converts 42 to unicode value for *

	// To actually convert 42 i.e. i into string we use string convertion packages
	var l string
	l = strconv.Itoa(i)
	fmt.Printf("%v %T\n", l, l)

	a := 1 == 1
	b := 2 == 2
	fmt.Printf("%v %v \n", a, b)

	var c bool
	fmt.Printf("%v\n", c) //without initializing 0 is asigned and boolean value of 0 is false

	var u uint = 12
	fmt.Printf("%v %T\n", u, u)

	//arithmatic operations
	p := 10
	q := 3
	fmt.Println(p + q)
	fmt.Println(p - q)
	fmt.Println(p / q)
	fmt.Println(p * q)
	fmt.Println(p % q)

	//go will NOT implicit type convert for you
	// thus (a+b) , where a and b are of diff Data types it will not work

	//Bit operators
	//10 = 1010
	//3 = 0011
	fmt.Println(p & q)  //AND    //0010 = 2
	fmt.Println(p | q)  //OR     //1011 = 11
	fmt.Println(p ^ q)  //EX-OR  //1001 = 9
	fmt.Println(p &^ q) //AND-NOT operator //0100 = 8

	//Bit Shifting
	m := 8              // 8 = 2^3
	fmt.Println(m << 3) //Bit shift left 3 places  = 2^3 * 2^3 = 2^6 = 64
	fmt.Println(m >> 3) //Bit shift right 3 places = 2^3 / 2^3 = 2^0 = 1

	//Floating point numbers
	n := 3.14
	r := 13.7
	s := 2.1E14
	//If you are using initializer syntax , everything is working as float64
	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T\n", r, r)
	fmt.Printf("%v %T\n", s, s)

	// Floating point numbers don't have % , Bitwise and bitshifting operators
	fmt.Println(n + r)
	fmt.Println(n - r)
	fmt.Println(n * r)
	fmt.Println(n / r)

	//Complex numbers //default is complex128
	x := 1 + 2i
	g := complex(3, 4) //complex function ->  3 is real and 4 is imaginary
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", real(x), real(x))
	fmt.Printf("%v %T\n", imag(x), imag(x))
	fmt.Printf("%v %T\n", g, g)

	e := 1 + 2i
	f := 2 + 5.2i
	fmt.Println(e + f)
	fmt.Println(e - f)
	fmt.Println(e * f)
	fmt.Println(e / f)

	//String
	h := "My name is Slim Shady..."
	uv := "Slim Shady"
	fmt.Printf("%v %T\n", h, h)
	fmt.Printf("%v %T\n", h[3], h[3]) // String can be treated as an array of unsigned integer
	// h[3] = "i"  // this is not allowed string is immutable

	//Arithmatic opertation with string (string concatination)
	fmt.Printf("%v\n", h+uv)

	//Converting string to collection of bytes (called as slice of bytes)
	ab := "This is test"
	cd := []byte(ab)
	fmt.Printf("%v %T\n", cd, cd) // this comes as array of ascii values or UTF values of the string

	//RUNE data type  Rune represents any UTF32 bit charactor
	var ac rune = 'a' //rune is declared with single quotes
	//Rune is true type alias, when you talk about rune you are talking about integer32
	fmt.Printf("%v %T\n", ac, ac)

	test()
	arrSl()
	testMapStruct()
	controlSt()
	loopPrac()
	contFlow()
	pointerStudy()
	functionStudy()
	studyInterfaces()
	studyConcurrancy()
}
