package utils

type Queue[T any] struct {
	elems []T
}

func (q *Queue[T]) Add(item T) {
	q.elems = append(q.elems, item)
}

func (q *Queue[T]) Size() int {
	return len(q.elems)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Remove() *T {
	if q.IsEmpty() {
		return nil
	}

	removed := q.elems[0]
	q.elems = q.elems[1:]
	return &removed
}
