package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func contFlow() {
	//defer : skips the flow of the program
	// puts the statement that is defered at the end of execution of program just
	// before return is called for the function
	// Defer functions are actually executed in LIFO order

	defer fmt.Println("First")
	defer fmt.Println("second")
	fmt.Println("Last")

	// defer to close the resources
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // closing the resource just after creating it and checking for err
	// that way we dont have worry to close after doing all the work. Defer will close the
	//resource
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(robots)

	//defer statement will take the arguments when its written not when its called
	// defer uses EAGER loading
	defEg := "start"
	defer fmt.Println(defEg) //Here defEg has already taken "start" as an argument for defEg
	defEg = "end"
	// Thus this will print "start" , event though print is called at last

	//PANIC and RECOVER functions
	//There is not concept of exceptions in go, instead there is panic
	//ACtaully most of things which other lang consider to be exception, is error for go,
	// like file not found
	//But sometimes we can run into panic situation, which is :
	ap, bp := 1, 0
	//ansp := ap / bp // this will generate panic:runtime error : .........
	fmt.Printf("tested ansp %v %v\n", ap, bp)

	//we can use built in panic function to throw panic
	fmt.Println("start panic")
	//panic("i am panicing")
	fmt.Println("end panic") //this is actually unreachable code due to panic
	//this will not be printed out , and function ends when paniced

	//
	//Order for execution is
	// first function is executed
	// then defer statement
	// then panic
	// then return is executed
	fmt.Println("Start ")
	defer fmt.Println("this was defered")
	//panic("I am panicing aaaa !!!!")
	fmt.Println("The END")

	checkPanicer()

}

func checkPanicer() {
	fmt.Println("Start my check")
	panicker()
	fmt.Println("Panic check ENDs ")
}

func panicker() {
	fmt.Println("Start Panicing in panicker ")
	// the err will have whatever error is returned from panic
	// check if that err is present in function execution and handle that error
	// now we can decide what we want to do with that error
	// either to continue parent functoin execution or not

	// PRoper place to use recover function is defer function
	// because when application panics , it will stop execution, it will only execute defer functions
	defer func() {
		if err := recover(); err != nil { // recover seems like a catch for panic
			fmt.Println("Recovered from Painc : ", err)
			panic("I will panic and stop everything") // this will stop the function
		}

	}()
	panic("Panicer panicing ....!!")
	//fmt.Println("Stop Panicking")
}

func mainForWebApp() {

	//handleFunc will register a function listner for url : "/"
	//the annonymous function is callback function , which is called when hit the url
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	//Listen and Serve function will return error
	//for go its not panic , it's erro
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error()) //running second time with same port will give an error
		// we explicitly decide to throw panic here when we get that error
	}

}
