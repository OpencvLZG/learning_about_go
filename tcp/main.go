package main

import (
	"bytes"
	"log"
	"net"
	"regexp"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":6677")
	// ...
	if err != nil {

	}
	for {
		client, err := listener.Accept()
		// ...

		// ...
		data := make([]byte, 1024)
		_, err = client.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		//println("client.read")
		//println(string(data))
		// Read request fully
		//request, err := bufio.NewReader(client).ReadBytes('\n')
		//if err != nil {
		//	log.Fatal(err)
		//}
		//_, err = client.Write(request)
		// Edit User-Agent header
		//request = bytes.Replace(request, []byte("GET"), []byte("User-Agent: MyProxy "), 1)
		////request = bytes.Replace(request, []byte("User-Agent: "), []byte("User-Agent: MyProxy "), 1)
		//println("bufio.NewReader")
		//fmt.Printf("%s", request)
		//_, err = client.Write(request)
		//if err != nil {
		//	return
		//}
		lines := bytes.Split(data, []byte("\n"))
		data = bytes.Replace(data, []byte("User-Agent: "), []byte("User-Agent: MyProxy "), 1)
		hostStr := ""

		for _, line := range lines {
			if bytes.HasPrefix(line, []byte("Host:")) {
				// Split Host line by :
				host := bytes.SplitN(line, []byte(":"), 2)

				// host[1] is the host value
				hostStr = string(host[1][1:])
				hostStr = strings.Trim(hostStr, "\r\n")

				portReg := regexp.MustCompile(`(:)\d{2,5}`)
				port := portReg.FindString(hostStr)

				switch port {
				case ":443":

				case "":
					hostStr += ":443"
				default:
					//hostStr += port
				}

			}
		}
		//println(hostStr)

		if err != nil {
			log.Fatal(err)
		}
		_, err = net.Dial("tcp", hostStr)
		_, err = client.Write([]byte("HTTP/1.1 200 Connection Established \r\n\r\n"))
		// Write edited request
		//println(string(data))
		method := bytes.Split(data, []byte(" "))
		println(string(method[0]))
		//for _, me := range method {
		//	println(string(me))
		//}
		//if string(method[1]) != "CONNECT" {
		//	go server.Write(data)
		//	go io.Copy(server, client)
		//} else {
		//	go io.Copy(server, client)
		//	go io.Copy(client, server)
		//}
		//

		// Proxy remaining data

	}
}
