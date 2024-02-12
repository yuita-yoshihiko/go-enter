// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

type Like struct {
	ID   string `json:"id"`
	User *User  `json:"user"`
	Book *Book  `json:"book"`
}

type NewBook struct {
	Title string `json:"title"`
	Genre string `json:"genre"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
