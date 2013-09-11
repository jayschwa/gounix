package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func cat(name string) (err error) {
	src := os.Stdin
	if name != "-" {
		src, err = os.Open(name)
		if err != nil {
			return err
		}
		defer src.Close()
	}
	_, err = io.Copy(os.Stdout, src)
	return err
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
	for _, file := range files {
		err := cat(file)
		if err != nil {
			log.Println(err)
		}
	}
}