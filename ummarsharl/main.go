package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Server struct {
	ServerIp       string `json:"server_ip"`
	ServerPort     string `json:"server_port"`
	ServerMethod   string `json:"server_method"`
	ServerProtocol string `json:"server_protocol"`
}

type User struct {
	Name string
	Age  int
}

func main() {
	btyes, err := os.ReadFile("./ummarsharl/config.json")
	if err != nil {
		log.Fatalln(err)
	}
	var server Server
	err = json.Unmarshal(btyes, &server)
	if err != nil {
		log.Fatalln(err)
	}
	//for k, v := range server {
	//	fmt.Printf("+%s: %v\n", k, v)
	//}
	fmt.Printf("%+v\n%v", server)
	//data := `{"Name": "John", "Age": 30}`

	//var user User
	//json.Unmarshal([]byte(data), &user)
	//
	//for key, value := range user {
	//	fmt.Printf("%s: %v\n", key, value)
	//}
}
