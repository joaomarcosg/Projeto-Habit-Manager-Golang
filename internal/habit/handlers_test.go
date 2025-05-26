package habit

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
)

type MockHabitRepository struct {
	Habits []entity.Habit
}

func (m *MockHabitRepository) CreateHabit(ctx context.Context, habit entity.Habit) (int64, error) {
	habit.ID = int64(len(m.Habits) + 1)
	m.Habits = append(m.Habits, habit)
	return habit.ID, nil
}

func (m *MockHabitRepository) ListHabits(ctx context.Context) ([]entity.Habit, error) {
	return m.Habits, nil
}

func (m *MockHabitRepository) DeleteHabit(ctx context.Context, id int64) (bool, error) {
	for i, h := range m.Habits {
		if h.ID == id {
			m.Habits = append(m.Habits[:i], m.Habits[i+1:]...)
			return true, nil
		}
	}

	return false, nil
}

func TestCreateHabit(t *testing.T) {

	mockRepo := &MockHabitRepository{}
	service := NewService(mockRepo)

	handler := handleCreateHabit(service)

	payload := []byte(
		`{
			"name": "ler",
            "category": "estudos",
            "description": "ler todos os dias",
            "frequency": "sunday,monday,tuesday,wednesday,thursday,friday,saturday",
            "start_date": "2025-05-24T00:00:00Z",
            "target_date": "2025-06-07T00:00:00Z",
            "priority": 10
		}`)

	req, err := http.NewRequest("POST", "/habits", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusCreated {
		t.Errorf("incorrect status: receveid %v, expected %v", status, http.StatusCreated)
	}

	if len(mockRepo.Habits) != 1 {
		t.Errorf("expected 1 habit in the repository, but there is %d", len(mockRepo.Habits))
	}

}
