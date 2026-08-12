package main

import (
	"fmt"
	"net/http"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetTasksHandler(w, r)
		} else if r.Method == http.MethodPost {
			CreateTaskHandler(w, r)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	}))

	mux.HandleFunc("/tasks/", enableCORS(TaskDetailHandler))

	fmt.Println("Servidor backend Go rodando na porta :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
