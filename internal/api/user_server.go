package api

import (
	"context"
	"errors"
	"gophkeeper/internal/auth"
	"gophkeeper/internal/logger"
	pbUser "gophkeeper/internal/proto/user"
	userservice "gophkeeper/internal/service/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService interface {
	Register(ctx context.Context, login, password string) error
	Login(ctx context.Context, login, password string) error
}

type UserServer struct {
	pbUser.UnimplementedUsersServer
	service userService
}

func NewUserServer(service userService) *UserServer {
	return &UserServer{service: service}
}

func (us *UserServer) RegisterUser(ctx context.Context, in *pbUser.RegisterUserRequest) (*pbUser.RegisterUserResponse, error) {
	if in.Login == "" || in.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "login and password must be not empty")
	}

	err := us.service.Register(ctx, in.Login, in.Password)

	if err != nil {
		if errors.Is(err, userservice.ErrorLoginExists) {
			return nil, status.Errorf(codes.AlreadyExists, "user with login %s already exists")
		}

		logger.Log.Error("Error while register user: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	generatedJWT, err := auth.BuildJWTString(in.Login)

	if err != nil {
		logger.Log.Error("Error while create token: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	logger.Log.Info("User %s registered", in.Login)
	res := &pbUser.RegisterUserResponse{AccessToken: generatedJWT}
	return res, nil
}

func (us *UserServer) LoginUser(ctx context.Context, in *pbUser.LoginUserRequest) (*pbUser.LoginUserResponse, error) {
	if in.Login == "" || in.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "login and password must be not empty")
	}

	err := us.service.Login(ctx, in.Login, in.Password)

	if err != nil {
		logger.Log.Error("Error while login user: ", err)
		if errors.Is(err, userservice.ErrorWrongCredentials) {
			return nil, status.Errorf(codes.Unauthenticated, "wrong credentials")
		}

		return nil, status.Errorf(codes.Internal, "internal error")
	}

	generatedJWT, err := auth.BuildJWTString(in.Login)

	if err != nil {
		logger.Log.Error("Error while create token: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	logger.Log.Debugf("User %s logged in", in.Login)
	res := &pbUser.LoginUserResponse{AccessToken: generatedJWT}
	return res, nil
}
