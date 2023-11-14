package main

import (
	"PersonalNotes/internal/controller"
	router "PersonalNotes/internal/http"
	"PersonalNotes/internal/repository"
	"PersonalNotes/internal/usecases"
	"PersonalNotes/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	var (
		noteRepository repository.NoteRepository = repository.NewFireStoreRepository(config.ProjectId, config.CollectionName)
		noteUsecase    usecases.NoteUseCases     = usecases.NewNoteUseCases(noteRepository)
		noteController controller.NoteController = controller.NewNoteController(noteUsecase)
		httpRouter     router.Router             = router.NewMuxRouter()
	)

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Up & Running...")
	})

	httpRouter.GET("/notes", noteController.GetNotes)
	httpRouter.POST("/notes", noteController.AddNotes)

	httpRouter.SERVE(config.ServerPort)
}
