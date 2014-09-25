package spf_test

import (
	"archive/spf"
	"bytes"
	"log"
	"fmt"
)

func ExampleWriter() {
	buf := new (bytes.Buffer)

	// create a new SPF archive.
	w := spf.NewWriter(buf)

	var files = []struct {Name, Body string} {
		{"readme.txt", "This is readme."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleReader() {
	// Open a new SPF archive for reading.
	r, err := spf.OpenReader("testdata/example.spf")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
	}
}
