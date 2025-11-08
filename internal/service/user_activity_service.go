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

	// Aggregate events by type and repository to avoid duplicate summaries.
	type aggregateKey struct {
		eventType string
		repoName  string
	}

	type aggregateValue struct {
		occurrences int
	}

	aggregated := make(map[aggregateKey]aggregateValue)
	var order []aggregateKey

	for _, activity := range activities {
		repoName := activity.Repo.Name

		switch activity.Type {
		case "PushEvent":
			key := aggregateKey{eventType: activity.Type, repoName: repoName}
			value, exists := aggregated[key]
			if !exists {
				order = append(order, key)
			}
			value.occurrences++
			aggregated[key] = value

		case "PullRequestEvent":
			key := aggregateKey{eventType: activity.Type, repoName: repoName}
			value, exists := aggregated[key]
			if !exists {
				order = append(order, key)
			}
			value.occurrences++
			aggregated[key] = value

		case "IssuesEvent":
			key := aggregateKey{eventType: activity.Type, repoName: repoName}
			value, exists := aggregated[key]
			if !exists {
				order = append(order, key)
			}
			value.occurrences++
			aggregated[key] = value

		case "ForkEvent":
			key := aggregateKey{eventType: activity.Type, repoName: repoName}
			value, exists := aggregated[key]
			if !exists {
				order = append(order, key)
			}
			value.occurrences++
			aggregated[key] = value

		default:
			// Ignore other event types
		}
	}

	for _, key := range order {
		value := aggregated[key]
		switch key.eventType {
		case "PushEvent":
			if value.occurrences > 1 {
				summaries = append(summaries, fmt.Sprintf("Pushed to %s (%d events)", key.repoName, value.occurrences))
			} else {
				summaries = append(summaries, fmt.Sprintf("Pushed to %s", key.repoName))
			}
		case "PullRequestEvent":
			if value.occurrences > 1 {
				summaries = append(summaries, fmt.Sprintf("Pull request activity in %s (%d events)", key.repoName, value.occurrences))
			} else {
				summaries = append(summaries, fmt.Sprintf("Pull request activity in %s", key.repoName))
			}
		case "IssuesEvent":
			if value.occurrences > 1 {
				summaries = append(summaries, fmt.Sprintf("Issue activity in %s (%d events)", key.repoName, value.occurrences))
			} else {
				summaries = append(summaries, fmt.Sprintf("Issue activity in %s", key.repoName))
			}
		case "ForkEvent":
			if value.occurrences > 1 {
				summaries = append(summaries, fmt.Sprintf("Forked repository %s (%d times)", key.repoName, value.occurrences))
			} else {
				summaries = append(summaries, fmt.Sprintf("Forked repository %s", key.repoName))
			}
		}
	}

	return summaries, nil
}
