package dbModel

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) getBooks() []*Book {
	return nil
}

func (s *sqliteHandler) getBook(id int) *Book {
	return nil
}

func (s *sqliteHandler) addBook(book Book) *Book {
	return nil
}

func (s *sqliteHandler) removeBook(id int) bool {
	return false
}

// func (s *sqliteHandler) completeBook(id int, complete bool) bool {
// 	return false
// }

func (s *sqliteHandler) close() {
	s.db.Close()
}

func newSqliteHandler() dbHandler {
	os.Remove("./test.db")
	database, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS books (
			id        INTEGER  PRIMARY KEY AUTOINCREMENT,
			Isbn      TEXT,
			Title BOOLEAN,
			createdAt DATETIME
		)`)
	statement.Exec()
	return &sqliteHandler{db: database}
}
