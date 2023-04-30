package storage

import (
	"database/sql"
)

func FindRandomQuestionID(db *sql.DB, num int) ([]int, error) {
	var questionIDs []int

	rows, err := db.Query("SELECT id FROM question ORDER BY random() LIMIT $1", num)
	if err != nil {
		return questionIDs, err
	}
	defer rows.Close()

	for rows.Next() {
		var questionID int
		err := rows.Scan(&questionID)
		if err != nil {
			return questionIDs, err
		}
		questionIDs = append(questionIDs, questionID)
	}

	return questionIDs, nil
}
