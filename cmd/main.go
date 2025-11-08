package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MiguelMachado-dev/github-user-activity-cli/internal/repository"
	"github.com/MiguelMachado-dev/github-user-activity-cli/internal/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please, provide a Github Username. Usage: github-activity <github-username>")
		os.Exit(1)
	}

	username := os.Args[1]

	// Initialize repository and service
	repo := repository.NewUserActivityRepository()
	service := service.NewUserActivityService(repo)

	activities, err := service.GetUserActivity(username)
	if err != nil {
		log.Fatalf("Error fetching user activity: %v", err)
	}

	if len(activities) == 0 {
		fmt.Printf("No recent activity found for user %s\n", username)
		return
	}

	fmt.Printf("Recent activity for user %s:\n", username)
	for _, activity := range activities {
		fmt.Println("- " + activity)
	}
}
