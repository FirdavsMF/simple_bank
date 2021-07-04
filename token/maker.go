package token

import (
	"time"
)

// Maker - это интерфейс для управления токенами
type Maker interface {
	// CreateToken создает новый токен для определенного имени пользователя и продолжительности
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken проверяет, действителен ли токен
	VerifyToken(token string) (*Payload, error)
}
