package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type ExitStatus interface {
	ExitStatus() int
}

func main() {
	var real, user, sys time.Duration
	var status int

	posix_fmt := flag.Bool("p", false, "use POSIX formatting")
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		cmd := exec.Command(args[0], args[1:len(args)]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		start := time.Now()
		err := cmd.Run()

		if process := cmd.ProcessState; process != nil {
			real = time.Since(start)
			user = process.UserTime()
			sys = process.SystemTime()

			if es, ok := process.Sys().(ExitStatus); ok {
				status = es.ExitStatus()
			} else if process.Success() {
				status = 0
			} else {
				status = 1
			}
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			if err, ok := err.(*exec.Error); ok && err.Err == exec.ErrNotFound {
				status = 127
			} else {
				status = 126
			}
		}
	}
	if *posix_fmt {
		fmt.Fprintf(os.Stderr, "real %f\nuser %f\nsys %f\n", real.Seconds(), user.Seconds(), sys.Seconds())
	} else {
		fmt.Fprintf(os.Stderr, "\nreal\t%v\nuser\t%v\nsys\t%v\n", real, user, sys)
	}
	os.Exit(status)
}
