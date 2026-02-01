package event

import "log"

type NoopPublisher struct{}

func NewNoopPublisher() *NoopPublisher {
    return &NoopPublisher{}
}
func (p *NoopPublisher) Publish(event any) error {
    log.Printf("NoopPublisher: event published: %+v\n", event)
    return nil
}