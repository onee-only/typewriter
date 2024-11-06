package storage

import (
	"context"

	"github.com/onee-only/typewriter/backend/internal/object"
)

// VoteStorage keeps the vote counts of cursor's next direction.
// Operations should be atomic since multiple concurrent entities can access simultaneously.
type VoteStorage interface {
	// Vote increases counter of given direction by 1.
	// It returns count of direction after the vote.
	Vote(ctx context.Context, dir object.Direction) (uint, error)
	// Status returns current vote count of all directions.
	// The order of directions will be order of [object.Direction].
	Status(ctx context.Context) (votes [4]uint, err error)
	// Reset resets vote count of directions.
	Reset(ctx context.Context) error
}
