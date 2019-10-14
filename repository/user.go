package repository

import "github.com/nayuneko/go-wire-mock-test/entity"

type UserRepository interface {
	FindByID(id int64) (*entity.User, error)
}
type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByID(id int64) (*entity.User, error) {
	return &entity.User{ID: id, Name: "name"}, nil
}
