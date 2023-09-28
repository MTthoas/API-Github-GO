package model

type Repository struct {
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	HTMLURL         string `json:"html_url"`
	Description     string `json:"description"`
	Language        string `json:"language"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	StargazersCount int    `json:"stargazers_count"`
	ForksCount      int    `json:"forks_count"`
	WatchersCount   int    `json:"watchers_count"`
	OpenIssuesCount int    `json:"open_issues_count"`
	DefaultBranch   string `json:"default_branch"`
	Owner           struct {
		Login     string `json:"login"`
		AvatarURL string `json:"avatar_url"`
		HTMLURL   string `json:"html_url"`
	} `json:"owner"`
}
