package main

import (
	"fmt"
	"io"
	"log"
	"time"
)

func main() {
	reader, writer := io.Pipe()
	data := make([]byte, 1024)
	go func() {
		for {
			_, err := reader.Read(data)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(string(data))
		}

	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			total, err := writer.Write([]byte("hello world"))
			fmt.Printf("total write %d byte\n", total)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	for {

	}
}
