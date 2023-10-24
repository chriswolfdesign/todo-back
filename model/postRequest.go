package model

type PostRequest struct {
	Text string `json:"text"`
	Completed bool `json:"completed"`
}