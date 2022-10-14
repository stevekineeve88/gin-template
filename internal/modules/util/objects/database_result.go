package objects

import "database/sql"

type DatabaseResult struct {
	Status       bool
	Message      string
	Data         *sql.Rows
	InsertId     int64
	AffectedRows int64
}
