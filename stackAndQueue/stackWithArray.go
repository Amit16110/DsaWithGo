package stackandqueue

type QueueWithArray struct {
	items []int
}

func (s *QueueWithArray) Enqueue(item int) {
	s.items = append(s.items, item)
}

func (s *QueueWithArray) Dequeue() int {
	if s.isEmpty() {
		return -1
	}

	item := s.items[0]

	s.items = s.items[1:]
	return item
}

func (s *QueueWithArray) isEmpty() bool {
	return len(s.items) == 0

}
