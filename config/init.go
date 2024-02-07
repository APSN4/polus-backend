package config

import (
	"polus-backend/app/controller"
	"polus-backend/app/repository"
	"polus-backend/app/service"
)

type Initialization struct {
	userRepo  repository.UserRepository
	userSvc   service.UserService
	UserCtrl  controller.UserController
	RoleRepo  repository.RoleRepository
	diaryRepo repository.DiaryRepository
	diarySvc  service.DiaryService
	DiaryCtrl controller.DiaryController
	noteRepo  repository.NoteRepository
	noteSvc   service.NoteService
	NoteCtrl  controller.NoteController
}

func NewInitialization(userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository,
	diaryRepo repository.DiaryRepository,
	diarySvc service.DiaryService,
	DiaryCtrl controller.DiaryController,
	noteRepo repository.NoteRepository,
	noteSvc service.NoteService,
	NoteCtrl controller.NoteController) *Initialization {
	return &Initialization{
		userRepo:  userRepo,
		userSvc:   userService,
		UserCtrl:  userCtrl,
		RoleRepo:  roleRepo,
		diaryRepo: diaryRepo,
		diarySvc:  diarySvc,
		DiaryCtrl: DiaryCtrl,
		noteRepo:  noteRepo,
		noteSvc:   noteSvc,
		NoteCtrl:  NoteCtrl,
	}
}
