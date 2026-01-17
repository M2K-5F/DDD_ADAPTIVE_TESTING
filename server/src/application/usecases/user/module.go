package user_usercases

import "adaptivetesting/src/application/interfaces"

type DependsModule struct {
	UserRdr interfaces.IUserReader
	Writer  interfaces.IWriter
}

type UserUseCases struct {
	deps *DependsModule
}

func New(depsModule *DependsModule) *UserUseCases {
	return &UserUseCases{
		deps: depsModule,
	}
}
