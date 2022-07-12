package event

type EventPublisher struct {
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{}
}

func (ep *EventPublisher) Publish(sender EventSender, event DomainEvent) error {
	//send event to event bus or mq
	return sender.Send(event)
}
