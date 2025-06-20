package messaging

type PubSub[T any] struct {
	queue chan T
}

func NewPubSub[T any](bufferSize int) *PubSub[T] {
	return &PubSub[T]{
		queue: make(chan T, bufferSize),
	}
}

func (p *PubSub[T]) Publish(message T) {
	p.queue <- message
}

func (p *PubSub[T]) Subscribe() <-chan T {
	return p.queue
}

func (p *PubSub[T]) Close() {
	close(p.queue)
}
