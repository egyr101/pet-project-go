package user

import (
	"context"
	"fmt"
	"pet-project/internal/models"

	"github.com/jackc/pgx/v5"
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
	var user models.User
	err := u.db.QueryRow(ctx, `
		SELECT *
			FROM users
			WHERE id = $1;
	`, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Balance,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("user.Get error: %w", err)
	}

	return &user, nil
}

func (u *UserRepository) Create(ctx context.Context, userReq *userRequest) (int, error) {
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
		WHERE id = $1
		RETURNING id`, userId,
	).Scan(&id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("user.Delete error: %w", err)
	}

	return id, nil
}

func (u *UserRepository) Update(ctx context.Context, userId int, userReq *userRequest) (int, error) {
	var id int

	err := u.db.QueryRow(ctx,
		`UPDATE users SET
			name = $1,
			email = $2,
			password_hash = $3
		WHERE id = $4
		RETURNING id`, userReq.Name, userReq.Email, userReq.PasswordHash, userId,
	).Scan(&id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("user.Update error: %w", err)
	}

	return id, nil
}
