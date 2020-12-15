package users

import (
	"fmt"
	"github.com/PMerdala/users-api/datasources/mysql/users_db"
	"github.com/PMerdala/users-api/utils/date_utils"
	"github.com/PMerdala/users-api/utils/errors"
	"github.com/PMerdala/users-api/utils/mysql_utils"
)

const (
	queryInsertUser              = "INSERT INTO users(first_name,last_name,email,date_created) values (?, ?, ?, ?);"
	queryUserById                = "SELECT id,first_name,last_name,email,date_created from users where id = ?;"
	queryUserByEmail             = "SELECT id,first_name,last_name,email,date_created from users where email = ?;"
	indexUniqueEmailErrorMessage = "email %s already exists"
	saveUserErrorMessage         = "error when trying to save user: %s"
	getUserByIdErrorMessage      = "error when trying to get user by id %d: %s"
	getUserByIdNotFound          = "Not found user by id %d"
	getUserByEmailErrorMessage   = "error when trying to get user by email %s: %s"
	getUserByEmailNotFound       = "Not found user by email %s"
)

func (user *User) Get() *errors.RestErr {
	stmt, prepareErr := users_db.Client.Prepare(queryUserById)
	if prepareErr != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf(getUserByIdErrorMessage, user.Id, prepareErr.Error()))
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParserError(getErr,
			fmt.Sprintf(getUserByIdNotFound,user.Id))
	}
	return nil
}

func (user *User) GetByEmail() *errors.RestErr {
	stmt, prepareErr := users_db.Client.Prepare(queryUserByEmail)
	if prepareErr != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf(getUserByEmailErrorMessage, user.Email, prepareErr.Error()))
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParserError(getErr,
			fmt.Sprintf(getUserByEmailNotFound,user.Email))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, prepareErr := users_db.Client.Prepare(queryInsertUser)
	if prepareErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf(saveUserErrorMessage, prepareErr.Error()))
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParserError(saveErr,
			fmt.Sprintf(indexUniqueEmailErrorMessage,user.Email))
	}
	userId, lastInsertIdErr := insertResult.LastInsertId()
	if lastInsertIdErr != nil {
		return mysql_utils.ParserError(lastInsertIdErr,
			fmt.Sprintf(saveUserErrorMessage,user.Email))
	}
	user.Id = userId
	return nil
}
