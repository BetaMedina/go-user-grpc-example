package repositories

import (
	"database/sql"
	"grpc-user/pkg/models"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (u *users) Save(user *models.UserModel) (int64, error) {
	statement, err := u.db.Prepare("INSERT INTO users (name,email,age,document) VALUES (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	row, err := statement.Exec(user.Name, user.Email, user.Age, user.Document)
	if err != nil {
		return 0, err
	}
	lastInsertedId, _ := row.LastInsertId()
	return lastInsertedId, nil
}

func (u *users) FindOne(id int64) (*models.UserModel, error) {
	var user models.UserModel
	row, err := u.db.Query("SELECT id,name,email,document,age FROM users where id = ?", id)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Email, &user.Document, &user.Age)
	}
	return &user, nil
}
