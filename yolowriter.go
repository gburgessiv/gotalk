package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type YoloBuffer struct {
	buf *bytes.Buffer
}

func newYoloBuffer() *YoloBuffer {
	innerBuffer := bytes.NewBuffer(nil)
	return &YoloBuffer{innerBuffer}
}

func (y *YoloBuffer) Read(b []byte) (n int, err error) {
	a, e := y.buf.Read(b)
	return a, e
}

func (y *YoloBuffer) Write(b []byte) (n int, err error) {
	n, err = y.buf.Write(b)
	if err != nil {
		return
	}
	_, err = y.buf.Write([]byte(" ...YOLO!\n"))
	return
}

func main() {
	buf := newYoloBuffer()
	fmt.Fprint(buf, "Got ma swaq on")
	fmt.Fprint(buf, "Ain't no one messin wit me")
	fmt.Fprint(buf, "'Cause I'm a baws")
	io.Copy(os.Stdout, buf)
}
