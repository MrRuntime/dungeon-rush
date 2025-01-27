package game

import "container/list"

type Queue[T any] struct {
	list *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: list.New(),
	}
}

func (q *Queue[T]) Enqueue(v T) {
	q.list.PushBack(v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.list.Len() == 0 {
		var zero T
		return zero, false
	}
	head := q.list.Front()
	q.list.Remove(head)
	return head.Value.(T), true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.list.Len() == 0 {
		var zero T
		return zero, false
	}
	return q.list.Front().Value.(T), true
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}
