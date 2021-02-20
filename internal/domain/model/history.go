package model

import "time"

// History は遊んだ履歴
type History struct {
	ID        uint32
	CreatedAt time.Time
	WithUsers []User
}
