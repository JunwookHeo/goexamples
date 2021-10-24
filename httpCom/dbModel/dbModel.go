package dbModel

type dbHandler interface {
	getBooks() []*Book
	getBook(id int) *Book
	addBook(book Book) *Book
	removeBook(id int) bool
	// completeBook(id int, complete bool) bool
}

// Book struct (Model)
type Book struct {
	ID     int     `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var handler dbHandler

func init() {
	//handler = newHandlerMap()
	handler = newSqliteHandler()
}

func GetBooks() []*Book {
	return handler.getBooks()
}

func GetBook(id int) *Book {
	return handler.getBook(id)
}

func AddBook(book Book) *Book {
	return handler.addBook(book)
}

func RemoveBook(id int) bool {
	return handler.removeBook(id)
}

// func CompleteBook(id int, complete bool) bool {
// 	return handler.completeBook(id, complete)
// }
