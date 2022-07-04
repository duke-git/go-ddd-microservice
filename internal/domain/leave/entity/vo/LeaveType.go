package vo

type LeaveType int

const (
	Internal LeaveType = iota + 1
	External
	Official
)
