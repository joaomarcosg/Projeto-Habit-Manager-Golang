package entity

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
}
