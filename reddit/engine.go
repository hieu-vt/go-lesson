package reddit

import (
	"sync"
)

type RedditEngine interface {
	Save(key string, data interface{})
	Get(key string) interface{}
}

type redditEngine struct {
	storage map[string]interface{}
	locker  *sync.RWMutex
}

func NewRedditEngine() *redditEngine {
	return &redditEngine{
		storage: make(map[string]interface{}),
		locker:  new(sync.RWMutex),
	}
}

func (rEngine *redditEngine) Save(key string, data interface{}) {
	rEngine.locker.Lock()
	rEngine.storage[key] = data
	rEngine.locker.Unlock()
}

func (rEngine *redditEngine) Get(key string) interface{} {
	rEngine.locker.RLock()
	defer rEngine.locker.RUnlock()

	return rEngine.storage[key]
}
