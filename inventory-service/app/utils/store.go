package utils

import "database/sql"

func GetReturnedRowNumber(err error) (int, error) {
	switch err {
	case sql.ErrNoRows:
		return 0, nil
	case nil:
		return 1, nil
	default:
		return -1, err
	}
}
