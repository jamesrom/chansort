package chansort_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jamesrom/chansort"
)

func TestAscending(t *testing.T) {
	messages := make(chan int)
	go func() { messages <- 1 }()
	go func() { messages <- 2 }()
	go func() { messages <- 3 }()
	go func() { messages <- 4 }()
	go func() { messages <- 5 }()
	time.AfterFunc(time.Second*2, func() {
		messages <- -10
	})

	sortedMessages := chansort.SortOrderable(messages, 15*time.Second)

	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
}

func TestDescending(t *testing.T) {
	messages := make(chan int)
	go func() { messages <- 1 }()
	go func() { messages <- 2 }()
	go func() { messages <- 3 }()
	go func() { messages <- 4 }()
	go func() { messages <- 5 }()
	time.AfterFunc(time.Second*2, func() {
		messages <- 10
	})

	sortedMessages := chansort.SortOrderable(messages, 15*time.Second)

	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
	fmt.Println(<-sortedMessages)
}
