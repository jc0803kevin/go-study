package main

import (
	"time"
)

type Message struct {
	ID uint32

	Data string

	CreateTime time.Time
}

type Queue struct {
	Msgs chan Message
}

func RunQueue() {
	queue := &Queue{Msgs: make(chan Message)}

	go Publish(queue)
	go Next(queue)

	time.Sleep(time.Second * 10)

}
