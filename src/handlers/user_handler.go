package handlers

import(
	"gorm.io/gorm"

	models "example/models"
	repositories "example/repositories"
)

type UserHandler struct {
	BaseHandler[models.People]
}

func NewUserHandler(db *gorm.DB) UserHandler {
	var peopleRepository repositories.IPeopleRepository = repositories.PeopleRepository {
		BaseRepository : repositories.BaseRepository[models.People] { Db: db },
	}
	userHandler := UserHandler{
		BaseHandler: BaseHandler[models.People]{
			baseRepository: peopleRepository,
		},
	}
	return userHandler
}

