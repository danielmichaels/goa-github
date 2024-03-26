package goagithub

import (
	"context"
	"log"

	usersv1 "github.com/danielmichaels/goa-github/gen/users_v1"
)

// users_v1 service example implementation.
// The example methods log the requests and return zero values.
type usersV1srvc struct {
	logger *log.Logger
}

// NewUsersV1 returns the users_v1 service implementation.
func NewUsersV1(logger *log.Logger) usersv1.Service {
	return &usersV1srvc{logger}
}

// ListUsers implements listUsers.
func (s *usersV1srvc) ListUsers(ctx context.Context, p *usersv1.ListUsersPayload) (res usersv1.UsersCollection, err error) {
	s.logger.Print("usersV1.listUsers")
	return
}

// RetrieveUser implements retrieveUser.
func (s *usersV1srvc) RetrieveUser(ctx context.Context, p *usersv1.RetrieveUserPayload) (res *usersv1.Users, err error) {
	res = &usersv1.Users{}
	s.logger.Print("usersV1.retrieveUser")
	return
}

// CreateUser implements createUser.
func (s *usersV1srvc) CreateUser(ctx context.Context, p *usersv1.CreateUserPayload) (res *usersv1.User, err error) {
	res = &usersv1.User{}
	s.logger.Print("usersV1.createUser")
	return
}

// DeleteUser implements deleteUser.
func (s *usersV1srvc) DeleteUser(ctx context.Context, p *usersv1.DeleteUserPayload) (err error) {
	s.logger.Print("usersV1.deleteUser")
	return
}
