package dto

import "smart-payment-dev-be/internal/core/domain"

type UserDTO struct{}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	User User `json:"user"`
}

type UsersResponse struct {
	Users []User `json:"users"`
	Total int    `json:"total"`
}

func (ud *UserDTO) CreateUserTransformIn(req CreateUserRequest) domain.User {
	return domain.User{
		Name:     req.Name,
		Password: req.Password,
	}
}

func (ud *UserDTO) UpdateUserTransformIn(req UpdateUserRequest) domain.User {
	return domain.User{
		Id:       req.Id,
		Name:     req.Name,
		Password: req.Password,
	}
}

func (ud *UserDTO) GetUserTransformOut(res domain.User) UserResponse {
	return UserResponse{
		User: User{
			Id:       res.Id,
			Name:     res.Name,
			Password: res.Password,
		},
	}
}

func (ud *UserDTO) GetUsersTransformOut(req []domain.User, total int) (res UsersResponse) {
	var users []User
	for _, item := range req {
		user := User{
			Id:       item.Id,
			Name:     item.Name,
			Password: item.Password,
		}
		users = append(users, user)
	}
	return UsersResponse{
		Users: users,
		Total: total,
	}
}
