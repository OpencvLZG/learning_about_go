package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"regexp"
	"strings"
)

type Pipe struct {
	Reader io.Reader
	Writer io.Writer
}

func (p *Pipe) Read(b []byte) (n int, err error) {
	return p.Reader.Read(b)
}

func (p *Pipe) Write(b []byte) (n int, err error) {
	return p.Writer.Write(b)
}

func main() {
	listener, err := net.Listen("tcp", ":6677")
	// ...
	if err != nil {

	}
	for {
		client, err := listener.Accept()

		data := make([]byte, 1024)
		_, err = client.Read(data)
		if err != nil {
			log.Fatal(err)
		}

		//reader, writer := io.Pipe()

		//go io.Copy(client, *middleReader)
		//hostStr := getHost(data)
		//
		//data = bytes.Replace(data, []byte("User-Agent: "), []byte("User-Agent: MyProxy "), 1)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//server, err := net.Dial("tcp", hostStr)
		//_, err = client.Write([]byte("HTTP/1.1 200 Connection Established \r\n\r\n"))
		//// Write edited request
		////println(string(data))
		//method := bytes.Split(data, []byte(" "))
		//println(string(method[0]))

	}
}

func getHost(lines []byte) string {
	hostStr := ""
	linesBy := bytes.Split(lines, []byte("\n"))
	for _, line := range linesBy {
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
	return hostStr
}
