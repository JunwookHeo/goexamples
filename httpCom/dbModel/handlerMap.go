package dbModel

type handlerMap struct {
	bookMap map[int]*Book
}

func (m *handlerMap) getBooks() []*Book {
	list := []*Book{}
	for _, v := range m.bookMap {
		list = append(list, v)
	}
	return list
}

func (m *handlerMap) getBook(id int) *Book {
	book, ok := m.bookMap[id]
	if ok == false {
		return nil
	}
	return book
}

func (m *handlerMap) addBook(book Book) *Book {
	book.ID = len(m.bookMap) + 1
	m.bookMap[book.ID] = &book
	return &book
}

func (m *handlerMap) removeBook(id int) bool {
	if _, ok := m.bookMap[id]; ok {
		delete(m.bookMap, id)
		return true
	}
	return false
}

// func (m *handlerMap) completeBook(id int, complete bool) bool {
// 	if book, ok := m.bookMap[id]; ok {
// 		book.Completed = complete
// 		return true
// 	}
// 	return false
// }

func newHandlerMap() dbHandler {
	m := &handlerMap{}
	m.bookMap = make(map[int]*Book)
	return m
}
