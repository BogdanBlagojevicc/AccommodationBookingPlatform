package service

import (
	"user-service/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}
