package state

import "sync"

// Tracks the database state, to be used for caching.
// will be used by multiple goroutines so must use Mutex
// if any db curd operation is performed, `isChanged` updates to true
// when a new caching operation is performed, that means that redis is updated so isChanged switches back to false.

type DatabaseCachingState struct {
	isChanged bool
	sync.Mutex
}

var DbCacheState = NewCachingState()

func NewCachingState() *DatabaseCachingState {
	return &DatabaseCachingState{isChanged: true}
}

func (state *DatabaseCachingState) SetIsChanged(value bool) {
	state.Lock()
	defer state.Unlock()
	state.isChanged = value
}

func (state *DatabaseCachingState) WasDBChanged() bool {
	state.Lock()
	defer state.Unlock()
	return state.isChanged
}