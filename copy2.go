package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	buf := bytes.NewBuffer(nil)
	speshulReader := io.TeeReader(os.Stdin, buf)
	fmt.Println("Type anything!")
	io.Copy(os.Stdout, speshulReader)

	fmt.Println("Thanks for visiting! Here's a record of everything you typed, courtesy of the NSA:")
	buf.WriteTo(os.Stdout)
}
