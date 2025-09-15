package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	//dont lock whole function only lock whre the modification will go on
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1
}


func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)  //coz this wasnt blocking code so main goroutine was exiting before all the goroutines were finished

	}

	wg.Wait()//wait for all goroutines to finish
	fmt.Println(myPost.views)
}
