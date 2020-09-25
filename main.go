package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/raviks789/check-massive/lib"
	log "github.com/sirupsen/logrus"
)

func main() {
	{
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		go handleSignal(ch)
	}
	go checkmass.Do()

	fmt.Printf("go")
}

func handleSignal(ch <-chan os.Signal) {
	if sig, ok := <-ch; ok {
		log.WithFields(log.Fields{"signal": sig}).Info("Shutting down")
		os.Exit(0)
	}
}
