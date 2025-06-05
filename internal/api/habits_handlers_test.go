package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

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

type HabitResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Frequency   string    `json:"frequency"`
	StartDate   time.Time `json:"start_date"`
	TargetDate  time.Time `json:"target_date"`
	Priority    int       `json:"priority"`
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

func TestListHabits(t *testing.T) {

	startDate := time.Date(2025, 5, 28, 0, 0, 0, 0, time.UTC)
	targetDate := time.Date(2025, 6, 28, 0, 0, 0, 0, time.UTC)

	mockRepo := &MockHabitRepository{

		Habits: []entity.Habit{
			{
				ID:          1,
				Name:        "ler",
				Category:    "estudos",
				Description: "ler livro",
				Frequency:   "sunday,monday,tuesday,wednesday,thursday,friday,saturday",
				StartDate:   startDate,
				TargetDate:  targetDate,
				Priority:    10,
			},
			{
				ID:          2,
				Name:        "treinar",
				Category:    "saúde",
				Description: "ir para academia",
				Frequency:   "monday,tuesday,wednesday,thursday,friday",
				StartDate:   startDate,
				TargetDate:  targetDate,
				Priority:    10,
			},
		},
	}

	service := NewService(mockRepo)
	handler := handleListHabits(service)

	req, err := http.NewRequest("GET", "/habits/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("incorrect status: receveid %v, expected %v", status, http.StatusOK)
	}

	var resp struct {
		Data  []HabitResponse `json:"data"`
		Error string          `json:"error,omitempty"`
	}

	err = json.Unmarshal(response.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	got := resp.Data

	expected := []HabitResponse{
		{
			ID:          1,
			Name:        "ler",
			Category:    "estudos",
			Description: "ler livro",
			Frequency:   "sunday,monday,tuesday,wednesday,thursday,friday,saturday",
			StartDate:   startDate,
			TargetDate:  targetDate,
			Priority:    10,
		},
		{
			ID:          2,
			Name:        "treinar",
			Category:    "saúde",
			Description: "ir para academia",
			Frequency:   "monday,tuesday,wednesday,thursday,friday",
			StartDate:   startDate,
			TargetDate:  targetDate,
			Priority:    10,
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("incorrect response\nreceived %#v\nexpected %#v", got, expected)
	}

}
