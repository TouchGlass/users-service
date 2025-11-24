package grpc

import (
	"context"
	userpb "github.com/TouchGlass/project-protos/proto/user"
	"github.com/TouchGlass/users-service/internal/user"
	"strconv"
)

type Handler struct {
	svc user.UserService
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	userModel := user.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err, createdUser := h.svc.CreateUser(userModel)
	if err != nil {
		return nil, err
	}

	resp := &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:       uint32(createdUser.ID),
			Email:    createdUser.Email,
			Password: createdUser.Password,
		},
	}
	return resp, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	userModel := user.User{
		Email:    req.Email,
		Password: req.Password,
	}
	id := req.Id
	strId := strconv.Itoa(int(id))

	updatedUser, err := h.svc.UpdateUser(strId, userModel)
	if err != nil {
		return nil, err
	}

	resp := &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:       id,
			Email:    updatedUser.Email,
			Password: updatedUser.Password,
		},
	}
	return resp, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.User, error) {
	id := req.Id
	strId := strconv.Itoa(int(id))
	user, err := h.svc.GetUserByID(strId)
	if err != nil {
		return nil, err
	}

	resp := &userpb.User{
		Id:       id,
		Email:    user.Email,
		Password: user.Password,
	}
	return resp, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.GetUsers()
	if err != nil {
		return nil, err
	}

	respUsers := make([]*userpb.User, 0, len(users))
	for _, user := range users {
		respUsers = append(respUsers, &userpb.User{
			Id:       uint32(user.ID),
			Email:    user.Email,
			Password: user.Password,
		})
	}

	resp := &userpb.ListUsersResponse{
		Users: respUsers,
	}
	return resp, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	id := req.Id
	strId := strconv.Itoa(int(id))

	err := h.svc.DeleteUserByID(strId)
	if err != nil {
		return &userpb.DeleteUserResponse{Success: false}, err
	}

	return &userpb.DeleteUserResponse{Success: true}, err
}
