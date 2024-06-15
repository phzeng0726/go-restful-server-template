package domain

type VacuumInfo struct {
	Id             uint  `gorm:"column:id;not null;primaryKey;"`
	LastVacuumTime int64 `gorm:"column:last_vacuum_time;"`
}

func (VacuumInfo) TableName() string {
	return "vacuum_info"
}
