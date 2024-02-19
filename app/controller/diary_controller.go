package controller

import (
	"github.com/gin-gonic/gin"
	"polus-backend/app/service"
)

type DiaryController interface {
	GetAllDiaryData(c *gin.Context)
	AddDiaryData(c *gin.Context)
	GetDiaryById(c *gin.Context)
	UpdateDiaryData(c *gin.Context)
	DeleteDiary(c *gin.Context)
}

type DiaryControllerImpl struct {
	svc service.DiaryService
}

func (u DiaryControllerImpl) GetAllDiaryData(c *gin.Context) {
	u.svc.GetAllDiary(c)
}

func (u DiaryControllerImpl) AddDiaryData(c *gin.Context) {
	u.svc.AddDiaryData(c)
}

func (u DiaryControllerImpl) GetDiaryById(c *gin.Context) {
	u.svc.GetDiaryById(c)
}

func (u DiaryControllerImpl) UpdateDiaryData(c *gin.Context) {
	u.svc.UpdateDiaryData(c)
}

func (u DiaryControllerImpl) DeleteDiary(c *gin.Context) {
	u.svc.DeleteDiary(c)
}

func DiaryControllerInit(DiaryService service.DiaryService) *DiaryControllerImpl {
	return &DiaryControllerImpl{
		svc: DiaryService,
	}
}
