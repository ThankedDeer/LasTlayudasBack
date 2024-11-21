package token

import "time"

type Make interface {
	CreateToken(email string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
