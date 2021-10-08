package chaos_test

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/chyroc/chaos"
)

// TeeWriter example
func Example_TeeWriter() {
	buf := new(bytes.Buffer)
	w := chaos.TeeWriter([]io.Writer{buf, os.Stdout}, nil)

	w.Write([]byte("Hi, World."))

	fmt.Println(buf.String())

	// output:
	// Hi, World.Hi, World.
}
