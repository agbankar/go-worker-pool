package main

import (
	"github.com/PrakharSrivastav/workers/b_concurrent/dispatcher"
	"time"
)

func main() {
	dispatcher.New(5).Start()
	time.Sleep(10 * 60 * time.Second)

}
