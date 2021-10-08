package chaos

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mustErrWriter int

func (r mustErrWriter) Write(p []byte) (n int, err error) {
	return int(r), fmt.Errorf("fake-err")
}

func Test_TeeWriter(t *testing.T) {
	as := assert.New(t)

	type args struct {
		writers []io.Writer
		opt     *TeeWriterOption
	}
	type in struct {
		write      string
		wantN      int
		errContain string
	}
	tests := []struct {
		name  string
		args  args
		write in
	}{
		{name: "1 - opt-nil", args: args{writers: []io.Writer{new(bytes.Buffer)}, opt: nil}, write: in{write: "x", wantN: 1}},
		{name: "1 - opt-break-false", args: args{writers: []io.Writer{new(bytes.Buffer)}, opt: &TeeWriterOption{}}, write: in{write: "x", wantN: 1}},
		{name: "1 - opt-break-true", args: args{writers: []io.Writer{new(bytes.Buffer)}, opt: &TeeWriterOption{BreakWhenErr: true}}, write: in{write: "x", wantN: 1}},

		{name: "2 - opt-nil", args: args{writers: []io.Writer{new(bytes.Buffer), mustErrWriter(1)}, opt: nil}, write: in{write: "xx", wantN: 2, errContain: "fake-err"}},
		{name: "2 - opt-break-false", args: args{writers: []io.Writer{new(bytes.Buffer), mustErrWriter(1)}, opt: &TeeWriterOption{}}, write: in{write: "xx", wantN: 2, errContain: "fake-err"}},
		{name: "2 - opt-break-true", args: args{writers: []io.Writer{new(bytes.Buffer), mustErrWriter(1)}, opt: &TeeWriterOption{BreakWhenErr: true}}, write: in{write: "xx", wantN: 2, errContain: "fake-err"}},
		{name: "2 - opt-break-true-first", args: args{writers: []io.Writer{mustErrWriter(1), new(bytes.Buffer)}, opt: &TeeWriterOption{BreakWhenErr: true}}, write: in{write: "xx", wantN: 0, errContain: "fake-err"}},

		{name: "2 - opt-break-false", args: args{writers: []io.Writer{mustErrWriter(1), new(bytes.Buffer)}, opt: &TeeWriterOption{}}, write: in{write: "xx", wantN: 2, errContain: "fake-err"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := fmt.Sprintf("%s - %s", tt.name, tt.write.write)
			w := TeeWriter(tt.args.writers, tt.args.opt)
			n, err := w.Write([]byte(tt.write.write))
			if tt.write.errContain != "" {
				as.NotNil(err, msg)
				as.Contains(err.Error(), tt.write.errContain)
			} else {
				as.Nil(err, msg)
				as.Equal(tt.write.wantN, n)
			}
		})
	}
}
