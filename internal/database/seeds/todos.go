package database

import (
	"github.com/jaswdr/faker"
	todos "github.com/jimvid/ripple/internal/app/todo"
	"log"
)

// Seed 20 todos
func (s *Seed) SeedTodo() {
	faker := faker.New()

	// Create table if it does not exists
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS todos (
        id          INT AUTO_INCREMENT PRIMARY KEY,
        title       VARCHAR(255),
        description TEXT,
        completed   BOOLEAN
    );`

	// Execute the SQL statement
	_, err := s.db.Exec(createTableSQL)

	if err != nil {
		log.Fatalf("error seeding todos")
	}

	log.Print("Created todos table")

	storage := todos.NewTodoStorage(s.db)

	for i := 0; i < 20; i++ {
		var todo = todos.NewTodo{
			Title:       faker.Lorem().Word(),
			Description: faker.Lorem().Sentence(5),
		}

		_, err := storage.CreateNewTodo(todo)

		if err != nil {
			log.Fatalf("Error seeding individual todos")
		}
	}

	if err != nil {
		log.Fatalf("error seeding todos: %v", err)
	}
	log.Print("Seeded todos")
}
