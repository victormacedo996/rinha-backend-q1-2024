package dblock

import "context"

type DbLock interface {
	GetDbLock(context.Context) (string, error)
	LockDb(context.Context) error
	UnlockDb(context.Context) error
}
