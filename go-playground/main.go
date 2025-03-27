package main

import (
	"fmt"
	"go-playground/models"
)

func main() {
	fmt.Println("Hello, Go playground")
	job := models.Job{
		ID: 1,
		Company: "antdevrealm",
		Title: "Go Developer",
		Status: "Applied",
		Feedback: "TestFeedback",
	}

	fmt.Println("Job Application:")
	fmt.Printf("Company: %s\nTitle: %s\nStatus: %s\nFeedback: %s\n", job.Company, job.Title, job.Status, job.Feedback)
	fmt.Println("Something to commit")
	fmt.Println("Another test")
}
