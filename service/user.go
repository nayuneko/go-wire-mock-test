package service

import (
	"github.com/nayuneko/go-wire-mock-test/entity"
	"github.com/nayuneko/go-wire-mock-test/repository"
)

type UserService interface {
	CreateEntry(id int64) error
}
type userService struct {
	userRepo  repository.UserRepository
	entryRepo repository.EntryRepository
}

func NewUserService(
	userRepo repository.UserRepository,
	entryRepo repository.EntryRepository,
) UserService {
	return &userService{userRepo, entryRepo}
}

func (s *userService) CreateEntry(id int64) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	entry := entity.Entry{
		UserID:   user.ID,
		UserName: user.Name,
	}
	if _, err := s.entryRepo.CreateEntry(&entry); err != nil {
		return err
	}
	return nil
}
