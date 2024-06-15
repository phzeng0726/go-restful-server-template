package domain

import "time"

type Automation struct {
	Id             int       `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	TrackingId     int       `gorm:"column:tracking_id" json:"trackingId"`
	DocumentId     string    `gorm:"column:document_id" json:"documentId"`
	NotesUrl       string    `gorm:"column:notes_url" json:"notesUrl"`
	Project        string    `gorm:"column:project" json:"project"`
	Requester      string    `gorm:"column:requester" json:"requester"`
	DepartmentCode string    `gorm:"column:department_code" json:"departmentCode"`
	Developer      string    `gorm:"column:developer" json:"developer"`
	AutomationType string    `gorm:"column:automation_type" json:"automationType"`
	Link           string    `gorm:"column:link" json:"link"`
	AssignDate     time.Time `gorm:"column:assign_date" json:"assignDate"`
	PlanDate       time.Time `gorm:"column:plan_date" json:"planDate"`
	ExpectDate     time.Time `gorm:"column:expect_date" json:"expectDate"`
	CloseDate      time.Time `gorm:"column:close_date" json:"closeDate"`
	DutStatus      string    `gorm:"column:dut_status" json:"dutStatus"`
	Status         string    `gorm:"column:status" json:"status"`
	Spec           string    `gorm:"column:spec" json:"spec"`
	Description    string    `gorm:"column:description" json:"description"`
	Notify         string    `gorm:"column:notify" json:"notify"`
	Images         string    `gorm:"column:images" json:"images"`
}

func (Automation) TableName() string {
	return "automation"
}
