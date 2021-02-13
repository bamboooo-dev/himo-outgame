package himo

// Theme はお題の DAO
type Theme struct {
	ID       int64  `db:"id, primarykey, autoincrement"`
	Sentence string `db:"sentence"`
	UserID   int64  `db:"user_id"`
}
