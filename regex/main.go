/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 2023/6/11
  @desc: //TODO
**/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	url := "http://www.youtube.com:123131231/"
	found, err := regexp.MatchString("g([a-z]+)ng", "golang")
	fmt.Printf("found=%v, err=%v \n", found, err)

	reg := regexp.MustCompile(`(?i)(http://|https://|\/|:([0-9]+))`)
	fmt.Printf("%q\n", reg.ReplaceAllString(url, ""))

	//host :=
	//found, err := regexp.MatchString("g([a-z]+)ng", "golang")
	//fmt.Printf("found=%v, err=%v \n", found, err)

}
