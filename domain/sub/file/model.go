package file

import (
	"errors"
	"time"
)

type File struct {
	ID        string
	Name      string
	Content   []byte
	UpdatedAt time.Time
}

func NewFile(id, name string, content []byte) *File {
	return &File{
		ID:        id,
		Name:      name,
		Content:   content,
		UpdatedAt: time.Now(),
	}
}

func (f *File) ReplaceContent(newContent []byte) error {
	if len(newContent) == 0 {
		return errors.New("new content cannot be empty")
	}
	f.Content = newContent
	f.UpdatedAt = time.Now()
	return nil
}
