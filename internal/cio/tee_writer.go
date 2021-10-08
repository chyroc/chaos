package cio

import (
	"io"
)

type TeeWriter struct {
	writers []io.Writer
	opt     *TeeWriterOption
}

type TeeWriterOption struct {
	// When an error occurs when writing data in multiple writers, exit immediately,
	// instead of writing all the writers,
	// the default is false
	// 当在多个 writer 中写数据发生错误的时候，立刻退出，而不是将所有 writer 都写完，默认是 false
	BreakWhenErr bool `json:"break_when_err"`
}

func NewTeeWriter(writers []io.Writer, opt *TeeWriterOption) *TeeWriter {
	return &TeeWriter{
		writers: writers,
		opt:     opt,
	}
}

func (r *TeeWriter) Write(p []byte) (n int, err error) {
	var finalErr error
	var writeCount int
	for _, v := range r.writers {
		n, err := v.Write(p)
		if err != nil {
			if r.opt.BreakWhenErr {
				return n, err
			} else {
				finalErr = err
				if n > writeCount {
					writeCount = n
				}
			}
		} else {
			if n > writeCount {
				writeCount = n
			}
		}
	}
	return writeCount, finalErr
}
