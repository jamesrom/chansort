package chansort_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jamesrom/chansort"
)

func TestSimple(t *testing.T) {
	messages := make(chan int)
	go func() { messages <- 1 }()
	go func() { messages <- 3 }()
	go func() { messages <- 4 }()
	go func() { messages <- 3 }()
	go func() { messages <- 6 }()
	time.AfterFunc(time.Second*2, func() {
		messages <- 10
	})

	sortedMessages := chansort.SortOrderable(messages)

	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
}
