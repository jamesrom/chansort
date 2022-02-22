package chansort

import (
	"constraints"
	"time"

	"github.com/jamesrom/priorityqueue"
)

// Less must describe a transitive ordering:
//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
type Less[T any] func(T, T) bool

// SortOrderable sorts channel messages in ascending order. Messages received
// inside the sliding-window buffer defined by _window_ are sent to the
// output channel in ascending order. That is to say: a message received
// at time _Z_  from the output channel is guaranteed to be the smallest
// message since _Z − window_.
func SortOrderable[T constraints.Ordered](in <-chan T, window time.Duration) <-chan T {
	defaultComparator := func(a, b T) bool { return a < b }
	return SortWithComparator(in, window, defaultComparator)
}

// SortWithComparator sorts channel messages in the order defined by the given
// comparator function. Messages received inside the sliding-window buffer
// defined by _window_ are sent to the output channel in order.
// That is to say: a message received at time _Z_ from the output channel is
// guaranteed to be the largest message since _Z − window_.
func SortWithComparator[T any](in <-chan T, window time.Duration, fn Less[T]) <-chan T {
	q := priorityqueue.NewWithComparator(fn)
	out := make(chan T)
	go func() {
		for {
			el := <-in
			q.Push(el)
			time.AfterFunc(window, func() {
				out <- q.Pop()
			})
		}
	}()
	return out
}
