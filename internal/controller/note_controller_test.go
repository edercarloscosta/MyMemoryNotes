package controller

import (
	"PersonalNotes/internal/entity"
	"PersonalNotes/internal/usecases"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNoteRepository struct {
	mock.Mock
}

func (m *MockNoteRepository) Create(note *entity.Note) (*entity.Note, error) {
	args := m.Called(note)
	return args.Get(0).(*entity.Note), args.Error(1)
}

func (m *MockNoteRepository) FindAll() ([]entity.Note, error) {
	args := m.Called()
	return args.Get(0).([]entity.Note), args.Error(1)
}

func TestAddNote(t *testing.T) {
	mockRepo := new(MockNoteRepository)

	noteMockExpected := &entity.Note{
		ID:    "1",
		Title: "Title 1",
		Text:  "Text 1",
	}

	mockRepo.On("Create", mock.AnythingOfType("*entity.Note")).Return(noteMockExpected, nil)

	var (
		noteUseCase    usecases.NoteUseCases = usecases.NewNoteUseCases(mockRepo)
		noteController NoteController        = NewNoteController(noteUseCase)
	)

	var jsonNote = []byte(`{
		"id": "1",
		"title": "Title 1",
		"text": "Text 1"
	}`)

	req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(jsonNote))

	handler := http.HandlerFunc(noteController.AddNotes)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	statusCode := response.Code

	if statusCode != http.StatusOK {
		t.Errorf("The handler returned a not expected status code: Got %v want %v", statusCode, http.StatusOK)
	}

	var note entity.Note
	json.NewDecoder(io.Reader(response.Body)).Decode(&note)

	assert.NotNil(t, noteMockExpected.ID)
	assert.Equal(t, noteMockExpected.Title, note.Title)
	assert.Equal(t, noteMockExpected.Text, note.Text)
}

func TestGetNotes(t *testing.T) {
	mockRepo := new(MockNoteRepository)

	mockRepo.On("FindAll").Return([]entity.Note{
		{
			ID:    "1",
			Title: "Title 1",
			Text:  "Text 1",
		},
		{
			ID:    "2",
			Title: "Title 2",
			Text:  "Text 2",
		},
	}, nil)

	var (
		noteUseCase    usecases.NoteUseCases = usecases.NewNoteUseCases(mockRepo)
		noteController NoteController        = NewNoteController(noteUseCase)
	)

	req, _ := http.NewRequest("GET", "/notes", nil)

	handler := http.HandlerFunc(noteController.GetNotes)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	statusCode := response.Code

	if statusCode != http.StatusOK {
		t.Errorf("The handler returned a not expected status code: Got %v want %v", statusCode, http.StatusOK)
	}

	var notes []entity.Note
	json.NewDecoder(io.Reader(response.Body)).Decode(&notes)

	for _, n := range notes {
		assert.NotNil(t, n)
		assert.NotNil(t, n.Title)
		assert.NotNil(t, n.Text)
	}
}
