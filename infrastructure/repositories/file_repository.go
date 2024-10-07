package infrastructure

import (
	"errors"
	"rocket/domain/file"
)

type InMemoryFileRepository struct {
	files map[string]*file.File
}

func NewInMemoryFileRepository() *InMemoryFileRepository {
	return &InMemoryFileRepository{
		files: make(map[string]*file.File),
	}
}

func (repo *InMemoryFileRepository) Save(f *file.File) error {
	repo.files[f.ID] = f
	return nil
}

func (repo *InMemoryFileRepository) FindById(id string) (*file.File, error) {
	if file, exists := repo.files[id]; exists {
		return file, nil
	}
	return nil, errors.New("file not found")
}

func (repo *InMemoryFileRepository) DeleteById(id string) error {
	delete(repo.files, id)
	return nil
}
