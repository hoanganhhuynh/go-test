package handlers

import(
	"gorm.io/gorm"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"time"

	models "example/models"
	repositories "example/repositories"
)

type UserHandler struct {
	PeopleRepository repositories.PeopleRepository
}

func NewUserHandler(db *gorm.DB) UserHandler {
	var peopleRepository repositories.PeopleRepository = repositories.PeopleRepository {
		BaseRepository : repositories.BaseRepository[models.People] { Db: db },
	}
	userHandler := UserHandler{
		PeopleRepository: peopleRepository,
	}
	return userHandler
}

func (userHandler UserHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	
	ctx := request.Context()

	t, err := userHandler.PeopleRepository.BaseRepository.GetAll(ctx)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else {
		err := json.NewEncoder(writer).Encode(t)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusFound)
		}
	}
}

func (userHandler UserHandler) GetById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	vars := mux.Vars(request)

	ctx := request.Context()

	id, parseIntError := strconv.ParseInt(vars["id"], 10, 64)
	if(parseIntError != nil) {
		http.Error(writer, parseIntError.Error(), http.StatusBadRequest)
	}
	
	t, err := userHandler.PeopleRepository.GetById(ctx, id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else {
		err := json.NewEncoder(writer).Encode(t)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusFound)
		}
	}
}

func (userHandler UserHandler) Create (writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	var model models.People
	err := json.NewDecoder(request.Body).Decode(&model)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := request.Context()
	model.CreateAt = time.Now()
	model.UpdateAt = time.Now()
	t, err := userHandler.PeopleRepository.BaseRepository.Create(ctx, model)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else {
		err := json.NewEncoder(writer).Encode(t)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusFound)
		}
	}
}

func (userHandler UserHandler) Update (writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	var model models.People
	err := json.NewDecoder(request.Body).Decode(&model)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := request.Context()
	model.UpdateAt = time.Now()
	t, err := userHandler.PeopleRepository.BaseRepository.Update(ctx, model)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else {
		err := json.NewEncoder(writer).Encode(t)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusFound)
		}
	}
}

