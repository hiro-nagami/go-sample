// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Title  string `json:"title"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
	UserID *User  `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
