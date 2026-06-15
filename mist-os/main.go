package main

import (
	"fmt"
	"mist-os/brain"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("i am mist")

	b, err := brain.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "startup error: ", err)
		os.Exit(1)
	}
	defer b.Close()

	if err := b.Body.Arm(); err != nil {
		fmt.Fprintln(os.Stderr, "arm error: ", err)
		os.Exit(1)
	}

	stop := make(chan struct{})
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		close(stop)
	}()
	b.Run(stop)
}
