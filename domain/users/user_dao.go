package users

import (
	"fmt"
	"github.com/PMerdala/users-api/datasources/mysql/users_db"
	"github.com/PMerdala/users-api/utils/date_utils"
	"github.com/PMerdala/users-api/utils/errors"
	"strings"
)

const (
	queryInsertUser              = "INSERT INTO users(first_name,last_name,email,date_created) values (?, ?, ?, ?);"
	queryUserById                = "SELECT id,first_name,last_name,email,date_created from users where id = ?;"
	queryUserByEmail             = "SELECT id,first_name,last_name,email,date_created from users where email = ?;"
	mysqlIndexUniqueEmailError   = "'users_email_uindex'"
	mysqlNotRowsError            = "sql: no rows in result set"
	indexUniqueEmailErrorMessage = "email %s already exists"
	saveUserErrorMessage         = "error when trying to save user: %s"
	getUserByIdErrorMessage      = "error when trying to get user by id %d: %s"
	getUserByIdNotFound          = "Not found user by id %d"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUserById)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf(getUserByIdErrorMessage, err.Error()))
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), mysqlNotRowsError) {
			return errors.NewNotFoundError(
				fmt.Sprintf(getUserByIdNotFound, user.Id))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf(getUserByIdErrorMessage, user.Id, err.Error()))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf(saveUserErrorMessage, err.Error()))
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), mysqlIndexUniqueEmailError) {
			return errors.NewBadRequestError(
				fmt.Sprintf(indexUniqueEmailErrorMessage, user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf(saveUserErrorMessage, err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf(saveUserErrorMessage, err.Error()))
	}
	user.Id = userId
	return nil
}
