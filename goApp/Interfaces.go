package main

import (
	"bytes"
	"fmt"
	"io"
)

// basics
// composing interfaces
// type conversion
// empty interfaces
// type switches
// implementing with values vs pointerStudy

// best practices
// Use small, many interfaces eg: io.Reader, io.Writer, interface{}
// Don't export inteface for types that will be consumed
// Don't export interfaces for types that will be used by package
// Design functions and methods to recieve interfaces whenever possible

func studyInterfaces() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Interface"))

	myInt := IntCounter(0)       //Create an IntCounter and cast an integer to that IntCounter i.e. 0
	var inc Incrementer = &myInt //Create Incrementer and assign to pointer object of myInt/
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}

	wc.Write([]byte("Hello"))
	wc.Close()

	var wc1 WriterCloser = &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}

	wc1.Write([]byte("Hello hi how r "))
	wc1.Close()

	//now bwc1 has memory address of bufferedWriterCloser
	//now the type conversion successed, therefore I can work with this as a BufferedWriterCloser not
	// as WriterCloser ..
	//If i wanted to access buffer directly, ill be able to do that directly, but as WriterCloser it's not
	// aware of internal fields of specific implementation, so i won't be able to access it
	bwc1 := wc.(*BufferedWriterCloser)
	fmt.Println(bwc1)

	//bwc1 := wc.(io.Reader)  io.Reader is interface and requires Read method on it
	// This is not possible and it panics
	//interface conversion : *main.BufferWriterCloser is not io.Reader : missing method Read

	//empty interface
	//interface which has no method in it. interface{},
	var myObj interface{} = &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}

	//Type conversion of myObj empty interface to WriterCloser INterface and calling Write and
	// Close functions on it
	if wcObj, OK1 := myObj.(WriterCloser); OK1 {
		fmt.Println("Success on Writer Closer Conversion")
		wcObj.Write([]byte("Hello How r"))
		wcObj.Close()
	}
	//we can not call any method on it.
	//Now, Everything can be cast into an object that has no method on it, even integers
	//int has no method so it can be cast to an empty interface
	if ioObj, ok2 := myObj.(io.Reader); ok2 {
		fmt.Println("Success on io.Reader Conversion")
		ioObj.Read([]byte("Hi 2"))
	}

	// if ok {
	// 	fmt.Println(r)
	// } else {
	// 	fmt.Println("Conversion Failed")
	// }

	//Type Switches
	var iT interface{} = true
	//var iT interface{} = "0"
	//var iT interface{} = 0
	switch iT.(type) {
	case int:
		fmt.Println("Its interger")
	case string:
		fmt.Println("Its string")
	case bool:
		fmt.Println("its bool")
	default:
		fmt.Println("yo i dont know")
	}

	// creating value type interface object
	//When implementing an interface, if I use a value type the methods that implement the interface
	// have to all have value reciever, however if I am implementing the interface with a pointer
	//then o just have to have a method there regardless of the method type.
	//method set for pointer type is method set for value recirever as well as methods with pointer reciver
	var myWCValue WriterCloser = &myWrtierCloser{}
	fmt.Println(myWCValue)

}

//interfaces don't describe data , they describe behaviour
// if your interface has single method then, interface name should be methodName+'er'
type Writer interface {
	Write([]byte) (int, error)
}

//struct for implementing interfaces
type ConsoleWriter struct{}

//in go we don't explicitly implement interfaces , we will implicitly implement interfaces
//any type can have method associated with it, here type ConsoleWriter(struct) has method
// Write associated with it which implements Writer interface implicitly
// we can call Write method on Writer instance
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

/// if we need more than one than one method ,but we cannot decompose the interface down.

//Already declared writer interface
// type Writer interface {
// 	Write([]byte) (int, error)
// }

type Closer interface {
	Close() error
}

//Embedding interfaces with in interfaces
//WriterCloser in implemented if the object has Write and Close methods on it
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

//Method on bufferedWriterCloser struct Object
func (bwc *BufferedWriterCloser) Write(date []byte) (int, error) {
	return 80, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	return nil
}

/// Type to pointer and value convertion on interface
type myWrtierCloser struct{}

func (mwc *myWrtierCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWrtierCloser) Close() error {
	return nil
}
