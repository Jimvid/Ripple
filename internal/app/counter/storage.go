package counter

import "github.com/jmoiron/sqlx"

var schema = `
CREATE TABLE count(
    id int NOT NULL AUTO_INCREMENT,
    count int NOT NULL,
)`

type CounterStorage struct {
	Conn *sqlx.DB
}

func NewCounterStorage(conn *sqlx.DB) *CounterStorage {
	return &CounterStorage{Conn: conn}
}

type Count struct {
	id    int
	count int
}

func (s *CounterStorage) CreateCounter() int {
	var count int
	err := s.Conn.Get(&count, "INSERT INTO counters (count) VALUES (?)", 0)
	if err != nil {
		panic(err)
	}
	return count
}

func (s *CounterStorage) GetCount(id int) int {
	var count int
	err := s.Conn.Get(&count, "SELECT count FROM counts WHERE id=?", id)
	if err != nil {
		panic(err)
	}
	return count
}

func (s *CounterStorage) IncrementCount(id int) {
	_, err := s.Conn.Exec("UPDATE counts SET count = count + 1 WHERE id=?", id)
	if err != nil {
		panic(err)
	}
}

func (s *CounterStorage) DecrementCount(count int) {
	_, err := s.Conn.Exec("UPDATE counts SET count = count - 1 WHERE id = 1")
	if err != nil {
		panic(err)
	}
}
