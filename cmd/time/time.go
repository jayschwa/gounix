package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

import "gounix.org/cmd"

type ExitStatus interface {
	ExitStatus() int
}

func main() {
	cmd.Init("time")
	defer cmd.Exit()

	var real, user, sys time.Duration

	posix_fmt := flag.Bool("p", false, "use POSIX formatting")
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		racer := exec.Command(args[0], args[1:len(args)]...)
		racer.Stdin = os.Stdin
		racer.Stdout = os.Stdout
		racer.Stderr = os.Stderr

		signals := make(chan os.Signal)
		signal.Notify(signals)

		start := time.Now()
		err := racer.Run()

		signal.Stop(signals)

		if process := racer.ProcessState; process != nil {
			real = time.Since(start)
			user = process.UserTime()
			sys = process.SystemTime()

			if es, ok := process.Sys().(ExitStatus); ok {
				cmd.ExitStatus = es.ExitStatus()
			} else if process.Success() {
				cmd.ExitStatus = 0
			} else {
				cmd.ExitStatus = 1
			}
		} else if err != nil {
			cmd.Errorln(err)
			if err, ok := err.(*exec.Error); ok && err.Err == exec.ErrNotFound {
				cmd.ExitStatus = 127
			} else {
				cmd.ExitStatus = 126
			}
		}
	}
	if *posix_fmt {
		fmt.Fprintf(os.Stderr, "real %f\nuser %f\nsys %f\n", real.Seconds(), user.Seconds(), sys.Seconds())
	} else {
		fmt.Fprintf(os.Stderr, "\nreal\t%v\nuser\t%v\nsys\t%v\n", real, user, sys)
	}
}
