// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: unfollow.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const unfollow = `-- name: Unfollow :exec
DELETE FROM feed_follows
USING feeds
WHERE feed_follows.feed_id = feeds.id
    AND feed_follows.user_id = $1 AND url = $2
`

type UnfollowParams struct {
	UserID uuid.UUID
	Url    string
}

func (q *Queries) Unfollow(ctx context.Context, arg UnfollowParams) error {
	_, err := q.db.ExecContext(ctx, unfollow, arg.UserID, arg.Url)
	return err
}
