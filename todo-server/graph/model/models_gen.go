// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Title  string `json:"title"`
	UserID int    `json:"userId"`
}

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
	UserID int    `json:"userId"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}