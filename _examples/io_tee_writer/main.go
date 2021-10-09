package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/chyroc/chaos"
)

func main() {
	buf := new(bytes.Buffer)
	w := chaos.TeeWriter([]io.Writer{buf, os.Stdout}, nil)

	w.Write([]byte("Hi, World."))

	fmt.Println(buf.String())

	// output:
	// Hi, World.Hi, World.
}
