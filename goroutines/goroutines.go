package main

import (
	"fmt"
	"sync"
	"time"
)

//🌟Goroutine is a lightweight thread managed by the Go runtime
//go keyword is used to create a goroutine
//syntax: go functionName(params)
//eg: go task(1)

//🌟WaitGroup is used to wait for a collection of goroutines to finish
//we can use Add, Done, Wait methods of WaitGroup
//Add method is used to add the number of goroutines to wait for
//Done method is used to indicate that a goroutine has finished its work
//Wait method is used to block until all goroutines have finished their work

func task(id int, w *sync.WaitGroup){
	//why passing function because we wanted to done signal when the goroutine completes its work
	defer w.Done() //decrement the counter when the goroutine completes its work
	fmt.Printf("Task %d is starting\n",id)
}
func task2(id int){
	fmt.Printf("Task %d is starting\n",id)
}

func main(){
	T1:=time.Now()
	var wg sync.WaitGroup

	for i:=0;i<=25000;i++{
		wg.Add(1) //increment the counter for each goroutine
		go task(i, &wg) //go keyword is used to create a goroutine, a goroutine is a lightweight thread managed by the Go runtime
	}

	// time.Sleep(2*time.Second) //sleep is used to wait for a goroutine to finish, here we are waiting for 2 seconds before main goroutine exits
	// fmt.Println("Main goroutine is exiting")

	wg.Wait() //wait for all goroutines to finish, this will block the main goroutine until the counter is zero
	T2:=time.Now()
	diff:=T2.Sub(T1) //takes approx 1.02s to execute 25000 goroutines
	fmt.Println("Time taken to execute all goroutines is ",diff)

	// for i := 0; i < 25000; i++ {
	// 	task2(i)
	// }
	// T2:=time.Now()
	// diff:=T2.Sub(T1)
	fmt.Println("Time taken to execute all tasks is ",diff) //takes approx 0.012ms to execute all tasks
}