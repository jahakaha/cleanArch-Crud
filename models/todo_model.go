package models

type TodoList struct {
	ID          int
	Title       string
	Description string
}

type UserList struct {
	ID     int
	UserID int
	ListID int
}

type TodoItem struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}
