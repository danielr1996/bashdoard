package main

import (
	"danielr1996/bashboard/client"
	"os"
	"os/signal"
)

func main() {
	client := client.New()

	stopCh := make(chan struct{})
	go client.StartWatching(stopCh)
	go client.PushUpdates(stopCh)
	sigCh := make(chan os.Signal, 0)
	signal.Notify(sigCh, os.Kill, os.Interrupt)
	<-sigCh
	close(stopCh)
}
