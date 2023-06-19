package main

import (
	"fmt"
	"reflect"
)

func main() {

	//ip := "127.0.0.1"

	// Use reflect to get name of ip
	//_ = reflect.ValueOf(ip)
	//name := reflect.TypeOf(ip).Elem().Name()
	//fmt.Println(name) // ip
	x := 123
	v := reflect.ValueOf(x)
	fmt.Println(v.Type().Name())   // 输出 "int"
	fmt.Println(v.Type().String()) // 输出 "int"
}
