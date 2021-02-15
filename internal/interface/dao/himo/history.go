package himo

import "time"

// History は履歴の DAO
type History struct {
	ID        uint32    `db:"id, primarykey, autoincrement"`
	CreatedAt time.Time `db:"created_at"`
}
