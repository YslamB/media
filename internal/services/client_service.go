package services

import (
	"context"
	"media/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientService struct {
	repo *repositories.ClientRepository
}

func NewClientService(db *pgxpool.Pool) *ClientService {
	return &ClientService{repositories.NewClientRepository(db)}

}

func (us *ClientService) GetUsers(ctx context.Context, id int) (int, error) {

	// otp := response.OTP{}
	// userPhone := us.repo.CheckUserExist(ctx, user.Phone)

	// if userPhone != "" {
	// 	return otp
	// }

	return 1, nil

}
