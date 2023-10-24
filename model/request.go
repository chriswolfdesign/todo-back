package model

type Request struct {
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

