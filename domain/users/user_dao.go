package users

import (
	"fmt"

	"github.com/michaeljwood/bookstore_users-api/datasources/mysql/users_db"
	"github.com/michaeljwood/bookstore_users-api/logger"
	"github.com/michaeljwood/bookstore_users-api/utils/errors"
)

const (
	errorNoRows           = "no rows in result set"
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	querySelectUser       = "SELECT * FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT * FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Preparex(querySelectUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.Unsafe().QueryRowx(user.Id)
	if err := result.StructScan(user); err != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Preparex(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare get save statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("error when trying to insert user", err)
		return errors.NewInternalServerError("database error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get inserted user id", err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Preparex(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Preparex(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Preparex(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	results := []User{}
	if err = stmt.Unsafe().Select(&results, status); err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError("database error")
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
