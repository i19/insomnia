package project_access

import "sync"

var (
	address string
	doOnce  sync.Once
)

func Init(addr string) {
	doOnce.Do(
		func() {
			// todo some check
			address = addr
		})
}

func GetAllowedProjectIDsByUID(userID string) ([]string, error) {
	if userID == "guest" {
		return []string{}, nil
	}
	return []string{"project_1"}, nil
}
