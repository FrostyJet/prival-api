package token

import "time"

type Token interface {
	Create(userID string, duration time.Duration) (string, error)
	Verify(token string) (*Payload, error)
}
