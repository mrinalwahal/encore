package user

import "time"

type (
	User struct {
		ID        int64     `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		FullName  string    `json:"full_name"`
		Active    bool      `json:"active"`
	}
)
