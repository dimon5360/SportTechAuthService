package storage

import (
	"app/main/models"
	"app/main/proto"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthUsersService struct {
	proto.UnimplementedAuthUsersServiceServer

	db * sql.DB
}

func CreateService() AuthUsersService {
	return AuthUsersService{}
}

func (s *AuthUsersService) Init(conn_string string) {

	db, err := sql.Open("postgres", conn_string)
	if err != nil {
		log.Fatal(err)
	}

	s.db = db
}

func (s *AuthUsersService) GetUserById(uuid string) models.User {

	rows, err := s.db.Query(fmt.Sprintf("select * from users where id = %s;", uuid))
	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at); err != nil {
			log.Fatal(err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return user
}

func (s *AuthUsersService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserInfoResponse, error) {

	user := s.GetUserById(req.Id)
	log.Print(user)

	return &proto.UserInfoResponse{
		Id:        fmt.Sprintf("%d", user.Id),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.Created_at),
		UpdatedAt: timestamppb.New(user.Updated_at),
	}, nil
}

func (s *AuthUsersService) AuthUser(ctx context.Context, req *proto.AuthUserRequest) (*proto.UserInfoResponse, error) {

	log.Print(req.Email + req.Password)

	rows, err := s.db.Query(fmt.Sprintf("select * from users where email = '%s';", req.Email))
	if err != nil {
		log.Print(err)
		return &proto.UserInfoResponse{}, err
	}

	var user models.User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at);
		if err != nil {
			log.Print(err)
			return &proto.UserInfoResponse{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return &proto.UserInfoResponse{}, err
	}

	if user.Id == 0 {
		return &proto.UserInfoResponse{}, fmt.Errorf("%s", "User not found")
	}

	if user.Email == req.Email && user.Password == req.Password {
		return &proto.UserInfoResponse{
			Id:        fmt.Sprintf("%d", user.Id),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.Created_at),
			UpdatedAt: timestamppb.New(user.Updated_at),
		}, nil
	}

	return &proto.UserInfoResponse{}, fmt.Errorf("%s", "Invalid email or password")
}
