package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

	repositories "example/repositories"
)

type BaseHandler[T any] struct {
	baseRepository repositories.IBaseRepository[T]
}

func (baseHandler BaseHandler[T]) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	
	t, err := baseHandler.baseRepository.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else {
		err := json.NewEncoder(writer).Encode(t)
		fmt.Println("Finished getting people")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusFound)
		}
	}
}

func (baseHandler BaseHandler[T]) GetById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	vars := mux.Vars(request)

	ctx := request.Context()

	id, parseIntError := strconv.ParseInt(vars["id"], 10, 64)
	if(parseIntError != nil) {
		http.Error(writer, parseIntError.Error(), http.StatusBadRequest)
	}
	
	t, err := baseHandler.baseRepository.GetById(ctx, id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else {
		err := json.NewEncoder(writer).Encode(t)
		fmt.Println("Finished getting people")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusFound)
		}
	}
}

func (baseHandler BaseHandler[T]) Create (writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	var model T
	err := json.NewDecoder(request.Body).Decode(&model)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := request.Context()

	t, err := baseHandler.baseRepository.Create(ctx, model)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else {
		err := json.NewEncoder(writer).Encode(t)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusFound)
		}
	}
}


