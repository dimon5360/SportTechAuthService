package storage

import (
	"app/main/utils"
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

func CreateService() *AuthUsersService {
	return &AuthUsersService{}
}

func (s *AuthUsersService) Init() {

	var connString string

	env := utils.Env()

	connString = fmt.Sprintf("postgresql://%s:%s@%s/%s",
		env.Value("POSTGRES_USER"),
		env.Value("POSTGRES_PASSWORD"),
		env.Value("POSTGRES_HOST"),
		env.Value("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", connString+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	s.db = db
}

func invalidUserInfoResponse(err error) (*proto.UserInfoResponse, error) {
	return &proto.UserInfoResponse{}, err
}

func (s *AuthUsersService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserInfoResponse, error) {

	user, err := s.GetUserByIdFromDatabase(req.Id)
	if err != nil {
		return invalidUserInfoResponse(err)
	}

	return &proto.UserInfoResponse{
		Id:        user.Id,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.Created_at),
		UpdatedAt: timestamppb.New(user.Updated_at),
	}, nil
}

func (s *AuthUsersService) AuthUser(ctx context.Context, req *proto.AuthUserRequest) (*proto.UserInfoResponse, error) {

	user, err := s.GetUserByEmailFromDatabase(req.Email)
	if err != nil {
		return invalidUserInfoResponse(fmt.Errorf("%s: %v", "Failed handling user data", err))
	}

	if !user.ValidateCredentials(req) {
		return invalidUserInfoResponse(fmt.Errorf("%s: %v", "Invalid email or password", err))
	}

	return &proto.UserInfoResponse{
		Id:        user.Id,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.Created_at),
		UpdatedAt: timestamppb.New(user.Updated_at),
	}, nil
}

func (s *AuthUsersService) CreateUser(ctx context.Context, req *proto.CreateUserRequst) (*proto.UserInfoResponse, error) {
	user, err := s.AddUserToDatabase(req.Email, req.Password)
	return &proto.UserInfoResponse{Id: user.Id}, err
}
