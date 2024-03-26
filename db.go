package goagithub

import (
	usersv1 "github.com/danielmichaels/goa-github/gen/users_v1"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	m map[string]*usersv1.Users
}

var preSeededUsers = map[string]*usersv1.Users{
	"rs": {
		Username:  "rs",
		Name:      "Olivier Poitrey",
		Followers: 4468,
		Repos:     116,
	},
	"danielmichaels": {
		Username:  "danielmichaels",
		Name:      "Daniel Michaels",
		Followers: 31,
		Repos:     68,
	},
}

func NewSafeMap() *SafeMap {
	sm := new(SafeMap)
	sm.m = make(map[string]*usersv1.Users)
	if len(sm.m) == 0 {
		sm.m = preSeededUsers
	}
	return sm
}

func (sm *SafeMap) List() []*usersv1.Users {
	sm.RLock()
	defer sm.RUnlock()
	users := make([]*usersv1.Users, 0, len(sm.m))
	for _, v := range sm.m {
		users = append(users, v)
	}
	return users
}

func (sm *SafeMap) Read(key string) (*usersv1.Users, bool) {
	sm.RLock()
	defer sm.RUnlock()
	item, ok := sm.m[key]
	return item, ok
}

func (sm *SafeMap) Write(key string, item *usersv1.Users) {
	sm.Lock()
	defer sm.Unlock()
	sm.m[key] = item
}

func (sm *SafeMap) Delete(key string) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.m, key)
}
