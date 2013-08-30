package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func catFile(name string) (err error) {
	input := os.Stdin
	if name != "-" {
		input, err = os.Open(name)
		if err != nil {
			return
		}
		defer input.Close()
	}
	_, err = io.Copy(os.Stdout, input)
	return
}

func main() {
	log.SetPrefix("cat: ")
	log.SetFlags(0)
	flag.Bool("u", true, "unbuffered reads")
	flag.Parse()
	files := flag.Args()
	if len(files) <= 0 {
		files = []string{"-"}
	}
	for _, f := range files {
		err := catFile(f)
		if err != nil {
			log.Println(err)
		}
	}
}
