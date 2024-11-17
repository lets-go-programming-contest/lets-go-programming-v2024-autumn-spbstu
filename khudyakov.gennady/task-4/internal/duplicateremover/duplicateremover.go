package duplicateremover

import (
	"sync"

	"github.com/KRYST4L614/task-4/internal/user"
)

type DuplicateRemover struct {
	Emails map[string]user.User
	Mu     sync.Mutex
}

func NewDuplicateRemover() *DuplicateRemover {
	return &DuplicateRemover{Emails: map[string]user.User{}, Mu: sync.Mutex{}}
}

func (dr *DuplicateRemover) GetUnique(in chan user.User) []user.User {
	result := []user.User{}
	var wg sync.WaitGroup
	for user := range in {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			dr.makeUnique(user, &result)
		}()
	}
	wg.Wait()
	return result
}

func (dr *DuplicateRemover) GetUniqueSafe(in chan user.User) []user.User {
	result := []user.User{}
	var wg sync.WaitGroup
	for user := range in {
		wg.Add(1)
		go func() {
			defer func() {
				dr.Mu.Unlock()
				wg.Done()
			}()
			dr.Mu.Lock()
			dr.makeUnique(user, &result)
		}()
	}
	wg.Wait()
	return result
}

func (dr *DuplicateRemover) makeUnique(user user.User, result *[]user.User) {
	_, contain := dr.Emails[user.Email]
	if !contain {
		*result = append(*result, user)
		dr.Emails[user.Email] = user
	}
}
