package main

import "fmt"

func Next(q *Queue) {

	for {
		msg := <-q.Msgs

		fmt.Printf("rec %s \n", &msg)
	}

}
