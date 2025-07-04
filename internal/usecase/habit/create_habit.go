package habit

import (
	"context"
	"time"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/validator"
)

type CreateHabitReq struct {
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Frequency   string    `json:"frequency"`
	StartDate   time.Time `json:"start_date"`
	TargetDate  time.Time `json:"target_date"`
	Priority    uint8     `json:"priority"`
}

func (req CreateHabitReq) Valid(ctx context.Context) validator.Evaluator {

	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.Name), "name", "this field cannot be empty")
	eval.CheckField(validator.NotBlank(req.Category), "category", "this field cannot be empty")
	eval.CheckField(validator.NotBlank(req.Description), "description", "this field cannot be empty")
	eval.CheckField(
		validator.MinChars(req.Description, 10) &&
			validator.MaxChars(req.Description, 255), "description", "this field must have a length between 10 and 255")
	eval.CheckField(validator.NotBlank(req.Frequency), "frequency", "this field cannot be empty")
	eval.CheckField(
		validator.MinPriority(req.Priority) &&
			validator.MaxPriority(req.Priority), "priority", "this field must have a length between 1 and 10")

	return eval
}
