package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func Publish(q *Queue) {

	for i := 0; i < 100000; i++ {
		if i%7 == 0 {

			msg := Message{
				ID:         uuid.New().ID(),
				Data:       fmt.Sprintf("semd data %d", i),
				CreateTime: time.Now(),
			}

			q.Msgs <- msg
		}
	}

}
