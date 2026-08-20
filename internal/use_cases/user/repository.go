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

func (u UserRepository) Get(ctx context.Context, id int) (*models.User, error) {
	rows, err := u.db.Query(ctx, `
		SELECT *
			FROM users
			WHERE id = $1;
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

func (u *UserRepository) Create(ctx context.Context, userReq userRequest) (int, error) {
	var id int
	err := u.db.QueryRow(ctx,
		`INSERT INTO users(name, email, password_hash)
		VALUES($1, $2, $3)
		RETURNING id;`, userReq.Name, userReq.Email, userReq.PasswordHash,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("user.Create error: %w", err)
	}

	return id, nil
}

func (u *UserRepository) Delete(ctx context.Context, userId int) (int, error) {
	var id int

	err := u.db.QueryRow(ctx,
		`DELETE FROM users
		WHERE id = $1`, id,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("user.Delete error: %w", err)
	}

	return id, nil
}
