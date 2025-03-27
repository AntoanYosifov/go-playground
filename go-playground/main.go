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
	}

	fmt.Println("Job Application:")
	fmt.Printf("Company: %s\nTitle: %s\nStatus: %s\n", job.Company, job.Title, job.Status)
}
