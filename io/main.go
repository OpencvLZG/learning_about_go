package main

import (
	"fmt"
	"io"
	"os"
)

// Pipe - Implements io.Reader and io.Writer
type Pipe struct {
	Reader io.Reader
	Writer io.Writer
}

func (p *Pipe) Read(buf []byte) (n int, err error) {
	return p.Reader.Read(buf)
}

func (p *Pipe) Write(buf []byte) (n int, err error) {
	return p.Writer.Write(buf)
}

// Client
func client(pipe *Pipe) {
	io.Copy(pipe, os.Stdin) // Copy from stdin to pipe
}

// Server
func server(pipe *Pipe) {
	io.Copy(os.Stdout, pipe) // Copy from pipe to stdout
}

func main() {
	// Create a pipe
	reader, writer := io.Pipe()
	
	pipe := &Pipe{Reader: reader, Writer: writer}

	// Start client and server
	go client(pipe)
	go server(pipe)

	// Wait for input
	var input string
	fmt.Scanln(&input)
}
