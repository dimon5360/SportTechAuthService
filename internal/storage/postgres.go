package storage

import (
	"app/main/models"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthUsersService) GetUserByIdFromDatabase(uuid uint64) (*models.User, error) {

	rows, err := s.db.Query(fmt.Sprintf("select * from users where id = %d;", uuid))
	if err != nil {
		log.Print(err)
		return &models.User{}, err
	}

	var user models.User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at)
		if err != nil {
			log.Print(err)
			return &models.User{}, err
		}
	}

	if err := rows.Err(); err != nil {
		log.Print(err)
		return &models.User{}, err
	}
	return &user, nil
}

func (s *AuthUsersService) GetUserByEmailFromDatabase(email string) (*models.User, error) {

	rows, err := s.db.Query(fmt.Sprintf("select * from users where email = '%s';", email))
	if err != nil {
		log.Print(err)
		return &models.User{}, err
	}

	var user models.User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at)
		if err != nil {
			log.Print(err)
			return &models.User{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return &models.User{}, err
	}

	if user.Id == 0 {
		return &models.User{}, fmt.Errorf("%s", "User not found")
	}
	return &user, nil
}

func (s *AuthUsersService) AddUserToDatabase(username string, email string, password string) (models.User, error) {

	format := "insert into users (username, email, password) values ('%s', '%s', '%s') ON conflict DO NOTHING RETURNING id, username;"

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	var user models.User
	rows := s.db.QueryRow(fmt.Sprintf(format, username, email, string(hash)))
	if err := rows.Err(); err != nil {
		log.Print(err)
		return user, err
	}

	err := rows.Scan(&user.Id, &user.Username)
	if err != nil {
		log.Print(err)
		return user, err
	}
	return user, nil
}
