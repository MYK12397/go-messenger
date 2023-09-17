package ports

import "github.com/MYK12397/go-messenger/internal/core/domain"

type Servicer interface {
	SaveMessage(message domain.Message) error
	ReadMessage(id string) (*domain.Message, error)
	ReadMessages() ([]*domain.Message, error)
	DeleteMessage(id string) error
}

type Repository interface {
	SaveMessage(message domain.Message) error
	ReadMessage(id string) (*domain.Message, error)
	ReadMessages() ([]*domain.Message, error)
	DeleteMessage(id string) error
}
