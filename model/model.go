package model

type TODOItem struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

type UpdateRequest struct {
	ID        int    `json:"id"`
	Body      string `json:"body,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

type CreateRequest struct {
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

type DeleteRequest struct {
	ID int `json:"id"`
}