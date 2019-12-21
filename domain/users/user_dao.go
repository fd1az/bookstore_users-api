package users

import (
	"fmt"

	"github.com/fdiaz7/bookstore_users-api/datasources/mysql/users_db"
	"github.com/fdiaz7/bookstore_users-api/utils/date_utils"
	"github.com/fdiaz7/bookstore_users-api/utils/errors"
	"github.com/fdiaz7/bookstore_users-api/utils/mysql_utils"
	_ "github.com/go-sql-driver/mysql"
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser    = "SELECT * FROM users WHERE id=? "
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err = users_db.Client.Ping(); err != nil {
		panic(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	fmt.Println(user)
	user.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}
