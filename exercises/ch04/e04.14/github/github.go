package github

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// IssueState ...
type IssueState string

// Issue States ...
const (
	Open   IssueState = "open"
	Closed IssueState = "closed"
)

// LockReason ...
type LockReason string

// Lock Reasons ...
const (
	OffTopic  LockReason = "off-topic"
	TooHeated LockReason = "too heated"
	Resolved  LockReason = "resolved"
	Spam      LockReason = "spam"
)

// AccessMgmt ...
type AccessMgmt struct {
	urlPrfx  string
	repoSet  bool
	repoSrch string
	crudPath string
	srchPath string
	lockSufx string
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
		lockSufx: "/lock",
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
	req := issueRequest{
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
func (m *AccessMgmt) ReadIssue(num int) (*IssueResponse, error) {
	if !m.repoSet {
		return nil, ErrRepoNotSet
	}

	method := http.MethodGet
	url := fmt.Sprintf("%s%s/%d", m.urlPrfx, m.crudPath, num)

	var res IssueResponse
	if err := sendDecode(method, url, m.hdrs, nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ReadIssues ...
func (m *AccessMgmt) ReadIssues() ([]*IssueResponse, error) {
	if !m.repoSet {
		return nil, ErrRepoNotSet
	}

	method := http.MethodGet
	url := fmt.Sprintf("%s%s?state=all", m.urlPrfx, m.crudPath)

	var res []*IssueResponse
	if err := sendDecode(method, url, m.hdrs, nil, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateIssue ...
func (m *AccessMgmt) UpdateIssue(num int, title, body string, state IssueState) (*IssueResponse, error) {
	if !m.repoSet {
		return nil, ErrRepoNotSet
	}

	method := http.MethodPatch
	url := fmt.Sprintf("%s%s/%d", m.urlPrfx, m.crudPath, num)
	req := issueRequest{
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

// LockIssue ...
func (m *AccessMgmt) LockIssue(num int, rsn LockReason) error {
	if !m.repoSet {
		return ErrRepoNotSet
	}

	method := http.MethodPut
	url := fmt.Sprintf("%s%s/%d%s", m.urlPrfx, m.crudPath, num, m.lockSufx)
	req := issueDeleteRequest{
		Locked: true,
		Reason: rsn,
	}

	return sendDecode(method, url, m.hdrs, &req, nil)
}

// UnlockIssue ...
func (m *AccessMgmt) UnlockIssue(num int) error {
	if !m.repoSet {
		return ErrRepoNotSet
	}

	method := http.MethodDelete
	url := fmt.Sprintf("%s%s/%d%s", m.urlPrfx, m.crudPath, num, m.lockSufx)

	return sendDecode(method, url, m.hdrs, nil, nil)
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
