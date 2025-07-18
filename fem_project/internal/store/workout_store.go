package store

import (
	"database/sql"
)

type Workout struct {
	ID 					int `json:"id"`
	Title 				string `json:"title"`
	Description 		string `json:"description"`
	DurationMinutes 	int `json:"duration_minutes"` 
	CaloriesBurned 		int `json:"calories_burned"`
	Entries 			[]WorkoutEntry `json:"entries"`
}

type WorkoutEntry struct {
	ID 					int `json:"id"`
	ExerciseName 		string `json:"exercise_name"`
	Sets   				int `json:"sets"`
	Reps    			*int `json:"reps"`
	DurationSeconds 	*int `json:"duration_seconds"`
	Weight 				*float64 `json:"weight"`
	Notes 				string `json:"notes"`
	OrderIndex 			int `json:"order_index"`
}

type PostgresWorkoutStore struct {
	db *sql.DB
}

func NewPostgresWorkoutStore(db *sql.DB) *PostgresWorkoutStore {
	return &PostgresWorkoutStore{db: db}
}

type WorkoutStore interface {
	CreateWorkout(*Workout) (*Workout, error)
	GetWorkoutByID(id int64) (*Workout, error)
}

func (pg *PostgresWorkoutStore) CreateWorkout(Workout *Workout) (*Workout, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := 
	`INSERT INTO workouts (title, description, duration_minutes, calories_burned)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	err = tx.QueryRow(query, Workout.Title, Workout.Description, Workout.DurationMinutes, Workout.CaloriesBurned).Scan(&Workout.ID)
	if err != nil {
		return nil, err
	}

	// Insert entries
	for _, entry := range Workout.Entries {
		query :=
		`INSERT INTO workout_entries (workout_id, exercise_name, sets, reps, duration_seconds, weight, notes, order_index)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

		err = tx.QueryRow(query, Workout.ID, entry.ExerciseName, entry.Sets, entry.Reps, entry.DurationSeconds, entry.Weight, entry.Notes, entry.OrderIndex).Scan(&entry.ID)
		if err != nil {
			return nil, err
		}
	}
	
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return Workout, nil
}

func (pg *PostgresWorkoutStore) GetWorkoutByID(id int64) (*Workout, error) {
	Workout := &Workout{}

	return Workout, nil
}
