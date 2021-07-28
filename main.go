package main

import (
	"github.com/agbankar/go-worker-pool/dispatcher"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	d := dispatcher.NewDispatcher(5)
	go func() {
		wg.Add(1)
		d.Start()
		wg.Done()
	}()
	<-sigs
	d.Stop()
	wg.Wait()

}
