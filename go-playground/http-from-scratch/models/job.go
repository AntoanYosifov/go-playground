package models

type Job struct {
	Title    string `json:"title"`
	Company  string `json:"company"`
	Location string `json:"location"`
}
