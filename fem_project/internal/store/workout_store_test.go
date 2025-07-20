package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the PostgreSQL driver
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=Nirush dbname=db_fs port=5432 sslmode=disable")
	if err != nil {
		t.Fatalf("db: open: %v", err)
	}

	// run the migrations
	err = Migrate(db, "../../migrations/")
	if err != nil {
		t.Fatalf("migrating test db error: %v", err)
	}

	_, err = db.Exec(`TRUNCATE workouts, workout_entries CASCADE`)

	if err != nil {
		t.Fatalf("truncating tables %v", err)
	}
 	return db
}

func  TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresWorkoutStore(db)

	tests := []struct {
		name string
		workout *Workout
		wantErr bool
	}{
		{
			name: "valid workout",
			workout: &Workout{
				Title:           "Morning Run",
				Description:     "A quick morning run",
				DurationMinutes: 60,
				CaloriesBurned:  500,
				Entries: []WorkoutEntry{
					{
						ExerciseName:    "Running",
						Sets:            1,
						Reps:            Intptr(1),
						Weight: 		Floatptr(0),
						Notes:           "Felt great",
						OrderIndex:      1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid workout with entries",
			workout: &Workout{
				Title:           "Full body workout",
				Description:     "Comprehensive workout",
				DurationMinutes: 30,
				CaloriesBurned:  200,
				Entries: []WorkoutEntry{
					{
						ExerciseName:    "Plank",
						Sets:            3,
						Reps:            Intptr(10),
						Notes:           "Keep core tight",
						OrderIndex:      1,
					},
					{
						ExerciseName:    "Squats",
						Sets:            4,
						Reps:            Intptr(10),
						Notes:           "Focus on form",
						OrderIndex:      2,
					},
				},
			},
			wantErr: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdWorkout, err := store.CreateWorkout(tt.workout)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.workout.Title, createdWorkout.Title)
			assert.Equal(t, tt.workout.Description, createdWorkout.Description)
			assert.Equal(t, tt.workout.DurationMinutes, createdWorkout.DurationMinutes)

			retrieved, err := store.GetWorkoutByID(int64(createdWorkout.ID))
			require.NoError(t, err)

			assert.Equal(t, createdWorkout.ID, retrieved.ID)
			assert.Equal(t, len(tt.workout.Entries), len(retrieved.Entries))

			for i, _ := range tt.workout.Entries {
				assert.Equal(t, tt.workout.Entries[i].ExerciseName, retrieved.Entries[i].ExerciseName)
				assert.Equal(t, tt.workout.Entries[i].Sets, retrieved.Entries[i].Sets)
				assert.Equal(t, tt.workout.Entries[i].OrderIndex, retrieved.Entries[i].OrderIndex)
			}

		})
	}
}

func Intptr(i int) *int{
	return &i
}

func Floatptr(i float64) *float64 {
	return &i
}
