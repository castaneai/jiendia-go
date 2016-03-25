package spf

import (
	"io"
	"os"
)

type Reader struct {
	reader io.ReaderAt
	File   []*file
}

type ReadCloser struct {
	f *os.File
	Reader
}

func OpenReader(name string) (*ReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	r := new(ReadCloser)
	if err := r.init(f, fi.Size()); err != nil {
		f.Close()
		return nil, err
	}
	r.f = f
	return r, nil
}

func NewReader(srcReader io.ReaderAt, size int64) (*Reader, error) {
	r := new(Reader)
	if err := r.init(srcReader, size); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Reader) init(sourceReader io.ReaderAt, size int64) error {
	sr := io.NewSectionReader(sourceReader, 0, size)
	sr.Seek(-140, os.SEEK_END)

}
