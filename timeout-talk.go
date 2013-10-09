package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func stdinChannel() <-chan []byte {
	ch := make(chan []byte)
	go func() {
		defer close(ch)
		for {
			buf := make([]byte, 1024)
			amt, err := os.Stdin.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Fatal("Error reading from stdin:", err)
				}
				break
			}
			ch <- buf[:amt]
		}
	}()

	return ch
}

func main() {
	ch := stdinChannel()

	fmt.Println("Tell me something.")
	select {
	case res, ok := <-ch:
		if !ok {
			fmt.Println("You said NOTHING?! QQ")
		} else {
			fmt.Println("Thanks for your swift answer of", string(res))
		}
		break
	case <-time.After(3 * time.Second):
		fmt.Println("\rYou're going to need to type faster than that...")
		break
	}
}
