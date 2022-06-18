package main

import (
	"danielr1996/bashdoard/providers"
	"danielr1996/bashdoard/sse"
	"os"
	"os/signal"
)

func main() {
	// Start Listening
	s := sse.New()
	stopCh := make(chan struct{})
	go s.Serve(stopCh)

	// Register Providers
	var ps []providers.Provider
	ps = append(ps, new(providers.DockerProvider))
	for _, provider := range ps {
		go provider.Push(s)
	}

	// Handle Stop signals
	sigCh := make(chan os.Signal, 0)
	signal.Notify(sigCh, os.Kill, os.Interrupt)
	<-sigCh
	close(stopCh)
}
