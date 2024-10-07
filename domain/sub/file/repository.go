package file

type FileRepository interface {
	Save(file *File) error
	FindById(id string) (*File, error)
	DeleteById(id string) error
}
