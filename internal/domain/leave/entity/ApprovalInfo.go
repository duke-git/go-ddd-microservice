package entity

import "go-ddd-microservice/internal/domain/leave/entity/vo"

type ApprovalInfo struct {
	ApprovalInfoId string
	Approver       vo.Approver
	ApprovalType   vo.ApprovalType
	Message        string
	Time           int64
}
