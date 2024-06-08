package main

import (
  "net/http"
  "io"

  "github.com/gorilla/mux"
)

type TasksService struct {
  store Store
}

func NewTasksService(s Store) *TasksService {
  return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
  r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
  r.HandleFunc("/tasks/{id}", s.handleCreateTask).Methods("GET")
}

func (s *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
  body, err := io.ReadAll(r.Body)
  if err != nil {
    return
  }

  defer r.Body.Close()



}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}
