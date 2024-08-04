package dto

import "time"

type PersonDTO struct {
	PersonID         string
	PersonName       string
	RoleID           int
	PersonType       int
	CreateTime       time.Time
	LastModifiedTime time.Time
	Status           int
}
