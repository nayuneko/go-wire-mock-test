package repository

import "github.com/nayuneko/go-wire-mock-test/entity"

type EntryRepository interface {
	CreateEntry(entry *entity.Entry) (*entity.Entry, error)
}
type entryRepository struct{}

func NewEntryRepository() EntryRepository {
	return &entryRepository{}
}

func (r *entryRepository) CreateEntry(entry *entity.Entry) (*entity.Entry, error) {
	return entry, nil
}
