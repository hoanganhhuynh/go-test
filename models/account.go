package models

type Account struct {
	BaseModel
	EmployeeId int64
	Amount int64
}