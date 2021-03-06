package main

import (
	"sync"
	"time"
	"log"
	. "github.com/baofengqqwwff/BtcQuant/engine"
)

func main() {
	engine := NewEngine()
	engine.RegisterProcessor(&Processor{ProcessorName:"timerProcessor",EventName: "Timer", EventHandler: timerProcessor})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(1000 * time.Second)
		defer wg.Done()
	}()
	wg.Wait()
}

func timerProcessor(event *Event)(*Event,error) {
	log.Println(event.Data)
	return nil,nil
}
