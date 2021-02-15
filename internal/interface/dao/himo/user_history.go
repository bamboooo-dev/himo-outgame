package himo

// UserHistory はユーザーと履歴を結びつける DAO
type UserHistory struct {
	UserID    uint32 `db:"user_id"`
	HistoryID uint32 `db:"history_id"`
}
