package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ByteNirush/fem_project/internal/store"
	"github.com/ByteNirush/fem_project/internal/utils"
	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
	logger       *log.Logger
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, logger *log.Logger) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
		logger:       logger,
	}
}

func (wh *WorkoutHandler) HandlerGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	workoutID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("ERROR: readIDParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid workout id"})
		return
	}	

	workout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		wh.logger.Printf("ERROR: getWorkoutByID: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "invalid sever error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": workout})

}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request){
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)

	if err != nil {
		wh.logger.Printf("ERROR: decodingCreateWorkout: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request sent"})
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		wh.logger.Printf("ERROR: createdWorkout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "failed to create workout"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": createdWorkout})
}

func (wh *WorkoutHandler) HandleGetAllWorkouts(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement getting all workouts
	fmt.Fprintf(w, "get all workouts\n")
}

func (wh *WorkoutHandler) HandleUpdateWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	exitingWorkout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		http.Error(w, "failed to get workout", http.StatusInternalServerError)
		return
	}

	if exitingWorkout == nil {
		http.NotFound(w, r)
		return
	}

	// at this point we asume we are able to find an existing workout
	var updatedWorkoutRequest struct {
		Title 		  *string `json:"title"`
		Description   *string `json:"description"`
		DurationMinutes *int `json:"duration_minutes"`
		CaloriesBurned *int `json:"calories_burned"`
		Entries       *[]store.WorkoutEntry `json:"entries"`
	} 

	err = json.NewDecoder(r.Body).Decode(&updatedWorkoutRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatedWorkoutRequest.Title != nil {
		exitingWorkout.Title = *updatedWorkoutRequest.Title
	}

	if updatedWorkoutRequest.Description != nil {
		exitingWorkout.Description = *updatedWorkoutRequest.Description
	}

	if updatedWorkoutRequest.DurationMinutes != nil {
		exitingWorkout.DurationMinutes = *updatedWorkoutRequest.DurationMinutes
	}

	if updatedWorkoutRequest.CaloriesBurned != nil {
		exitingWorkout.CaloriesBurned = *updatedWorkoutRequest.CaloriesBurned
	}

	if updatedWorkoutRequest.Entries != nil {
		exitingWorkout.Entries = *updatedWorkoutRequest.Entries
	}

	err = wh.workoutStore.UpdateWorkout(exitingWorkout)
	if err != nil {
		fmt.Println("update workout error",err)
		http.Error(w, "failed to update workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exitingWorkout)
}


func (wh *WorkoutHandler) HandleDeleteWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = wh.workoutStore.DeleteWorkout(workoutID)
	if err == sql.ErrNoRows {
		http.Error(w, "workout not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "error deleting workout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}