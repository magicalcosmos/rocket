package application

import "project-with-ddd/domain/test_case"

type TestCaseAppService struct {
	testCaseService *test_case.TestCaseService
}

func NewTestCaseAppService(service *test_case.TestCaseService) *TestCaseAppService {
	return &TestCaseAppService{
		testCaseService: service,
	}
}

func (s *TestCaseAppService) CreateTestCase(id, name string, steps []string) error {
	testCase := test_case.NewTestCase(id, name, steps)
	return s.testCaseService.Save(testCase)
}

func (s *TestCaseAppService) DeleteTestCase(id string) error {
	return s.testCaseService.Delete(id)
}
