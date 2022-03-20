package user

import (
	"context"
	"fmt"
	"grpc-user/pkg/models"
	"grpc-user/pkg/store"
	"grpc-user/pkg/store/repositories"
)

type Server struct{}

func (s *Server) SaveUser(ctx context.Context, message *CreateUserInput) (*CreateUserOutput, error) {
	db, err := store.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	user := models.UserModel{
		Name:     message.Name,
		Email:    message.Email,
		Document: message.Document,
		Age:      message.Age,
	}
	repository := repositories.NewUserRepository(db)
	userId, err := repository.Save(&user)
	if err != nil {
		return nil, err
	}
	return &CreateUserOutput{IdUser: userId, Message: "User created"}, nil
}

func (s *Server) ReadUser(ctx context.Context, message *ReadUserInput) (*ReadUserOutput, error) {
	db, err := store.Connect()
	if err != nil {
		return nil, fmt.Errorf("Error on connect with database: %v", err)
	}
	defer db.Close()
	repository := repositories.NewUserRepository(db)
	userRow, err := repository.FindOne(message.Id)
	if err != nil {
		return nil, fmt.Errorf("Error on find user: %v", err)
	}
	return &ReadUserOutput{Id: userRow.Id, Name: userRow.Name, Email: userRow.Email, Document: userRow.Document, Age: userRow.Age}, nil
}
