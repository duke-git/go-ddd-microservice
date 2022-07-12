package event

import (
	"go-ddd-microservice/internal/domain/leave/entity"
	"go-ddd-microservice/internal/infrastructure/event"

	"github.com/duke-git/lancet/convertor"
	"github.com/duke-git/lancet/datetime"
)

type LeaveEvent struct {
	event.DomainEvent
	LeaveEventType LeaveEventType
}

func NewLeaveEvent(eventType LeaveEventType, leave entity.Leave) *LeaveEvent {
	data, err := convertor.ToJson(leave)
	if err != nil {
		return nil
	}
	id, err := convertor.ToJson(leave)
	if err != nil {
		return nil
	}

	return &LeaveEvent{
		event.DomainEvent{
			Id:        id,
			TimeStamp: datetime.GetZeroHourTimestamp(),
			Data:      data,
		},
		eventType,
	}
}

func (le *LeaveEvent) Send() error {
	//send event to event bus or mq
	return nil
}
