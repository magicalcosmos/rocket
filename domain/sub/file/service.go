package file

type FileService struct {
	repository FileRepository
}

func NewFileService(repo FileRepository) *FileService {
	return &FileService{
		repository: repo,
	}
}

func (s *FileService) UploadFile(file *File) error {
	return s.repository.Save(file)
}

func (s *FileService) ReplaceFileContent(id string, newContent []byte) error {
	file, err := s.repository.FindById(id)
	if err != nil {
		return err
	}
	err = file.ReplaceContent(newContent)
	if err != nil {
		return err
	}
	return s.repository.Save(file)
}
