package habit

import (
	"context"
	"time"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/validator"
)

type UpdateHabitStatusReq struct {
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

func (req UpdateHabitStatusReq) Valid(ctx context.Context) validator.Evaluator {

	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.Status), "status", "this field cannot be empty")
	eval.CheckField(validator.HabitStatus(req.Status), "status", "must be done or not_done")

	return eval

}
