package main

import (
	"flag"
	"io"
	"os"
)

import "gounix.org/cmd"

func main() {
	cmd.Init("cat")
	defer cmd.Exit()

	flag.Bool("u", true, "unbuffered reads")
	flag.Parse()
	files := flag.Args()
	if len(files) <= 0 {
		files = []string{"-"}
	}
	for _, file := range files {
		err := cat(file)
		if err != nil {
			cmd.Errorln(err)
		}
	}
}

func cat(filename string) (err error) {
	var src *os.File
	if filename == "-" {
		src = os.Stdin
	} else {
		src, err = os.Open(filename)
		if err != nil {
			return err
		}
		defer src.Close()
	}
	_, err = io.Copy(os.Stdout, src)
	return err
}
