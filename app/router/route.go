package router

import (
	"github.com/gin-gonic/gin"
	"polus-backend/config"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.GetAllUserData)
		user.POST("", init.UserCtrl.AddUserData)
		user.GET("/:userID", init.UserCtrl.GetUserById)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)

		userComponents := api.Group("/component")
		userComponents.GET("/:userID", init.UserCtrl.UpdateComponentUserData)

		diary := api.Group("/diary")
		diary.GET("", init.DiaryCtrl.GetAllDiaryData)
		diary.POST("", init.DiaryCtrl.AddDiaryData)
		diary.GET("/:diaryID", init.UserCtrl.GetUserById)
		diary.PUT("/:diaryID", init.UserCtrl.UpdateUserData)
		diary.DELETE("/:diaryID", init.UserCtrl.DeleteUser)

		note := api.Group("/note")
		note.GET("", init.NoteCtrl.GetAllNoteData)
		note.POST("", init.NoteCtrl.AddNoteData)
		note.GET("/:noteID", init.NoteCtrl.GetNoteById)
		note.PUT("/:noteID", init.NoteCtrl.UpdateNoteData)
		note.DELETE("/:noteID", init.NoteCtrl.DeleteNote)
	}

	return router
}
