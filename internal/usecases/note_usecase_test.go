package usecases

import (
	"PersonalNotes/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}	

func (mock *mockRepository) Create(note *entity.Note) (*entity.Note, error){
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Note), args.Error(1)
} 

func (mock *mockRepository) FindAll() ([]entity.Note, error){
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Note), args.Error(1)
} 

func TestValidateEmptyPost(t *testing.T) {
	testUseCase := NewNoteUseCases(nil)

	err := testUseCase.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The note is empty", err.Error())
}	

func TestValidateEmptyPostTitle(t *testing.T){
	note := entity.Note{
		ID: "1",
		Title: "",
		Text: "text",
	} 

	testUseCase := NewNoteUseCases(nil)
	
	err := testUseCase.Validate(&note)

	assert.NotNil(t, err)
	assert.Equal(t, "The Title is empty", err.Error())
} 

func TestValidateEmptyPostText(t *testing.T){
	note := entity.Note{
		ID: "1",
		Title: "title",
		Text: "",
	} 

	testUseCase := NewNoteUseCases(nil)
	
	err := testUseCase.Validate(&note)

	assert.NotNil(t, err)
	assert.Equal(t, "The Text is empty", err.Error())
} 

func TestFindAll(t *testing.T){
	mockRepo := new(mockRepository)

	note := entity.Note{
		ID: "1",
		Title: "title",
		Text: "text",
	} 
	mockRepo.On("FindAll").Return([]entity.Note{note}, nil)

	testService := NewNoteUseCases(mockRepo)
	result, _ := testService.FindAll()
	
	mockRepo.AssertExpectations(t)

	assert.Equal(t, note.ID, result[0].ID)
	assert.Equal(t, note.Title, result[0].Title)
	assert.Equal(t, note.Text, result[0].Text)
} 

func TestCreate(t *testing.T){
	mockRepo := new(mockRepository)

	note := entity.Note{
		ID: "1",
		Title: "title",
		Text: "text",
	} 
	mockRepo.On("Create").Return(&note, nil)

	testService := NewNoteUseCases(mockRepo)
	result, _ := testService.Create(&note)
	
	mockRepo.AssertExpectations(t)

	assert.Equal(t, note.ID, result.ID)
	assert.Equal(t, note.Title, result.Title)
	assert.Equal(t, note.Text, result.Text)

} 