package chansort

import (
	"constraints"
	"time"

	"github.com/jamesrom/priorityqueue"
)

func SortOrderable[T constraints.Ordered](in chan T) chan T {
	q := priorityqueue.NewWithOrderable[T]()
	out := make(chan T)
	go func() {
		for {
			el := <-in
			q.Push(el)
			time.AfterFunc(time.Second*15, func() {
				out <- q.Pop()
			})
		}
	}()
	return out
}
