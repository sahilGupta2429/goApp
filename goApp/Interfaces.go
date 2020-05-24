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

	//we can not call any method on it.
	//Now, Everything can be cast into an object that has no method on it, even integers
	//int has no method so it can be cast to an empty interface
	ok := myObj.(io.Reader)
	fmt.Println(ok)

	// if ok {
	// 	fmt.Println(r)
	// } else {
	// 	fmt.Println("Conversion Failed")
	// }

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
