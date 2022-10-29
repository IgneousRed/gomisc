package gomisc

type Queue[T any] struct {
	slice      []T
	start, end int
}
type QueueEmpty string

func (q QueueEmpty) Error() string {
	return string(q)
}
func QueueNew[T any](initial []T) Queue[T] {
	if len(initial) < 4 {
		initial = SliceNewCopy(initial, 4)
	}
	return Queue[T]{slice: initial}
}
func (q Queue[T]) wrap(value int) int {
	return Wrap(value, len(q.slice))
}
func (q Queue[T]) Len() int {
	return q.wrap(q.end - q.start)
}
func (q Queue[T]) Slice() []T {
	if q.end < q.start {
		result := make([]T, q.Len())
		len := copy(result, q.slice[q.start:])
		copy(result[len:], q.slice[:q.end])
		return result
	}
	return q.slice[q.start:q.end]
}
func (q *Queue[T]) Push(value T) error {
	if q.wrap(q.start-q.end) <= 1 {
		q.slice = SliceExpand(q.slice)
		q.end = q.Len()
		q.start = 0
	}
	q.slice[q.end] = value
	q.end = q.wrap(q.end + 1)
	return nil
}
func (q Queue[T]) First() (T, error) {
	if q.start == q.end {
		var result T
		return result, QueueEmpty("Queue empty")
	}
	return q.slice[q.start], nil
}
func (q Queue[T]) Last() (T, error) {
	if q.start == q.end {
		var result T
		return result, QueueEmpty("Queue empty")
	}
	return q.slice[q.wrap(q.end-1)], nil
}
func (q *Queue[T]) Pop() (T, error) {
	if q.start == q.end {
		var result T
		return result, QueueEmpty("Queue empty")
	}
	result := q.slice[q.start]
	q.start = q.wrap(q.start + 1)
	return result, nil
}
