package service

import (
	"errors"
	"github.com/potterbl/agent/internal/repository"
	"github.com/potterbl/agent/pkg/types"
)

type UserService interface {
	CreateUser(req *types.CreateUserRequest) (*types.UserResponse, error)
	GetUserByID(id uint) (*types.UserResponse, error)
	GetAllUsers() (*types.UsersListResponse, error)
	UpdateUser(id uint, req *types.UpdateUserRequest) (*types.UserResponse, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser создает нового пользователя
func (s *userService) CreateUser(req *types.CreateUserRequest) (*types.UserResponse, error) {
	// Валидация бизнес-правил
	if len(req.Name) < 2 {
		return nil, errors.New("name must be at least 2 characters long")
	}

	user := &types.User{
		Email: req.Email,
		Name:  req.Name,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return s.userToResponse(user), nil
}

// GetUserByID получает пользователя по ID
func (s *userService) GetUserByID(id uint) (*types.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return s.userToResponse(user), nil
}

// GetAllUsers получает список всех пользователей
func (s *userService) GetAllUsers() (*types.UsersListResponse, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userResponses := make([]types.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = *s.userToResponse(&user)
	}

	return &types.UsersListResponse{
		Users: userResponses,
		Total: len(userResponses),
	}, nil
}

// UpdateUser обновляет данные пользователя
func (s *userService) UpdateUser(id uint, req *types.UpdateUserRequest) (*types.UserResponse, error) {
	// Получаем существующего пользователя
	existingUser, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Обновляем только переданные поля
	updatedUser := *existingUser
	
	if req.Name != nil {
		if len(*req.Name) < 2 {
			return nil, errors.New("name must be at least 2 characters long")
		}
		updatedUser.Name = *req.Name
	}
	
	if req.Email != nil {
		updatedUser.Email = *req.Email
	}

	if err := s.userRepo.UpdateUser(id, &updatedUser); err != nil {
		return nil, err
	}

	return s.userToResponse(&updatedUser), nil
}

// DeleteUser удаляет пользователя
func (s *userService) DeleteUser(id uint) error {
	// Проверим, что пользователь существует
	if _, err := s.userRepo.GetUserByID(id); err != nil {
		return err
	}

	return s.userRepo.DeleteUser(id)
}

// userToResponse конвертирует User в UserResponse
func (s *userService) userToResponse(user *types.User) *types.UserResponse {
	return &types.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}