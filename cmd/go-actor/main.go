package main

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

// This example will demonstrate how to create actors for producer-consumer use case.
// Producer will create incremented number on every 1 second interval and
// consumer will print whatever number it receives.
func main() {
	mailbox := actor.NewMailbox[int]()

	// Produce and consume workers are created with same mailbox
	// so that produce worker can send messages directly to consume worker
	pw := &produceWorker{outC: mailbox.SendC()}
	cw1 := &consumeWorker{inC: mailbox.ReceiveC(), id: 1}

	// Note: Example creates two consumers for the sake of demonstration
	// since having one or more consumers will produce the same result.
	// Message on stdout will be written by first consumer that reads from mailbox.
	cw2 := &consumeWorker{inC: mailbox.ReceiveC(), id: 2}

	// Create actors using these workers and combine them to singe actor
	a := actor.Combine(
		mailbox,
		actor.New(pw),
		actor.New(cw1),
		actor.New(cw2),
	)

	// Finally all actors are started and stopped at once
	a.Start()
	defer a.Stop()

	select {}
}

// produceWorker will produce incremented number on 1 second interval
type produceWorker struct {
	outC chan<- int
	num  int
}

func (w *produceWorker) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-time.After(time.Second):
		w.num++
		w.outC <- w.num

		return actor.WorkerContinue

	case <-c.Done():
		return actor.WorkerEnd
	}
}

// consumeWorker will consume numbers received on inC channel
type consumeWorker struct {
	inC <-chan int
	id  int
}

func (w *consumeWorker) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case num := <-w.inC:
		fmt.Printf("consumed %d \t(worker %d)\n", num, w.id)

		return actor.WorkerContinue

	case <-c.Done():
		return actor.WorkerEnd
	}
}
