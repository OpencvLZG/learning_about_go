/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 2023/5/11
  @desc: //TODO
**/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:], " ")
	fmt.Println(os.Args[0], " ")
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i], "")
	}
}
