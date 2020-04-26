package database

import (
	"github.com/shun-shun123/clean_architecture_todo/src/app/domain"
)

type ToDoRepository struct {
	DBHandler
}

func (repo *ToDoRepository) Create(todo domain.ToDo) (domain.ToDo, error) {
	// インフラ層のDBHandlerだけど呼び出してるのはinterfaces層で定義されたインタフェース
	// 「依存は外から内のみ」という原則を守ってる
	result, err := repo.DBHandler.Create(todo)
	if err != nil {
		return result, err
	}
	return result, err
}

func (repo *ToDoRepository) Read(todo domain.ToDo) (domain.ToDo, error) {
	result, err := repo.DBHandler.Read(todo)
	if err != nil {
		return result, err
	}
	return result, err
}

func (repo *ToDoRepository) Update(todo domain.ToDo) (domain.ToDo, error) {
	result, err := repo.DBHandler.Update(todo)
	if err != nil {
		return result, err
	}
	return result, err
}

func (repo *ToDoRepository) Delete(todo domain.ToDo) (domain.ToDo, error) {
	result, err := repo.DBHandler.Delete(todo)
	if err != nil {
		return result, err
	}
	return result, err
}