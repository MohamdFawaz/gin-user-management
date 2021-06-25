package repository

import (
	"gin-user-management/lib"
)

type UserRepository struct {
	lib.Database
	logger lib.Logger
}

func NewUserRepository(database lib.Database, logger lib.Logger) UserRepository {
	return UserRepository{
		Database: database,
		logger: logger,
	}
}
