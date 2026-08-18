package queue

import (
	"errors"
	"fmt"
	"iter"
	"strings"
)

var (
	ErrEmptyQueue = errors.New("queue is empty")
)

// Queue Implementation
type Queue[T any] struct {
	items []T
}

func New[T any](items ...T) *Queue[T] {
	q := &Queue[T]{
		items: make([]T, 0, len(items)),
	}
	q.Enqueue(items...)
	return q
}

func (q *Queue[T]) Enqueue(items ...T) {
	q.items = append(q.items, items...)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, ErrEmptyQueue
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[T]) Front() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, ErrEmptyQueue
	}
	return q.items[0], nil
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Clear() {
	q.items = nil
}

func (q *Queue[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range q.items {
			if !yield(item) {
				return
			}
		}
	}
}

func (q *Queue[T]) All2() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range q.items {
			if !yield(i, item) {
				return
			}
		}
	}
}

func (q *Queue[T]) Drain() iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			item, ok := q.Dequeue()
			if ok != nil || !yield(item) {
				return
			}
		}
	}
}

func (q *Queue[T]) String() string {
	var sb strings.Builder
	sb.WriteString("Queue[")
	for i, item := range q.items {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", item))
	}
	sb.WriteString("]")
	return sb.String()
}
