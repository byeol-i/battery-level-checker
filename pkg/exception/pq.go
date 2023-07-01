package exception

import (
	"github.com/lib/pq"
)

func isDuplicateKeyError(err error, constraint string) bool {
	pqErr, ok := err.(*pq.Error)
	if !ok {
		return false
	}

	return pqErr.Constraint == constraint
}

func CastingError(err error) string {
	switch {
	case isDuplicateKeyError(err, "User_pkey"):
		return "Already exist user"
	default:
		return err.Error()
	}
}