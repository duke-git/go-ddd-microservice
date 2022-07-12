package po

type LeavePO struct {
	Id            string `gorm:"type:varchar(20);primary_key" json:"id"`
	ApplicantId   string `gorm:"type:varchar(255)" json:"applicantId"`
	ApplicantName string `gorm:"type:varchar(255)" json:"applicantName"`
	ApplicantType int    `gorm:"type:int(10)" json:"ApplicantType"`
	ApproverId    string `gorm:"type:varchar(255)" json:"approverId"`
	ApproverName  string `gorm:"type:varchar(255)" json:"approverName"`
	LeaveType     int    `gorm:"type:int(10)" json:"leaveType"`
	Status        int    `gorm:"type:int(10)" json:"status"`
	StartTime     string `gorm:"type:varchar(255)" json:"startTime"`
	EndTime       string `gorm:"type:varchar(255)" json:"endTime"`
	Duration      int64  `gorm:"type:int(50)" json:"duration"`
}

func (*LeavePO) TableName() string {
	return "leave"
}
