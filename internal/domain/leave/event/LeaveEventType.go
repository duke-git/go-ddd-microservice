package event

type LeaveEventType int

const (
	CREATE_EVENT LeaveEventType = iota + 1
	AGREE_EVENT
	REJECT_EVENT
	APPROVED_EVENT
)
