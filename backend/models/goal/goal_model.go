package goal_model

import (
	"database/sql"
	"fmt"
	"log"
)

type GetGoal struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	TargetPerDay string `json:"target_per_day"`
	LongTermTarget string `json:"long_term_target"`
	UserID int64 `json:"user_id"`
}

func Find(db *sql.DB, id int64) (goal GetGoal, err error) {
	log.Println("finding goal...")

	query := "SELECT id, name, target_per_day, long_term_target, user_id FROM goal WHERE id = $1"

	if err = db.QueryRow(query, id).Scan(&goal.ID, &goal.Name, &goal.TargetPerDay, &goal.LongTermTarget, &goal.UserID); err != nil {
		err = fmt.Errorf("find goal:%v", err)
		return
	}

	return
}

func Create(db *sql.DB, name string, targetPerDay string, longTermTarget string, userID int64) (goalID int64, err error) {
	log.Println("inserting goal...")

	tx, err := db.Begin()

	if err != nil {
		return goalID, err
	}

	query := "INSERT INTO goal (name, target_per_day, long_term_target, user_id) VALUES ($1, $2, $3, $4) RETURNING id"

	{
		stmt, err := tx.Prepare(query)

		if err != nil {
			return goalID, err
		}

		defer stmt.Close()

		err = stmt.QueryRow(name, targetPerDay, longTermTarget, userID).Scan(&goalID)

		if err != nil {
			return goalID, err
		}
	}

	{
		err = tx.Commit()

		if err != nil {
			return goalID, err
		}
	}

	return goalID, nil
}
