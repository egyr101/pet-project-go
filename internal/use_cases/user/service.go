package user

import "context"

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u UserService) Get(ctx context.Context, id int) (*userResponse, error) {

	user, err := u.repo.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, err
	}
	return &userResponse{
		Name:    user.Name,
		Email:   user.Email,
		Balance: user.Balance,
	}, nil
}

func (u *UserService) Create(ctx context.Context, userReq userRequest) (int, error) {

	id, err := u.repo.Create(ctx, userReq)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (u *UserService) Delete(ctx context.Context, userId int) (int, error) {

	id, err := u.repo.Delete(ctx, userId)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (u *UserService) Update(ctx context.Context, userId int, userReq userRequest) (int, error) {

	id, err := u.repo.Update(ctx, userId, userReq)
	if err != nil {
		return 0, err
	}

	return id, err
}
