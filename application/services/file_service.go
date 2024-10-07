package application

import "project-with-ddd/domain/file"

type FileAppService struct {
	fileService *file.FileService
}

func NewFileAppService(fileService *file.FileService) *FileAppService {
	return &FileAppService{
		fileService: fileService,
	}
}

func (s *FileAppService) UploadFile(id, name string, content []byte) error {
	newFile := file.NewFile(id, name, content)
	return s.fileService.UploadFile(newFile)
}

func (s *FileAppService) ReplaceFileContent(id string, newContent []byte) error {
	return s.fileService.ReplaceFileContent(id, newContent)
}
