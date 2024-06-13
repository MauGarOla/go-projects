package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Title   string `json:Title`
	Content string `json:Content`
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
	json.NewEncoder(w).Encode(exampleTasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(exampleTasks) + 1
	exampleTasks = append(exampleTasks, newTask)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks)

	log.Fatal(http.ListenAndServe(":3000", router))
}
