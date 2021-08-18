package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handle(signal os.Signal)  {
	fmt.Println("Received:",signal)
}

func main()  {
	sigs := make(chan os.Signal,1)
	signal.Notify(sigs)
	for{
		sig := <- sigs
		switch sig {
		case os.Interrupt:
			handle(sig)
		case syscall.SIGTERM:
			handle(sig)
			os.Exit(0)
		case syscall.SIGBUS:
			fmt.Println("")
		default:
			fmt.Println()
		}
	}
}
