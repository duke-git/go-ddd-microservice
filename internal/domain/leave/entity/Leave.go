package entity

import (
	"go-ddd-microservice/internal/domain/leave/entity/vo"
)

type Leave struct {
	Id                   string
	Applicant            vo.Applicant
	Approver             vo.Approver
	Type                 vo.LeaveType
	Status               vo.Status
	StartTime            string
	EndTime              string
	Duration             int64
	LeaderMaxLevel       int
	CurrentApprovalInfos []*ApprovalInfo
}
