package todos

import (
	"github.com/jmoiron/sqlx"
)

type TodoStorage struct {
	Conn *sqlx.DB
}

type Todo struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Completed   bool   `db:"completed"`
}

type NewTodo struct {
	Title       string
	Description string
}

type UpdateTodo struct {
	Id          int
	Title       string
	Description string
	Completed   bool
}

type CompletTodo struct {
	Id        int
	Completed bool
}

func NewTodoStorage(conn *sqlx.DB) *TodoStorage {
	return &TodoStorage{Conn: conn}
}

// Create Todo
func (s *TodoStorage) CreateNewTodo(data NewTodo) (int, error) {
	statement := "INSERT INTO todos (title, description, completed) VALUES (?, ?, ?)"
	res, err := s.Conn.Exec(statement, data.Title, data.Description, false)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get all todos
func (s *TodoStorage) GetAllTodos() ([]Todo, error) {
	statement := "SELECT * FROM todos"
	var todos []Todo

	err := s.Conn.Select(&todos, statement)

	if err != nil {
		return nil, err
	}

	return todos, err

}

// Delete TODO
func (s *TodoStorage) DeleteTodo(id int) error {
	statement := "DELETE FROM todos WHERE Id = ?"

	_, err := s.Conn.Exec(statement, id)

	if err != nil {
		return err
	}

	return err
}

// Update TODO

func (s *TodoStorage) UpdateTodo(data UpdateTodo) error {
	statement := "UPDATE todos SET title = ?, description = ?, completed = ? WHERE id = ?"

	_, err := s.Conn.Exec(statement, data.Title, data.Description, data.Completed, data.Id)

	if err != nil {
		return err
	}

	return nil
}

// Complete TODO

func (s *TodoStorage) CompleteTodo(data CompletTodo) error {
	statement := "UPDATE todos SET completed = ? WHERE Id = ?"

	_, err := s.Conn.Exec(statement, data.Completed, data.Id)

	if err != nil {
		return err
	}

	return nil
}
