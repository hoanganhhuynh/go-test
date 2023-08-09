package models

import(
	"time"
)

type People struct {
	BaseModel
	FirstName string
	LastName string
	Country string
	Dob time.Time
}
