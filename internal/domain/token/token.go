package token

import "time"

type Token struct {
	ID        int
	UserID    int
	Token     string
	ExpiresAt time.Time
	Type      string // 'auth', 'refresh', 'recovery'
}
