package main

import (
	"fmt"
	"time"
)

func main() {
	jamChannel := make(chan struct{}, 10)
	for i := 0; i < 9; i++ {
		jamChannel <- struct{}{}
	}
	startTime := time.Now()
	for i := 0; i < 20; i++ {
		go readDb(jamChannel, startTime)
	}

	for {

	}

}

func readDb(jam chan struct{}, startTime time.Time) {
	jam <- struct{}{}

	endTime := time.Now()
	time.Sleep(time.Second * 1)
	fmt.Printf("finish %s time\n", endTime.Sub(startTime))
	<-jam

}
