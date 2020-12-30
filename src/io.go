package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func iotest() {
	fmt.Println("--------------------io--------------------")
	f, err := ioutil.TempFile("", "sample_*.txt")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	h := sha256.New()
	w := io.MultiWriter(f, h)
	_, err = io.WriteString(w, "hello")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	err = f.Close()
	defer os.RemoveAll(f.Name())
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Printf("%v\n", f.Name())
}
