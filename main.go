package main

import (
	"github.com/agbankar/go-worker-pool/dispatcher"
	"time"
)

func main() {
	dispatcher.New(5).Start()
	time.Sleep(10 * 60 * time.Second)

}
