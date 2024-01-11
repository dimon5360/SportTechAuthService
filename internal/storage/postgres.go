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
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at);
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

func (s *AuthUsersService) AddUserToDatabase(username string, email string, password string) (uint64, error) {

	sql_query := "insert into users (username, email, password) values ('%s', '%s', '%s') ON conflict DO NOTHING RETURNING id;"

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	query := fmt.Sprintf(sql_query, username, email, string(hash))

	lastInsertId := 0
	rows := s.db.QueryRow(query)
	if err := rows.Err(); err != nil {
		log.Print(err)
		return 0, err
	}

	err := rows.Scan(&lastInsertId)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return uint64(lastInsertId), nil
}