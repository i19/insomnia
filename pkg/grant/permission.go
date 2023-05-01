package grant

import (
	"sync"
	"time"

	"insomnia/internal/config"
	"insomnia/internal/service/project_access"
	"insomnia/internal/service/session"
)

type Grant struct {
	Token string

	uid      string
	isAdmin  bool
	ttlTo    int64
	projects map[string]struct{}
}

func (g *Grant) IsAdmin() bool {
	return g.isAdmin
}
func (g *Grant) HasProject(projectID string) bool {
	_, ok := g.projects[projectID]
	return ok
}

func (g *Grant) isFresh() bool {
	return g.ttlTo < time.Now().Unix()
}

var (
	grantCatch = make(map[string]*Grant)
	locker     sync.RWMutex
)

func GetBySession(token string) (*Grant, error) {
	if !config.Config.SessionCatch.On {
		return getByToken(token)
	}

	locker.RLock()
	g, ok := grantCatch[token]
	locker.RUnlock()
	if ok && g.isFresh() {
		return g, nil
	}

	g, err := getByToken(token)
	if err != nil {
		return nil, err
	}

	locker.Lock()
	g.ttlTo = time.Now().Unix() + config.Config.SessionCatch.TTL
	grantCatch[token] = g
	locker.Unlock()

	return g, nil
}

func getByToken(token string) (*Grant, error) {
	user, err := session.Get(token)
	if err != nil {
		return nil, err
	}

	projects, err := project_access.GetAllowedProjectIDsByUID(user.UID)
	if err != nil {
		return nil, err
	}

	g := Grant{
		Token:    token,
		uid:      user.UID,
		isAdmin:  user.IsAdmin,
		projects: make(map[string]struct{}, len(projects)),
	}
	for _, projectID := range projects {
		g.projects[projectID] = struct{}{}
	}

	return &g, nil
}
