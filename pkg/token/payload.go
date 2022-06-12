package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewPayload(userID string, duration time.Duration) (*Payload, error) {
	payload := &Payload{}

	uniqueId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	payload.ID = uniqueId.String()
	payload.UserID = userID
	payload.IssuedAt = now
	payload.ExpiresAt = now.Add(duration)

	return payload, nil
}

func (p *Payload) Valid() error {
	if p.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("token has expired")
	}

	return nil
}
