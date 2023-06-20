/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 2023/5/10
  @desc: //TODO
**/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ioOutput() {
	for {
		fmt.Println("data")
		time.Sleep(1 * time.Second) // Sleep 1 second
		println(time.Now().Date())
		break
	}
}
func main() {
	wg.Add(1) // Add goroutine
	go func() {
		ioOutput()
		wg.Done() // Notify when /done
	}()

	wg.Wait() // Wait for goroutine to finish
}
