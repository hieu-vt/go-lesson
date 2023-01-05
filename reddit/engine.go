package reddit

import (
	"sync"
)

type RedditEngine interface {
	Save(key interface{}, data interface{}) bool
	Remove(key interface{}) bool
	Get(key interface{}) interface{}
}

type redditEngine struct {
	storage map[interface{}]interface{}
	locker  *sync.RWMutex
}

func NewRedditEngine() *redditEngine {
	return &redditEngine{
		storage: make(map[interface{}]interface{}),
		locker:  new(sync.RWMutex),
	}
}

func (rEngine *redditEngine) Save(key interface{}, data interface{}) bool {
	if data != nil && key != nil {
		rEngine.locker.Lock()
		rEngine.storage[key] = data
		rEngine.locker.Unlock()
		return true
	} else {
		return false
	}
}

func (rEngine *redditEngine) Remove(key interface{}) bool {
	rEngine.locker.Lock()
	defer rEngine.locker.Unlock()
	if _, ok := rEngine.storage[key]; ok {
		rEngine.storage[key] = nil
		return false
	} else {
		return false
	}
}

func (rEngine *redditEngine) Get(key interface{}) interface{} {
	rEngine.locker.RLock()
	defer rEngine.locker.RUnlock()

	return rEngine.storage[key]
}
