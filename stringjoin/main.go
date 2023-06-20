/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 2023/5/11
  @desc: //TODO
**/

package main

import "strings"

func main() {
	var string1 string
	string1 = "a"
	string2 := "b"
	var string3 string = "c"
	println(string1)
	println(string2)
	println(string3)
	var stringList []string
	stringList = append(stringList, string1)
	stringList = append(stringList, string2)
	stringList = append(stringList, string3)

	for _, s := range stringList {
		println(s)
	}
	println(strings.Join(stringList, "-"))

}
