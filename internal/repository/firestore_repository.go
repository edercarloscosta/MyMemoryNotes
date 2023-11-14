package repository

import (
	"PersonalNotes/internal/entity"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

type repo struct {
	ProjectId      string
	CollectionName string
}

func NewFireStoreRepository(projId, collName string) NoteRepository {
	return &repo{
		ProjectId:      projId,
		CollectionName: collName,
	}
}

func (r *repo) Create(note *entity.Note) (*entity.Note, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, r.ProjectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(r.CollectionName).Add(ctx, map[string]interface{}{
		"ID":    uuid.New().String(),
		"Title": note.Title,
		"Text":  note.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new note: %v", err)
		return nil, err
	}

	return note, nil
}

func (r *repo) FindAll() ([]entity.Note, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, r.ProjectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var notes []entity.Note
	it := client.Collection(r.CollectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed getting a note: %v", err)
			return nil, err
		}

		note := entity.Note{
			ID:    (doc.Data()["ID"]).(string),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		notes = append(notes, note)
	}
	return notes, nil
}
