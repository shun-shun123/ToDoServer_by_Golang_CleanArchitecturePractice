package database

import "github.com/shun-shun123/clean_architecture_todo/src/app/domain"

type DBHandler interface {
	Create(todo domain.ToDo) (domain.ToDo, error)
	Read(todo domain.ToDo) (domain.ToDo, error)
	Update(todo domain.ToDo) (domain.ToDo, error)
	Delete(todo domain.ToDo) (domain.ToDo, error)
}

