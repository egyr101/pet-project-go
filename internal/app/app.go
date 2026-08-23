package app

import (
	"context"
	"fmt"
	"net/http"
	"pet-project/internal/config"
	"pet-project/internal/database"
	"pet-project/internal/logger"
	"pet-project/internal/middlewares"
	"pet-project/internal/use_cases/user"
)

type App struct {
	server *http.Server
}

func New(cfg *config.Config) (*App, error) {

	app := new(App)
	pool, err := database.NewPool(context.Background(), cfg.PostgresConfig)
	if err != nil {
		return nil, err
	}

	userRepo := user.NewUserRepository(pool)
	userService := user.NewUserService(userRepo)

	mux := http.NewServeMux()

	getUserHandler := http.Handler(middlewares.RequestID(middlewares.Recovery(middlewares.Logger(user.GetUserHandler(userService)))))
	createUserHandler := http.Handler(middlewares.RequestID(middlewares.Recovery(middlewares.Logger(user.CreateUserHandler(userService)))))
	updateUserHandler := http.Handler(middlewares.RequestID(middlewares.Recovery(middlewares.Logger(user.UpdateUserHandler(userService)))))
	deleteUserHandler := http.Handler(middlewares.RequestID(middlewares.Recovery(middlewares.Logger(user.DeleteUserHandler(userService)))))

	mux.Handle("GET /user/{id}", getUserHandler)
	mux.Handle("POST /user", createUserHandler)
	mux.Handle("PUT /user/{id}", updateUserHandler)
	mux.Handle("DELETE /user/{id}", deleteUserHandler)

	app.server = &http.Server{
		Handler:      mux,
		Addr:         ":" + cfg.Port,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	return app, nil
}

func (a *App) Run() error {
	logger.Info("http server starting...")
	err := a.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server was stop with err: %w", err)
	}

	logger.Info("server was stop")

	return nil
}

// Временное решение для создания бд
// func CreateDb(ctx context.Context, pool *pgxpool.Pool, createDBs ...func(context.Context, *pgxpool.Pool) error) error {
// 	for _, fn := range createDBs {
// 		if err := fn(ctx, pool); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
