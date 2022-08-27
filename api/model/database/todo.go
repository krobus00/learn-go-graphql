package database

type Todo struct {
	ID     string `db:"id"`
	Text   string `db:"text"`
	IsDone bool   `db:"is_done"`
	DateColumn
}
