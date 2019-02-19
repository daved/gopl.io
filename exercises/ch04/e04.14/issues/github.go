package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/daved/gopl.io/exercises/ch04/e04.14/github"
)

type ownerRepo struct {
	owner, repo string
}

type issues struct {
	data []*github.IssueResponse
	last time.Time
}

type gitHub struct {
	user, token string
	cache       map[ownerRepo]*issues
	mu          sync.Mutex
	expiry      time.Duration
}

func newGitHub(user, token string, expiry time.Duration) *gitHub {
	return &gitHub{
		user:   user,
		token:  token,
		cache:  make(map[ownerRepo]*issues),
		expiry: expiry,
	}
}

func (gh *gitHub) issues(owner, repo string) (*issues, error) {
	or := ownerRepo{owner, repo}

	is, ok := gh.access(or)
	if ok && time.Now().Sub(is.last) < gh.expiry {
		return is, nil
	}

	m, err := github.NewAccessMgmt(gh.user, gh.token, owner, repo)
	if err != nil {
		return nil, fmt.Errorf("issues: create access management: %s", err)
	}

	gis, err := m.ReadIssues()
	if err != nil {
		return nil, fmt.Errorf("issues: cannot read issues: %s", err)
	}

	is = &issues{
		data: gis,
		last: time.Now(),
	}

	gh.insert(or, is)

	return is, nil
}

func (gh *gitHub) insert(or ownerRepo, is *issues) {
	gh.mu.Lock()
	defer gh.mu.Unlock()

	gh.cache[or] = is
}

func (gh *gitHub) access(or ownerRepo) (*issues, bool) {
	gh.mu.Lock()
	defer gh.mu.Unlock()

	v, ok := gh.cache[or]

	return v, ok
}
