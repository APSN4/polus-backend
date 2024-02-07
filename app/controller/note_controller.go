package controller

import (
	"github.com/gin-gonic/gin"
	"polus-backend/app/service"
)

type NoteController interface {
	GetAllNoteData(c *gin.Context)
	AddNoteData(c *gin.Context)
	GetNoteById(c *gin.Context)
	UpdateNoteData(c *gin.Context)
	DeleteNote(c *gin.Context)
}

type NoteControllerImpl struct {
	svc service.NoteService
}

func (u NoteControllerImpl) GetAllNoteData(c *gin.Context) {
	u.svc.GetAllNote(c)
}

func (u NoteControllerImpl) AddNoteData(c *gin.Context) {
	u.svc.AddNoteData(c)
}

func (u NoteControllerImpl) GetNoteById(c *gin.Context) {
	u.svc.GetNoteById(c)
}

func (u NoteControllerImpl) UpdateNoteData(c *gin.Context) {
	u.svc.UpdateNoteData(c)
}

func (u NoteControllerImpl) DeleteNote(c *gin.Context) {
	u.svc.DeleteNote(c)
}

func NoteControllerInit(noteService service.NoteService) *NoteControllerImpl {
	return &NoteControllerImpl{
		svc: noteService,
	}
}
