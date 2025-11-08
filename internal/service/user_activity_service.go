package service

import (
	"fmt"

	"github.com/MiguelMachado-dev/github-user-activity-cli/internal/repository"
)

type UserActivityService interface {
	GetUserActivity(username string) ([]string, error)
}

// userActivityServiceImpl implements UserActivityService
type userActivityServiceImpl struct {
	repo repository.UserActivityRepository
}

func NewUserActivityService(repo repository.UserActivityRepository) UserActivityService {
	return &userActivityServiceImpl{repo: repo}
}

func (s *userActivityServiceImpl) GetUserActivity(username string) ([]string, error) {
	activities, err := s.repo.GetUserActivity(username)
	if err != nil {
		return nil, err
	}

	var summaries []string

	for _, activity := range activities {
		repoName := activity.Repo.Name

		switch activity.Type {
		case "PushEvent":
			count := len(activity.Payload.Commits)
			summaries = append(summaries, fmt.Sprintf("Pushed %d commit(s) to %s", count, repoName))

		case "PullRequestEvent":
			action := activity.Payload.Action
			summaries = append(summaries, fmt.Sprintf("%s a pull request in %s", action, repoName))

		case "IssuesEvent":
			action := activity.Payload.Action
			summaries = append(summaries, fmt.Sprintf("%s an issue in %s", action, repoName))

		case "ForkEvent":
			summaries = append(summaries, fmt.Sprintf("Forked repository %s", repoName))

		default:
			// Ignore other event types
		}
	}

	return summaries, nil
}
