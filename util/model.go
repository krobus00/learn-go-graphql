package util

import "database/sql"

func NewNullString(s string, valid bool) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  valid,
	}
}
