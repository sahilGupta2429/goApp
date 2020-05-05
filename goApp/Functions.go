package main

import "fmt"

func functionStudy() {
	sayMessage("Hello functions")
	name := "sahil"
	sayGreetings("Hello", name)
	fmt.Println(name)

	greet := "Good Morning"
	sayRefGreetings(&greet, &name)
	fmt.Println(name)

	sum("The sum is ", 1, 2, 3, 4, 5)
	//1,2,3,4,5 goes into function argument as a slice
	// The variatic paramenter should be added as last argument and
	//function can only have one variatic argument

	//return values
	fmt.Println(returnTest1("Ravish"))
	fmt.Println(*returnTestPointer("Akshay"))
	fmt.Println(returnTestNamedReturn("Chintu"))

	anymFunc()
	funcAsVar()

	//method calling
	//creat struct instance as g and assign it values
	g := greeter{
		greeting: "Hello",
		name:     "GO",
	}
	//call greet method defined on the structure
	g.greet()
	//the stucture passed is passed by value

}

func sayMessage(msg string) {
	fmt.Println(msg)
}

//passing in value (copy value) in function argument
func sayGreetings(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}

//passing in pointer instead of values in a function argument
//it can change the value on calling scope as well
func sayRefGreetings(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "shivangi"
	fmt.Println(*greeting, *name)
}

//Using variatic parameters as arguments in go function
func sum(msg string, values ...int) {
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

//A simple return statement with string return value
func returnTest1(name string) string {
	return "Hello" + name
}

// we can return pointers in go
// when we return the pointer and function is over it mantains the data in heap for that return value
func returnTestPointer(name string) *string {
	greet := "Hello" + name
	return &greet
}

// named return parameter, return value is declared at function start
func returnTestNamedReturn(name string) (greeting string) {
	greeting = "Hello" + name
	return
}

func anymFunc() {
	//i has scope of anymFunc , but if we call anonymous function without argument
	// then in case of async call it will show wierd result
	//Thus, passing an argument i will ensure the correct call of anonymous function
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}

// function as a variable

func funcAsVar() {
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Can not divide by 0")
		}
		return a / b, nil
	}
	d, err := divide(10, 2)
	d1, err1 := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(d, d1)
}

type greeter struct {
	greeting string
	name     string
}

//Methods
//This is a method, greet() method is refrence to greeter struct
// the greeter struct as g will be available in that method
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
