package test_case

import "time"

type TestCase struct {
	ID        string
	Name      string
	Steps     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTestCase(id, name string, steps []string) *TestCase {
	return &TestCase{
		ID:        id,
		Name:      name,
		Steps:     steps,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (tc *TestCase) AddStep(step string) {
	tc.Steps = append(tc.Steps, step)
	tc.UpdatedAt = time.Now()
}

func (tc *TestCase) RemoveStep(index int) {
	if index >= 0 && index < len(tc.Steps) {
		tc.Steps = append(tc.Steps[:index], tc.Steps[index+1:]...)
		tc.UpdatedAt = time.Now()
	}
}
