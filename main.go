package main

import (
	"os"
	"os/signal"
	"syscall"
	"checkmass"
	log "github.com/sirupsen/logrus"
)

func main() {
	{
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		go handleSignal(ch)
	}
	go checkmass.Do()
}

func handleSignal(ch <-chan os.Signal) {
	if sig, ok := <-ch; ok {
		log.WithFields(log.Fields{"signal": sig}).Info("Shutting down")
		os.Exit(0)
	}
}
