package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/shun-shun123/clean_architecture_todo/src/app/domain"
	"io/ioutil"
	"os"
)

type DBHandler struct {
	file *os.File
}

func NewDBHandler() *DBHandler {
	var dbHandler = &DBHandler{}
	file, err := os.Create("db.json")
	if err != nil {
		return nil
	}
	dbHandler.file = file
	return dbHandler
}

func readDBContent() domain.ToDos {
	file, _ := os.Open("db.json")
	defer file.Close()
	buffer, err := ioutil.ReadAll(file)
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
	todos := readDBContent()
	handler.file, _ = os.Create("db.json")
	defer handler.file.Close()
	todos.Todos = append(todos.Todos, todo)
	writeData, err := json.Marshal(todos)
	_, err = handler.file.Write(writeData)
	if err != nil {
		return todo, err
	}
	return todo, err
}

func (handler *DBHandler) Read(todo domain.ToDo) (domain.ToDo, error) {
	// DBから全部読み出す
	todos := readDBContent()
	handler.file, _ = os.Create("db.json")
	defer handler.file.Close()
	err := fmt.Errorf("ID %v is not found", todo.ID)
	for _, v := range todos.Todos {
		if v.ID == todo.ID {
			todo = v
			err = nil
			break
		}
	}
	data, _ := json.Marshal(todos)
	handler.file.Write(data)
	return todo, err
}

func (handler *DBHandler) Update(todo domain.ToDo) (domain.ToDo, error) {
	todos := readDBContent()
	handler.file, _ = os.Create("db.json")
	defer handler.file.Close()
	err := fmt.Errorf("ID %v is not found", todo.ID)
	for i, v := range todos.Todos {
		fmt.Printf("ID: %v\n", v.ID)
		if v.ID == todo.ID {
			todos.Todos[i] = todo
			err = nil
			break
		}
	}
	data, _ := json.Marshal(todos)
	handler.file.Write(data)
	return todo, err
}

func (handler *DBHandler) Delete(todo domain.ToDo) (domain.ToDo, error) {
	todos := readDBContent()
	handler.file, _ = os.Create("db.json")
	defer handler.file.Close()
	err := fmt.Errorf("ID %v is not found", todo.ID)
	for i, v := range todos.Todos {
		if v.ID == todo.ID {
			fmt.Printf("ID: %v is found", todo.ID)
			todos.Todos = append(todos.Todos[:i], todos.Todos[i+1:]...)
			err = nil
			break
		}
	}
	data, marshalErr := json.Marshal(todos)
	if marshalErr != nil {
		return todo, marshalErr
	}
	_, writeErr := handler.file.Write(data)
	if writeErr != nil {
		return todo, writeErr
	}
	return todo, err
}