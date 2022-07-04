package vo

type Status int

const (
	APPROVIN Status = iota + 1
	APPROVED
	REJECTED
)
