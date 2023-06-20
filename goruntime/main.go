/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 2023/4/24
  @desc: //TODO
**/

package main

import "fmt"

func main() {

	foods := []string{"asda", "1111", "asd", "asd", "asd"}
	for i := range foods {
		fmt.Println(i)
	}

	something := []string{"meat", "fish", "bag", "pen", "esp"}
	for i, s := range something {
		result := fmt.Sprintf("%d--%q", i, s)
		fmt.Println(result)
		//print(i)
		//println(s)
	}
}
