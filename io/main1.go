package main

import (
	"io"
	"log"
)

type IoReadAndWrite struct {
	Reader io.Reader
	Writer io.Writer
}

func (w *IoReadAndWrite) Write(buf []byte) (n int, err error) {
	return w.Writer.Write(buf)
}
func (w *IoReadAndWrite) Read(buf []byte) (n int, err error) {
	return w.Reader.Read(buf)
}

func main() {

	reader, writer := io.Pipe()
	io_object := IoReadAndWrite{
		Reader: reader,
		Writer: writer,
	}

	go func() {
		writeData := []byte("hello world")
		_, err := io_object.Write(writeData)
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		readData := make([]byte, 1024)
		_, err := io_object.Read(readData)
		if err != nil {
			log.Fatal(err)
		}
		println(string(readData))
	}()
	for {
		break
	}

}
