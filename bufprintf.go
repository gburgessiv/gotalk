package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "Stop! %s time!", "hammer")
	fmt.Println("Buffer says:", string(buf.Bytes()))
}
