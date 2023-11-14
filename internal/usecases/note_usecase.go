package usecases

import (
	"PersonalNotes/internal/entity"
	"PersonalNotes/internal/repository"
	"errors"
)

type NoteUseCases interface {
	Validate(note *entity.Note) error
	Create(note *entity.Note) (*entity.Note, error)
	FindAll() ([]entity.Note, error)
}

type usecase struct {}

var (
	repo repository.NoteRepository
)

func NewNoteUseCases(rp repository.NoteRepository) NoteUseCases {
	repo = rp
    return &usecase{}
}

func (*usecase) Validate(note *entity.Note) error {
	if note == nil {
		err := errors.New("The note is empty")
		return err
	}
	if note.Title == "" {
		err := errors.New("The Title is empty")
		return err
	}
	if note.Text == "" {
		err := errors.New("The Text is empty")
		return err
	}

    return nil
}

func (*usecase) Create(note *entity.Note) (*entity.Note, error) {
	created, err := repo.Create(note)
	if err != nil {
		err := errors.New("Error on Create a new note")
		return nil, err
	}
    return created, nil
}

func (*usecase) FindAll() ([]entity.Note, error) {    
	notes, err := repo.FindAll()
	if err != nil {
		err := errors.New("Error on retrieve the notes")
		return nil, err
	}

	return notes, nil
}