package messaging

import "errors"

type PubSub[T any] struct {
	queue chan T
}

func NewPubSub[T any](bufferSize int) *PubSub[T] {
	return &PubSub[T]{
		queue: make(chan T, bufferSize),
	}
}

func (p *PubSub[T]) Publish(message T) error {
	select {
	case p.queue <- message:
		return nil
	default:
		return errors.New("queue is full")
	}
}

func (p *PubSub[T]) Subscribe() <-chan T {
	return p.queue
}

func (p *PubSub[T]) Close() {
	close(p.queue)
}
