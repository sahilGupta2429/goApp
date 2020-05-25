package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//Creating Routines
//Syncronization : Wait Groups and Mutexes
// Parallelism

// Best Practices
//dont create goroutines in libraries
//let consumer control concurrancy
//when you create a go rountine, know how it will end
//avoid subtle memory loss
//check for race confition at compile time

var wg = sync.WaitGroup{}
var countr int = 0

var mu = sync.RWMutex{} //Read Write Mutex
//Simple mutex is lock unlock mutex sync.Mutex
//As many routines can read RWMutex but only on can write at a time
//If anything is reading, we can't write at all.

func studyConcurrancy() {
	//main function which calls studyConcurrancy function is itself another go routine
	//So as soon as it called another go routine with sayHello(), it will stop and nothing
	// will be printed from sayHello()
	//go sayHello()
	//However when we sleep the main go routine it sleeps and go routine for sayHello will print
	//Hello Routine instead
	time.Sleep(100 * time.Millisecond)

	//go routine can take argument from parent go routine, but its not reccomended to create
	//such dependency
	//The reson it works is because go has concept of closures, which means
	//func has access to var in outer scope
	var msg = "hello"
	go func() {
		fmt.Println(msg)
	}()

	go func(msg string) {
		fmt.Println("1" + msg)
	}(msg)
	msg = "bye"
	time.Sleep(100 * time.Millisecond)

	//We don't want to use sleep, so alternative is waitgroup
	var ewg = "Example Wait group"
	//syncronyze main go routine to another goroutine, we start that by adding another goroutine
	//the main function is added as goroutine 0
	wg.Add(1)
	go func() {
		fmt.Println(ewg)
		//we tell the wait group, that go routine is done
		wg.Done()
	}()
	//exit the application by doing waiting on wait group
	wg.Wait()

	//We can change the number of thread by
	runtime.GOMAXPROCS(1)
	//This will tell the number of OS threads that are available for number of cores in CPU
	fmt.Printf("Threads : %v \n", runtime.GOMAXPROCS(-1))

	//This will print random counter value where both go routines fight to completion parllelly
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mu.RLock()
		go sayHello()
		mu.Lock()
		go increment()
	}
	//Locking the mutexes in a single context, before launching go routine
	// and asyncronously unclocking them in routines once  I am done with it
	wg.Wait()

	//Mutex : Its a lock that application is gonna honor

}

//Go Routine thread is diff from other threads in other lang
// There threads used are basically OS thread and are very expensive to work on and create and
// destroy those threads
// But go routine thread are basially abstraction over those low level threads and are thus
// cheap and require very less resources to run, create and destroy the threads.
func sayHello() {
	fmt.Println("Hello Routine" + strconv.Itoa(countr))
	mu.RUnlock()
	wg.Done()
}

func increment() {
	countr++
	mu.Unlock()
	wg.Done()
}
