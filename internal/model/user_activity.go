package model

import "time"

type Actor struct {
	ID          int64  `json:"id"`
	Login       string `json:"login"`
	DisplayName string `json:"display_login"`
	AvatarURL   string `json:"avatar_url"`
	URL         string `json:"url"`
}

type Repo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Commit struct {
	SHA     string `json:"sha"`
	Message string `json:"message"`
	URL     string `json:"url"`
}

type PullRequest struct {
	ID     int64  `json:"id"`
	Number int    `json:"number"`
	State  string `json:"state"`
	Title  string `json:"title"`
	URL    string `json:"html_url"`
}

type Payload struct {
	Ref    string `json:"ref,omitempty"`
	Before string `json:"before,omitempty"`
	Head   string `json:"head,omitempty"`
	PushID int64  `json:"push_id,omitempty"`
	RepoID int64  `json:"repository_id,omitempty"`
}

type UserActivity struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	Public    bool      `josn:"public"`
	CreatedAt time.Time `josn:"created_at"`
}
