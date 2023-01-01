package gomisc

// First In, First Out.
type Queue[T any] struct {
	slice      []T
	start, end int
}

// Error.
type QueueEmpty string

func (q QueueEmpty) Error() string {
	return string(q)
}

// Wrap on len.
func (q Queue[T]) wrap(value int) int {
	return Wrap(value, len(q.slice))
}

// Elements in Queue.
func (q Queue[T]) Len() int {
	return q.wrap(q.end - q.start)
}

// Slice of all elements.
func (q Queue[T]) Slice() []T {
	if q.end < q.start {
		result := make([]T, q.Len())
		copy(result[copy(result, q.slice[q.start:]):], q.slice[:q.end])
		return result
	}
	return q.slice[q.start:q.end]
}

// Put an element at the end.
func (q *Queue[T]) Push(value T) error {
	if q.wrap(q.start-q.end) <= 1 {
		q.start, q.end = 0, q.Len()
		q.slice = SliceExpand(q.Slice(), 4)
	}
	q.slice[q.end], q.end = value, q.wrap(q.end+1)
	return nil
}

// The first element.
func (q Queue[T]) First() (T, error) {
	if q.start == q.end {
		var result T
		return result, QueueEmpty("Queue empty")
	}
	return q.slice[q.start], nil
}

// The last element.
func (q Queue[T]) Last() (T, error) {
	if q.start == q.end {
		var result T
		return result, QueueEmpty("Queue empty")
	}
	return q.slice[q.wrap(q.end-1)], nil
}

// Take the first element.
func (q *Queue[T]) Pop() (T, error) {
	if q.start == q.end {
		var result T
		return result, QueueEmpty("Queue empty")
	} else if q.Len() < len(q.slice)/3 {
		q.start, q.end = 0, q.Len()
		q.slice = SliceShrink(q.Slice(), 4)
	}
	result := q.slice[q.start] // TODO: Merge lines
	q.start = q.wrap(q.start + 1)
	return result, nil
}
