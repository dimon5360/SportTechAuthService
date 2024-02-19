package storage

import (
	"app/main/models"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthUsersService) GetUserWithFormat(sql string) (*models.User, error) {

	rows, err := s.db.Query(sql)
	if err != nil {
		log.Print(err)
		return &models.User{}, err
	}

	var user models.User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
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

func (s *AuthUsersService) GetUserByIdFromDatabase(uuid uint64) (*models.User, error) {

	return s.GetUserWithFormat(fmt.Sprintf("select * from users where id = %d;", uuid))
}

func (s *AuthUsersService) GetUserByEmailFromDatabase(email string) (*models.User, error) {

	return s.GetUserWithFormat(fmt.Sprintf("select * from users where email = '%s';", email))
}

func (s *AuthUsersService) AddUserToDatabase(email string, password string) (*models.User, error) {

	format := "insert into users (email, password) values ('%s', '%s') ON conflict DO NOTHING RETURNING id;"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	var user models.User
	rows := s.db.QueryRow(fmt.Sprintf(format, email, string(hash)))
	if err := rows.Err(); err != nil {
		log.Print(err)
		return &user, err
	}

	err := rows.Scan(&user.Id)
	if err != nil {
		log.Print(err)
		return &user, err
	}
	return &user, nil
}
