package chaos

import (
	"io"

	"github.com/chyroc/chaos/internal/cio"
)

type TeeWriterOption struct {
	// When an error occurs when writing data in multiple writers, exit immediately,
	// instead of writing all the writers,
	// the default is false
	// 当在多个 writer 中写数据发生错误的时候，立刻退出，而不是将所有 writer 都写完，默认是 false
	BreakWhenErr bool `json:"break_when_err"`
}

// TeeWriter write data to multi writer
func TeeWriter(writers []io.Writer, opt *TeeWriterOption) io.Writer {
	if opt == nil {
		opt = new(TeeWriterOption)
	}
	return cio.NewTeeWriter(writers, &cio.TeeWriterOption{
		BreakWhenErr: opt.BreakWhenErr,
	})
}
