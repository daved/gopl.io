package github

import "time"

// User ...
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// IssueRequest ...
type IssueRequest struct {
	Title string     `json:"title,omitempty"`
	Body  string     `json:"body,omitempty"`
	State IssueState `json:"state,omitempty"`
}

// IssueResponse ...
type IssueResponse struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// IssuesSearchResponse ...
type IssuesSearchResponse struct {
	TotalCount int `json:"total_count"`
	Items      []*IssueResponse
}
