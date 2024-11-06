package memory

import (
	"context"
	"sync"

	"github.com/onee-only/typewriter/backend/internal/object"
	"github.com/onee-only/typewriter/backend/internal/storage"
)

type VoteStorage struct {
	status [4]uint
	lock   sync.RWMutex
}

var _ storage.VoteStorage = (*VoteStorage)(nil)

func NewVoteStorage() *VoteStorage {
	return &VoteStorage{
		status: [4]uint{},
		lock:   sync.RWMutex{},
	}
}

func (s *VoteStorage) Reset(_ context.Context) error {
	s.lock.Lock()

	s.status = [4]uint{}
	s.lock.Unlock()

	return nil
}

func (s *VoteStorage) Status(_ context.Context) (votes [4]uint, err error) {
	s.lock.RLock()
	stat := s.status
	s.lock.RUnlock()

	return stat, nil
}

func (s *VoteStorage) Vote(_ context.Context, dir object.Direction) (uint, error) {
	s.lock.Lock()

	incr := s.status[dir] + 1
	s.status[dir] = incr
	s.lock.Unlock()

	return incr, nil
}
