package storage

import (
	"app/main/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func CreateStorage() Storage {
	var store Storage
	return store
}

func (s *Storage) Init(conn_string string) {

	db, err := sql.Open("postgres", conn_string)
	if err != nil {
		log.Fatal(err)
	}

	s.db = db
}

func (s *Storage) GetUserById(uuid string) models.User {

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
