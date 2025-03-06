package contracts

import (
	"context"

	"github.com/ahargunyllib/hackathon-fiber-starter/domain/dto"
	"github.com/ahargunyllib/hackathon-fiber-starter/domain/entity"
	"github.com/google/uuid"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	RegisterUser(ctx context.Context, user entity.User) (uuid.UUID, error)
}

type AuthService interface {
	RegisterUser(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
	LoginUser(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
}
