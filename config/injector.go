// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"polus-backend/app/controller"
	"polus-backend/app/repository"
	"polus-backend/app/service"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var roleRepoSet = wire.NewSet(repository.RoleRepositoryInit,
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
)

var diaryServiceSet = wire.NewSet(service.DiaryServiceInit,
	wire.Bind(new(service.DiaryService), new(*service.DiaryServiceImpl)),
)

var diaryRepoSet = wire.NewSet(repository.DiaryRepositoryInit,
	wire.Bind(new(repository.DiaryRepository), new(*repository.DiaryRepositoryImpl)),
)

var diaryCtrlSet = wire.NewSet(controller.DiaryControllerInit,
	wire.Bind(new(controller.DiaryController), new(*controller.DiaryControllerImpl)),
)

var noteServiceSet = wire.NewSet(service.NoteServiceInit,
	wire.Bind(new(service.NoteService), new(*service.NoteServiceImpl)),
)

var noteRepoSet = wire.NewSet(repository.NoteRepositoryInit,
	wire.Bind(new(repository.NoteRepository), new(*repository.NoteRepositoryImpl)),
)

var noteCtrlSet = wire.NewSet(controller.NoteControllerInit,
	wire.Bind(new(controller.NoteController), new(*controller.NoteControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, roleRepoSet, diaryRepoSet, diaryCtrlSet, diaryServiceSet, noteRepoSet, noteServiceSet, noteCtrlSet)
	return nil
}
