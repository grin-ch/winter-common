package rand_util

import "github.com/google/uuid"

// MustUUIDv4
func MustUUIDv4() string {
	return uuid.New().String()
}

// UUIDv4
func UUIDv4() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil.String(), err
	}
	return u.String(), nil
}
