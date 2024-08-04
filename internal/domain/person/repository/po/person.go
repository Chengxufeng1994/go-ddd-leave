package po

import "time"

// PersonPO persistent object (GORM)
type Person struct {
	PersonID         string `gorm:"primaryKey"`
	PersonName       string
	DepartmentID     string
	PersonType       int
	LeaderID         string
	RoleLevel        int
	CreateTime       time.Time
	LastModifiedTime time.Time
	Status           int
	Relationship     Relationship `gorm:"foreignKey:PersonID;references:PersonID"`
}

func (p Person) TableName() string {
	return "person"
}
