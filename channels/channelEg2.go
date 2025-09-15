package main

import (
	"math/rand"
	"fmt"
	"time"
)

// func emailSender(emailChan chan string,done chan bool){
// 	defer func() { done <- true }() //signal that this goroutine is done

// 	for email:= range emailChan {
// 		fmt.Printf("sending email to", email)
// 		time.Sleep(time.Second * 1) //simulate some work
// 	}
// }

func main(){
	emailChan:=make(chan string,10) //buffered channel
	done:=make(chan bool)

	go emailSender(emailChan,done) //start goroutine to process numbers

	for i:=0; i<10; i++{
		emailChan <- fmt.Sprintf("user%d",rand.Intn(1000)) //send email to channel
	}
	//this is important to close the channel after sending all emails
	//otherwise the emailSender goroutine will block forever waiting for more emails
	//and the program will never exit
	close(emailChan) //close the channel to signal that no more values will be sent
	fmt.Println("Main goroutine is exiting")
}