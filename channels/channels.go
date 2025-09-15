package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// 🔹 What is a channel?
// 	   A channel in Go is like a pipe that allows one goroutine to send values and another goroutine to receive them.

// processNum takes a channel of integers as an argument.
// numChan is of type chan int → a channel that carries integers.
// Inside, it waits (<-numChan) to receive a value from the channel.
// When it gets a value, it prints: "Received number: <value>".
// ⚡ Note: <-numChan blocks until a value is available in the channel.
func processNum(numChan chan int) {//send to another goroutine
	// fmt.Println("Received number:", <-numChan) // receive value from channel
	for num:= range numChan { //range over channel until its closed
		fmt.Println("Received number:", num)
		time.Sleep(time.Second * 1) //simulate some work
	}
}

//recieve to main goroutine 
func sum (result chan int, num1 int, num2 int){
	result <- num1+num2 //send value to channel
}

func emailSender(emailChan<-chan string,done chan bool){ 
	//emailChan is a receive-only channel //done is a bidirectional channel// if done to be send only then chan<- bool
	defer func() { done <- true }() //signal that this goroutine is done

	for email:= range emailChan {
		fmt.Printf("sending email to %s\n", email)
		time.Sleep(time.Second * 1) //simulate some work
	}
}

func main() {
	// Create a new unbuffered channel for integers.
	// Since it’s unbuffered, every send (ch <- val) must be paired with
	// a receive (<-ch) happening at the same time — otherwise it blocks.

// 	👉 This creates an unbuffered channel, meaning it has no storage capacity.
// That means:
// If one goroutine tries to send (ch <- value) and no other goroutine is waiting to receive (<-ch), the sending goroutine will block (pause and wait).
// If one goroutine tries to receive (<-ch) and no other goroutine has sent anything yet, the receiving goroutine will also block (pause and wait).

// So send and receive must happen at the same time.
// This is why unbuffered channels are often used for synchronization.
	numChan := make(chan int)

	// Start a new goroutine that runs processNum(numChan).
	// This goroutine will block inside fmt.Println("... <-numChan")
	// until it receives a value.
	go processNum(numChan)

	// Send value to channel. Channels are blocking by default.
	// Since processNum goroutine is already waiting (<-numChan),
	// the value is delivered immediately, so no deadlock occurs.
	// If the goroutine wasn’t running, this line would block forever (deadlock).
	// numChan <- 42
	// // Keep sending random numbers
	// for i:=0; i<1; i++{
	// 	numChan <- rand.Intn(100)
	// }

	// Sleep for 1 second to allow goroutine to finish before main exits.
	// time.Sleep(time.Second * 1)
	// close(numChan) //close the channel to signal that no more values will be sent
	// fmt.Println("Main goroutine is exiting")


	// result:=make(chan int)

	// go sum(result,3,4)

	// res:=<-result //receive value from channel
	// fmt.Println("Sum is:",res)


	emailChan:=make(chan string,10) //buffered channel
	done:=make(chan bool)

	go emailSender(emailChan,done) //start goroutine to process numbers

	for i:=0; i<10; i++{//sprintf to format string
		emailChan <- fmt.Sprintf(" %d@gmail.com",i) //send email to channel
	}
	fmt.Println("All emails sent to channel")

	close(emailChan) //close the channel to signal that no more values will be sent
	<-done //wait for emailSender goroutine to finish
	fmt.Println("All emails processed, main goroutine is exiting")


	// accepting data from multiple channel 
	channel1:=make(chan string)
	channnel2:=make(chan int)

	// goroutine 1
	go func(){
		channel1 <- "hello"
	}()

	go func(){
		channnel2 <- 42
	}()

	// main goroutine
	select{
	case chan1val:=<-channel1:
		fmt.Println("received from channel 1:",chan1val)
	case  chan2val:=<-channnel2:
		fmt.Println("received from channel 2:",chan2val)	

	}

	//🌟Note: Channels are used for communication between goroutines.
//Blocking code below:
	// messageChan:=make(chan string)
	// messageChan<-"Hello, World!" //send value to channel ,channels are blocked by default so this will cause deadlock
	// message:=<-messageChan //receive value from channel
	// fmt.Println(message)
}