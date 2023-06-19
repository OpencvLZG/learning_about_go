/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 2023/4/24
  @desc: //TODO
**/

package main

func main() {
	dataChan := make(chan int, 2)
	go func() {
		dataChan <- 123
	}()

	//dataChan <- 1243
	//dataChan <- 12
	n := <-dataChan
	//n = <-dataChan
	println(n)

}
