package repository

import (
	"PersonalNotes/internal/entity"
)

type NoteRepository interface {
	Create(note *entity.Note) (*entity.Note, error)
	FindAll() ([]entity.Note, error)
}