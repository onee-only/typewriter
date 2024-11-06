package memory

import (
	"context"
	"sync"
	"testing"

	"github.com/onee-only/typewriter/backend/internal/object"
	"github.com/stretchr/testify/suite"
)

type VoteStorageTestSuite struct {
	suite.Suite
	storage *VoteStorage
}

func TestVoteStorageTestSuite(t *testing.T) {
	suite.Run(t, new(VoteStorageTestSuite))
}

func (s *VoteStorageTestSuite) SetupTest() {
	s.storage = NewVoteStorage()
}

func (s *VoteStorageTestSuite) TestStatus() {
	ctx := context.Background()

	votes, err := s.storage.Status(ctx)

	s.Require().NoError(err)
	s.Equal(s.storage.status, votes)
	s.Equal([4]uint{0, 0, 0, 0}, votes)
}

func (s *VoteStorageTestSuite) TestVote() {
	ctx := context.Background()

	err := s.storage.Vote(ctx, object.DirectionUp)
	s.Require().NoError(err)

	votes, err := s.storage.Status(ctx)
	s.Require().NoError(err)
	s.Equal(uint(1), votes[object.DirectionUp])
}

func (s *VoteStorageTestSuite) TestReset() {
	ctx := context.Background()

	err := s.storage.Vote(ctx, object.DirectionUp)
	s.Require().NoError(err)

	err = s.storage.Reset(ctx)
	s.Require().NoError(err)

	votes, err := s.storage.Status(ctx)
	s.Require().NoError(err)
	s.Equal([4]uint{0, 0, 0, 0}, votes)
}

func (s *VoteStorageTestSuite) TestConsistency() {
	ctx := context.Background()

	dir := object.DirectionUp
	workers := uint(3)
	numIncrs := uint(1000)

	incr := func() {
		err := s.storage.Vote(ctx, dir)
		s.Require().NoError(err)
	}

	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for range numIncrs {
				incr()
			}
		}()
	}

	wg.Wait()

	votes, err := s.storage.Status(ctx)
	s.Require().NoError(err)
	s.Equal(workers*numIncrs, votes[dir])
}
