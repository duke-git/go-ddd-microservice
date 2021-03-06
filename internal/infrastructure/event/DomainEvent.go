package event

type EventSender interface {
	Send(event DomainEvent) error
}

type DomainEvent struct {
	Id        string
	TimeStamp int64
	Source    string
	Data      string
}
