package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mu sync.Mutex

func main() {

	for i := 0; i < 1000; i++ {
		go incr()
	}

	time.Sleep(time.Second)

	fmt.Println(count) // 输出可能小于1000

	//select {}
	syChannel := make(chan struct{})
	select {
	case <-syChannel:

	}
	println()
}

func incr() {
	defer mu.Unlock()
	mu.Lock()
	count++
	//for {
	//	println(1)
	//}
}
