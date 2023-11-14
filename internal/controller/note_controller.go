package controller

import (
	"PersonalNotes/internal/entity"
	"PersonalNotes/internal/errors"
	"PersonalNotes/internal/usecases"
	"encoding/json"

	"net/http"
)

type NoteController interface {
	AddNotes(resp http.ResponseWriter, req *http.Request)
	GetNotes(resp http.ResponseWriter, req *http.Request)
}

type controller struct {}

var (
	usecase usecases.NoteUseCases
)

func NewNoteController(uc usecases.NoteUseCases) NoteController {
	usecase = uc
	return &controller{}
}

func (*controller) AddNotes(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var note entity.Note
	err := json.NewDecoder(req.Body).Decode(&note)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{
			Message: "Error unmarshalling data",
		})
		return
	}
	err = usecase.Validate(&note)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{
			Message: err.Error(),
		})
		return
	}
	result, err := usecase.Create(&note)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{
			Message: "Error creating the note",
		})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

func (*controller) GetNotes(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	notes, err := usecase.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{
			Message: "Error getting the notes",
		})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(notes)
}
