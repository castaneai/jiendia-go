package spf

import (
	"io"
	"bufio"
)

type Writer struct {
	dir []*header
}

type header struct {
	*FileHeader
}

func NewWriter(w io.Writer) *Writer {

}
