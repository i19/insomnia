package session

import (
	"strings"
	"sync"
)

var (
	address string
	doOnce  sync.Once
)

func Init(addr string) {
	doOnce.Do(
		func() {
			// todo some check
			address = addr
		},
	)

}

type user struct {
	UID     string
	IsAdmin bool
}

func Get(token string) (*user, error) {
	if strings.HasPrefix(token, "admin") {
		return &user{
			UID:     "admin_user",
			IsAdmin: true,
		}, nil
	} else if strings.HasPrefix(token, "grant") {
		return &user{
			UID:     "grant_user",
			IsAdmin: false,
		}, nil
	} else {
		return &user{
			UID:     "guest",
			IsAdmin: false,
		}, nil
	}
}
