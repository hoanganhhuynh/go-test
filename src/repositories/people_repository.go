package repositories

import(
	models "example/models"
)

type PeopleRepository struct {
	BaseRepository[models.People]
}

type IPeopleRepository interface {
	IBaseRepository[models.People]
}