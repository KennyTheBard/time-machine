package main

import (
	"fmt"
	"time"

	tm "github.com/KennyTheBard/timeline/timemachine"
)

func main() {

	tm.Machine.Init()
	tm.Machine.Start()
	defer tm.Machine.Stop()

	duration, _ := time.ParseDuration("1s")
	tm.Machine.AddTimePoint(tm.TimePoint{
		Point:    time.Now().Add(duration),
		WaitTime: duration,
		IsRepetable: func() bool {
			return true
		},
		Action: func() {
			fmt.Println("Hello world!")
		},
	})

	// blocking line
	block := make(chan int)
	<-block
}
