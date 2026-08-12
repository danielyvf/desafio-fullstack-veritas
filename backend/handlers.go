package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	tasks   = make(map[string]Task)
	tasksMu sync.RWMutex
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	tasksMu.RLock()
	taskList := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	tasksMu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskList)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if newTask.Status == "" {
		newTask.Status = StatusTodo
	}

	if err := newTask.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTask.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	tasksMu.Lock()
	tasks[newTask.ID] = newTask
	tasksMu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func TaskDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		http.Error(w, "ID é obrigatório", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		var updatedTask Task
		if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		updatedTask.ID = id
		if err := updatedTask.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tasksMu.Lock()
		if _, exists := tasks[id]; !exists {
			tasksMu.Unlock()
			http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
			return
		}
		tasks[id] = updatedTask
		tasksMu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedTask)

	case http.MethodDelete:
		tasksMu.Lock()
		if _, exists := tasks[id]; !exists {
			tasksMu.Unlock()
			http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
			return
		}
		delete(tasks, id)
		tasksMu.Unlock()

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
