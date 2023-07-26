package models

import(
	"time"
)

type BaseModel struct {
	Id int64
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
}