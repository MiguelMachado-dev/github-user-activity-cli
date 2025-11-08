package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MiguelMachado-dev/github-user-activity-cli/internal/model"
)

type UserActivityRepository interface {
	GetUserActivity(username string) ([]model.UserActivity, error)
}

type userActivityRepositoryImpl struct {
	httpClient *http.Client
	baseURL    string
}

func NewUserActivityRepository() UserActivityRepository {
	return &userActivityRepositoryImpl{
		httpClient: &http.Client{Timeout: time.Second * 10},
		baseURL:    "https://api.github.com/users",
	}
}

func (r *userActivityRepositoryImpl) GetUserActivity(username string) ([]model.UserActivity, error) {
	url := fmt.Sprintf("%s/%s/events", r.baseURL, username)

	resp, err := r.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user activity: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user activity: received status code %d", resp.StatusCode)
	}

	var activities []model.UserActivity
	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return activities, nil
}
