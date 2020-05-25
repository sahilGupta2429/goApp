package main

import (
	"fmt"
	"sync"
	"time"
)

//Channel basics
//Restricting flow
//buffered channels
//closed channels
//for range loops with channels
//select statements

//to sync the diff go routines
var chWg = sync.WaitGroup{}

//
const (
	logInfo    = "INFO"
	logWarning = "WWARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) //0 memory allocation for creating empty struct
//This is called signal only channel

func studyChannels() {
	//channels to sync data flow between these go routines
	//channel can be made only with make functions, its complicated and many things are going on
	// in background for that
	//channel is strongly types, thus it only takes integers
	ch := make(chan int, 50) //adding argument 50 makes buffered channel from unbuffered channel
	chWg.Add(2)
	//Recieving channel go routine, pulling/recieving data from channel
	//To make sure you only make single type of go routine, i.e. recieving or outgoing
	//we add arguments recieveing or sending channels
	//Data is going out of the channel, thus recieve only channel go routine
	go func(ch <-chan int) {
		//i := <-ch      //recieving i from the channel,
		//fmt.Println(i) // printing the value recieved from channel i.e. 43 (IT PASSES BY VALUE)

		//There can be infinine number of elements in a channel, so either close the channel
		//Or the code will panic, or check for error in for look.
		for i := range ch {
			fmt.Println(i)
		}
		chWg.Done()

		//This channel will recieve the value send into the channel and then sleep after Done()
		//Thus only one value will be processed for sending value into the channel
	}(ch)
	//Sending channel go routine, sending data into the channel
	//Data is going into the channel, thus send only channel go routine
	go func(ch chan<- int) {
		i := 43
		ch <- i //sending i into the channel
		//Only One message can be in the channel at one time, This is
		//BY DEFAULT we are working with unbuffered channels
		i = 34
		ch <- i
		close(ch) //Closing the channel, we are done working with channel
		chWg.Done()
	}(ch)
	chWg.Wait()

	//select statements - Logger

	//Without shutting down this go routine, this channel in will shut forcefully and
	//shut this go routine
	go logger() //launching the logger go routine

	//To shut down the go routine there are several ways
	// 1. Using defer function
	// defer func() {
	// 	close(logCh)
	// }()

	logCh <- logEntry{time.Now(), logInfo, "App is Starting"}

	logCh <- logEntry{time.Now(), logInfo, "App Ends"}
	time.Sleep(100 * time.Millisecond)

	//Pass an empty struct into the channel and thus closing the channel
	doneCh <- struct{}{}
}

func logger() {
	// for entry := range logCh {
	// 	fmt.Printf("%v - [%v] %v \n", entry.time, entry.severity, entry.message)
	// }

	//select statements work like switch statements, but only work in context of channels
	//They allows goroutine to monitoe several channels at once
	//They block if all channels block,
	//So when there is no msg on channel it will block by default and when message comes it executes
	//on the proper case
	//If multiple channels recive value simultaneously, behaviour is undefined
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v \n", entry.time, entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}

}
