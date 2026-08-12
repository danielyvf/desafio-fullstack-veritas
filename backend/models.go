package main

import (
	"errors"
	"strings"
)

type Status string

const (
	StatusTodo       Status = "FAZER"
	StatusInProgress Status = "EM_PROGRESSO"
	StatusDone       Status = "CONCLUIDAS"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"titulo"`
	Description string `json:"descricao"`
	Status      Status `json:"status"`
}

func (t *Task) Validate() error {
	if strings.TrimSpace(t.Title) == "" {
		return errors.New("o título é obrigatório")
	}

	if t.Status != StatusTodo && t.Status != StatusInProgress && t.Status != StatusDone {
		return errors.New("status inválido")
	}

	return nil
}
