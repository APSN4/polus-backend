package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"polus-backend/app/constant"
	"polus-backend/app/domain/dao"
	"polus-backend/app/pkg"
	"polus-backend/app/repository"
	"strconv"
	"time"
)

type NoteService interface {
	GetAllNote(c *gin.Context)
	GetNoteById(c *gin.Context)
	AddNoteData(c *gin.Context)
	UpdateNoteData(c *gin.Context)
	DeleteNote(c *gin.Context)
}

type NoteServiceImpl struct {
	noteRepository repository.NoteRepository
}

func (u NoteServiceImpl) UpdateNoteData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update note data by id")
	noteID, _ := strconv.Atoi(c.Param("noteID"))

	var request dao.Note
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.noteRepository.FindNoteById(noteID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.UserID = request.UserID
	data.DiaryID = request.DiaryID
	data.PhotoCloudsUrl = request.PhotoCloudsUrl
	data.NatureEvents = request.NatureEvents
	data.Temperature = request.Temperature
	data.Supplement = request.Supplement
	data.LocationX = request.LocationX
	data.LocationY = request.LocationY
	data.AddressText = request.AddressText
	data.UpdatedAt = time.Now()
	u.noteRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u NoteServiceImpl) GetNoteById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get note by id")
	noteID, _ := strconv.Atoi(c.Param("noteID"))

	data, err := u.noteRepository.FindNoteById(noteID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u NoteServiceImpl) AddNoteData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data note")
	var request dao.Note
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.noteRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u NoteServiceImpl) GetAllNote(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data note")

	data, err := u.noteRepository.FindAllNote()
	if err != nil {
		log.Error("Happened Error when find all note data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u NoteServiceImpl) DeleteNote(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data note by id")
	noteID, _ := strconv.Atoi(c.Param("noteID"))

	err := u.noteRepository.DeleteNoteById(noteID)
	if err != nil {
		log.Error("Happened Error when try delete data note from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func NoteServiceInit(noteRepository repository.NoteRepository) *NoteServiceImpl {
	return &NoteServiceImpl{
		noteRepository: noteRepository,
	}
}
