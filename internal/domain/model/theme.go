package model

// Theme はお題
type Theme struct {
	ID       int64
	Sentence string
	Creator  User
}
