package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Start typing, please.")
	io.Copy(os.Stdout, os.Stdin)
}
