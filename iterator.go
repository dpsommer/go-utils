package goutils

import "iter"

type Next[T any] func() (T, bool)

type Current[T any] func() T

type Iterator[T any] struct {
	Seq iter.Seq[T]
	Current[T]
	Next[T]
}

func NewIterator[T any](tt []T) Iterator[T] {
	var i int

	// return an iterator and a closure that increments the iterator so that we
	// can resume iteration from the same key if a child is running
	return Iterator[T]{
		Seq: func(yield func(T) bool) {
			for j := 0; j < len(tt); j += 1 {
				if !yield(tt[j]) {
					return
				}
			}
		},
		Current: func() T { return tt[i] },
		Next: func() (T, bool) {
			if i+1 >= len(tt) {
				return *new(T), false
			}
			i += 1
			return tt[i], true
		},
	}
}
