package user

import (
	"context"
	"fmt"
	"pet-project/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: pool,
	}
}

func (u UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	rows, err := u.db.Query(ctx, `
		SELECT *
			FROM users
			WHERE id = $1
	`, id)

	if err != nil {
		return nil, fmt.Errorf("user.GetByID error: %w", err)
	}

	var user models.User

	if err := rows.Scan(&user); err != nil {
		return nil, fmt.Errorf("user.GetByID error: %w", err)
	}

	return &user, nil
}

func (u *UserRepository) Create(ctx context.Context, userDto createUserDto) error {
	_, err := u.db.Exec(ctx,
		`INSERT INTO users(name, email, password_hash, balance)
		VALUES($1, $2, $3, $4)`, userDto.Name, userDto.Email, userDto.PasswordHash, userDto.Balance)

	if err != nil {
		return fmt.Errorf("user.Create error: %w", err)
	}

	return nil
}
