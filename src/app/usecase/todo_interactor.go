package usecase

import "github.com/shun-shun123/clean_architecture_todo/src/app/domain"

type ToDoInteractor struct {
	ToDoRepository ToDoRepository
}

func (interactor *ToDoInteractor) Create(todo domain.ToDo) (domain.ToDo, error) {
	data, err := interactor.ToDoRepository.Create(todo)
	return data, err
}

func (interactor *ToDoInteractor) Read(todo domain.ToDo) (domain.ToDo, error) {
	data, err := interactor.ToDoRepository.Read(todo)
	return data, err
}

func (interactor *ToDoInteractor) Update(todo domain.ToDo) (domain.ToDo, error) {
	data, err := interactor.ToDoRepository.Update(todo)
	return data, err
}

func (interactor *ToDoInteractor) Delete(todo domain.ToDo) (domain.ToDo, error) {
	data, err := interactor.ToDoRepository.Delete(todo)
	return data, err
}