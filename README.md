# GitHub User Activity CLI

A command-line tool that fetches and displays recent GitHub user activity in a clean, human-readable format.

## Overview

This tool aggregates GitHub user events from the public GitHub API and presents them in a summarized format, making it easy to quickly see what a user has been up to across their repositories.

Read more about the project [here]("https://roadmap.sh/projects/github-user-activity")

## Features

- **Activity Aggregation**: Groups similar events to reduce noise and provide cleaner output
- **Multiple Event Types**: Supports PushEvent, PullRequestEvent, IssuesEvent, and ForkEvent
- **Simple Interface**: Easy-to-use command-line interface
- **Rate Limit Aware**: Built-in timeout and error handling for API requests
- **Clean Architecture**: Well-structured codebase with separation of concerns

## Prerequisites

- Go 1.25.4 or later
- Internet connection to access GitHub's API

## Installation

### From Source

1. Clone the repository:
```bash
git clone https://github.com/MiguelMachado-dev/github-user-activity-cli.git
cd github-user-activity-cli
```

2. Build the application:
```bash
go build -o github-activity ./cmd/main.go
```

3. (Optional) Move the binary to your PATH:
```bash
sudo mv github-activity /usr/local/bin/
```

## Usage

### Basic Usage

```bash
# Using the built binary
./github-activity <github-username>

# Example
./github-activity torvalds
```

### Using Go Run

```bash
go run ./cmd/main.go <github-username>
```

### Help

If no username is provided, the tool will display usage instructions:
```bash
./github-activity
# Output: Please, provide a Github Username. Usage: github-activity <github-username>
```

## Architecture

The project follows a clean architecture pattern with three distinct layers:

### 1. Presentation Layer (`cmd/`)
- **Main.go**: Entry point that handles CLI arguments and orchestrates the application flow
- Manages dependency injection between layers

### 2. Business Logic Layer (`internal/service/`)
- **UserActivityService**: Contains the core logic for processing and aggregating user activity data
- Implements intelligent grouping of events by type and repository
- Converts raw GitHub events into human-readable summaries

### 3. Data Access Layer (`internal/repository/`)
- **UserActivityRepository**: Handles communication with the GitHub API
- Implements HTTP client with timeout configuration
- Manages API response parsing and error handling

### 4. Models (`internal/model/`)
- **UserActivity**: Data structures representing GitHub API responses
- Defines all supported GitHub event types and their properties

## Project Structure

```
github-user-activity-cli/
├── cmd/
│   └── main.go                           # Application entry point
├── internal/
│   ├── model/
│   │   └── user_activity.go              # Data models and structs
│   ├── repository/
│   │   └── user_activity_repository.go   # GitHub API client
│   └── service/
│       └── user_activity_service.go      # Business logic
├── go.mod                                # Go module definition
└── README.md                             # This file
```

## Example Output

### Active User Example
```bash
./github-activity octocat
```

Output:
```
Recent activity for user octocat:
- Pushed to octocat/Hello-World (3 events)
- Pull request activity in octocat/Spoon-Knife
- Forked repository octocat/Octocat-Repo
- Issue activity in octocat/git-consortium (2 events)
```

### Inactive or New User Example
```bash
./github-activity newuser123
```

Output:
```
No recent activity found for user newuser123
```

## Error Handling

The tool handles various error scenarios gracefully:

### Network Errors
```bash
Error fetching user activity: failed to fetch user activity: Get "https://api.github.com/users/invalid-username/events": context deadline exceeded
```

### Invalid Username
```bash
Error fetching user activity: failed to fetch user activity: received status code 404
```

### API Rate Limiting
```bash
Error fetching user activity: failed to fetch user activity: received status code 403
```

## Supported Event Types

The tool currently processes and aggregates the following GitHub event types:

| Event Type | Description | Example Output |
|------------|-------------|----------------|
| **PushEvent** | Code pushes to repositories | `Pushed to username/repo (3 events)` |
| **PullRequestEvent** | Pull request activities (opened, closed, merged) | `Pull request activity in username/repo` |
| **IssuesEvent** | Issue activities (opened, closed, labeled) | `Issue activity in username/repo (2 events)` |
| **ForkEvent** | Repository forking | `Forked repository username/repo` |

### Event Aggregation

Events are intelligently grouped by:
- Event type (PushEvent, PullRequestEvent, etc.)
- Repository name

This reduces duplicate entries and provides a cleaner summary of user activity.

## API Details

- **Base URL**: `https://api.github.com/users`
- **Endpoint**: `/{username}/events`
- **Rate Limiting**: Subject to GitHub's unauthenticated API limits (60 requests per hour)
- **Timeout**: 10 seconds per request
- **Authentication**: None required (uses public endpoints)

## Development

### Build Commands
```bash
# Build the application
go build -o github-activity ./cmd/main.go

# Format code
go fmt ./...

# Vet code for potential issues
go vet ./...

# Update dependencies
go mod tidy

# Run tests (when available)
go test ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Disclaimer

This tool uses GitHub's public API and is subject to GitHub's terms of service and rate limiting policies. Please use responsibly and avoid making excessive requests.