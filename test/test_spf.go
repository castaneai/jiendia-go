package spf_test

import (
	"fmt"
	"log"

	"../src/archive/spf"
)

func ExampleReader() {
	// Open a new SPF archive for reading.
	r, err := spf.OpenReader("testdata/CLAIRE.SPF")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for i, f := range r.File {
		fmt.Printf("[%d] %s:\n", f.Path)
	}
}
