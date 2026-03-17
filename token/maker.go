package token

import "time"

// Interface to manage tokens
type Maker interface {
	// Creates token for a username with specific duration
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
