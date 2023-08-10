package datasources

import (
	"log"

	"github.com/hwoodall30/sqlite-gql/graph/model"
)

func (d *DataSource) GetAllTodos() ([]*model.Todo, error) {
	var todos []*model.Todo
	err := d.DB.DBConnection.Select(&todos, "SELECT id, text, done FROM Todos")
	if err != nil {
		log.Fatal(err)
	}
	return todos, nil
}

func (d *DataSource) GetTodoByID(id string) (*model.Todo, error) {
	var todo model.Todo
	err := d.DB.DBConnection.Get(&todo, "SELECT id, text, done FROM Todos WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (d *DataSource) CreateTodo(text string) (*model.Todo, error) {
	var todo model.Todo

	res, err := d.DB.DBConnection.Exec("INSERT INTO Todos (id, text, done) VALUES (null, ?, ?)", text, false)
	if err != nil {
		return nil, err
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = d.DB.DBConnection.Get(&todo, "SELECT id, text, done FROM Todos WHERE id = ?", insertedId)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
