package controllers

import (
	"github.com/shun-shun123/clean_architecture_todo/src/app/domain"
	"github.com/shun-shun123/clean_architecture_todo/src/app/interfaces/database"
	"github.com/shun-shun123/clean_architecture_todo/src/app/usecase"
)

type ToDoController struct {
	Interactor usecase.ToDoInteractor
}

func NewToDoController(dbHandler database.DBHandler) *ToDoController {
	return &ToDoController{
		Interactor: usecase.ToDoInteractor{
			ToDoRepository: &database.ToDoRepository{
				DBHandler: dbHandler,
			},
		},
	}
}

func (controller *ToDoController) Create(c Context) {
	todo := domain.ToDo{}
	c.Bind(&todo)
	data, err := controller.Interactor.Create(todo)
	if err != nil {
		c.JSON(501, err)
		return
	}
	c.JSON(201, data)
}

func (controller *ToDoController) Read(c Context) {
	todo := domain.ToDo{}
	c.Bind(&todo)
	data, err := controller.Interactor.Read(todo)
	if err != nil {
		c.JSON(501, err)
		return
	}
	c.JSON(201, data)
}

func (controller *ToDoController) Update(c Context) {
	todo := domain.ToDo{}
	c.Bind(&todo)
	data, err := controller.Interactor.Update(todo)
	if err != nil {
		c.JSON(501, err)
		return
	}
	c.JSON(201, data)
}

func (controller *ToDoController) Delete(c Context) {
	todo := domain.ToDo{}
	c.Bind(&todo)
	data, err := controller.Interactor.Update(todo)
	if err != nil {
		c.JSON(501, err)
		return
	}
	c.JSON(201, data)
}
