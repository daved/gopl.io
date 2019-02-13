package github

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// IssueState ...
type IssueState string

// State{status} ...
const (
	StateOpen   IssueState = "open"
	StateClosed IssueState = "closed"
)

// AccessMgmt ...
type AccessMgmt struct {
	urlPrfx  string
	repoSet  bool
	repoSrch string
	crudPath string
	srchPath string
	hdrs     headers
}

// NewAccessMgmt ...
func NewAccessMgmt(user, token, owner, repo string) (*AccessMgmt, error) {
	verHdr := &header{"Accept", "application/vnd.github.v3.text-match+json"}
	authHdr, err := newHeaderAuth("https://api.github.com/user", user, token)
	if err != nil {
		return nil, err
	}

	m := AccessMgmt{
		urlPrfx:  "https://api.github.com",
		repoSet:  owner != "" && repo != "",
		repoSrch: "repo:" + owner + "/" + repo,
		crudPath: "/repos/" + owner + "/" + repo + "/issues",
		srchPath: "/search/issues",
		hdrs: headers{
			verHdr,
			authHdr,
		},
	}

	return &m, nil
}

// CreateIssue ...
func (m *AccessMgmt) CreateIssue(title, body string) (*IssueResponse, error) {
	if !m.repoSet {
		return nil, ErrRepoNotSet
	}

	method := http.MethodPost
	url := m.urlPrfx + m.crudPath
	req := IssueRequest{
		Title: title,
		Body:  body,
	}

	var res IssueResponse
	if err := sendDecode(method, url, m.hdrs, &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ReadIssue ...
func (m *AccessMgmt) ReadIssue() error {
	return nil
}

// UpdateIssue ...
func (m *AccessMgmt) UpdateIssue(num int, title, body string, state IssueState) (*IssueResponse, error) {
	if !m.repoSet {
		return nil, ErrRepoNotSet
	}

	method := http.MethodPatch
	url := m.urlPrfx + m.crudPath + fmt.Sprintf("/%d", num)
	req := IssueRequest{
		Title: title,
		Body:  body,
		State: state,
	}

	var res IssueResponse
	if err := sendDecode(method, url, m.hdrs, &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteIssue ...
func (m *AccessMgmt) DeleteIssue() error {
	return nil
}

// SearchIssues queries the GitHub issue tracker.
func (m *AccessMgmt) SearchIssues(terms []string) (*IssuesSearchResponse, error) {
	if m.repoSet {
		terms = append(terms, m.repoSrch)
	}

	method := http.MethodGet
	url := m.urlPrfx + m.srchPath + "?q=" + url.QueryEscape(strings.Join(terms, " "))

	var res IssuesSearchResponse
	if err := sendDecode(method, url, m.hdrs, nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
