package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"ID"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

type allTask []task

var exampleTasks = allTask{
	{
		ID:      1,
		Title:   "First task",
		Content: "This is the first task.",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST API.")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(exampleTasks)
}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Insert a valid ID", http.StatusBadRequest)
		return
	}
	for _, task := range exampleTasks {
		if taskID == task.ID {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Insert a valid Task", http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(exampleTasks) + 1
	exampleTasks = append(exampleTasks, newTask)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func deleteTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Insert a valid ID", http.StatusBadRequest)
		return
	}
	for i, task := range exampleTasks {
		if taskID == task.ID {
			exampleTasks = append(exampleTasks[:i], exampleTasks[i+1:]...)
			fmt.Fprintf(w, "The Task with the ID %v has been remove succesfully", taskID)
		}
	}
}

func updateTaskByID(w http.ResponseWriter, r *http.Request) {
	var updatedTask task
	vars := mux.Vars(r)

	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Insert a valid ID", http.StatusBadRequest)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Insert a valid Task", http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range exampleTasks {
		if taskID == task.ID {
			updatedTask.ID = taskID
			before := append(exampleTasks[:i], updatedTask)
			exampleTasks = append(before, exampleTasks[i+1:]...)
			fmt.Fprintf(w, "The Task with the ID %v has been updated succesfully", taskID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTaskByID).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTaskByID).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTaskByID).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
