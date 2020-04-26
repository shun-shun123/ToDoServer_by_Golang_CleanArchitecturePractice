package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/shun-shun123/clean_architecture_todo/src/app/domain"
	"io"
	"io/ioutil"
	"os"
)

type DBHandler struct {
	io.Writer
	io.Reader
}

func NewDBHandler() *DBHandler {
	var dbHandler = &DBHandler{}
	file, err := os.Create("db.json")
	if err != nil {
		return nil
	}
	dbHandler.Writer = file
	dbHandler.Reader = file
	return dbHandler
}

func (handler *DBHandler) readDBContent() domain.ToDos {
	buffer, err := ioutil.ReadAll(handler.Reader)
	if err != nil {
		return domain.ToDos{}
	}
	todos := domain.ToDos{}
	err = json.Unmarshal(buffer, &todos)
	if err != nil {
		todos = domain.ToDos{}
	}
	return todos
}

func (handler *DBHandler) Create(todo domain.ToDo) (domain.ToDo, error) {
	todos := handler.readDBContent()
	todos.Todos = append(todos.Todos, todo)
	writeData, err := json.Marshal(todos)
	_, err = handler.Writer.Write(writeData)
	if err != nil {
		return todo, err
	}
	return todo, err
}

func (handler *DBHandler) Read(todo domain.ToDo) (domain.ToDo, error) {
	// DBから全部読み出す
	todos := handler.readDBContent()
	for _, v := range todos.Todos {
		if v.ID == todo.ID {
			return v, nil
		}
	}
	err := fmt.Errorf("ID %v is not found", todo.ID)
	return todo, err
}

func (handler *DBHandler) Update(todo domain.ToDo) (domain.ToDo, error) {
	todos := handler.readDBContent()
	err := fmt.Errorf("ID %v is not found", todo.ID)
	for i, v := range todos.Todos {
		if v.ID == todo.ID {
			todos.Todos[i] = todo
			err = nil
			break
		}
	}
	return todo, err
}

func (handler *DBHandler) Delete(todo domain.ToDo) (domain.ToDo, error) {
	todos := handler.readDBContent()
	err := fmt.Errorf("ID %v is not found", todo.ID)
	for i, v := range todos.Todos {
		if v.ID == todo.ID {
			todos.Todos = append(todos.Todos[:i], todos.Todos[i+1:]...)
			break
		}
	}
	data, err := json.Marshal(todos)
	if err != nil {
		return todo, err
	}
	_, err = handler.Writer.Write(data)
	if err != nil {
		return todo, err
	}
	return todo, err
}