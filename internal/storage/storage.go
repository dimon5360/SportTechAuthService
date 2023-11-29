package storage

import (
	"app/main/models"
	"context"
	"database/sql"
	"fmt"
	"log"

	proto "github.com/dimon5360/SportTechProtos/gen/go/proto"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthUsersService struct {
	proto.UnimplementedAuthUsersServiceServer

	db *sql.DB
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

func (s *AuthUsersService) GetUserById(uuid uint64) models.User {

	rows, err := s.db.Query(fmt.Sprintf("select * from users where id = %d;", uuid))
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
		Id:        user.Id,
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
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at)
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
			Id:        user.Id,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.Created_at),
			UpdatedAt: timestamppb.New(user.Updated_at),
		}, nil
	}

	return &proto.UserInfoResponse{}, fmt.Errorf("%s", "Invalid email or password")
}

func (s *AuthUsersService) CreateUser(ctx context.Context, req *proto.CreateUserRequst) (*proto.UserInfoResponse, error) {

	log.Print(req.Username, req.Email, req.Password)

	query := fmt.Sprintf("insert into users (username, email, password) values ('%s', '%s', '%s') ON conflict DO NOTHING;",
		req.Username,
		req.Email,
		req.Password)

	res, err := s.db.Exec(query)
	if err != nil {
		log.Print(err)
		return &proto.UserInfoResponse{}, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
		return &proto.UserInfoResponse{}, err
	}

	if n == 0 {
		return &proto.UserInfoResponse{}, fmt.Errorf("user already exists")
	}

	return &proto.UserInfoResponse{}, err
}