package user

import (
	"context"
	"errors"
	"fmt"

	"encore.dev/storage/sqldb"
)

//encore:api public path=/user/get/:id
func Get(ctx context.Context, id string) (*User, error) {

	var user User
	err := sqldb.QueryRow(ctx, GET, id).Scan(&user.ID, &user.FullName, &user.Active)
	return &user, err
}

//encore:api public path=/user/get
func GetAll(ctx context.Context) (*Response, error) {

	var users []User

	rows, err := sqldb.Query(ctx, GET_ALL)

	for index := 0; rows.Next(); index++ {
		var user User
		if err := rows.Scan(&user.ID, &user.FullName, &user.Active); err != nil {
			return &Response{}, err
		}

		users = append(users, user)
	}
	return &Response{Users: users, Message: "users fetched"}, err
}

//encore:api method=POST public path=/user/add
func Add(ctx context.Context, payload User) (*User, error) {
	var user User
	err := sqldb.QueryRow(ctx, INSERT, payload.FullName, payload.Active).Scan(&user.ID, &user.FullName, &user.Active)
	return &user, err
}

//encore:api public path=/user/delete/:id
func Delete(ctx context.Context, id string) (*Response, error) {

	result, err := sqldb.Exec(ctx, DELETE, id)
	if result.RowsAffected() == 0 {
		return &Response{Message: "user not deleted"}, errors.New("user not found")
	}
	return &Response{Message: "user deleted"}, err
}

//encore:api public path=/user/delete
func DeleteAll(ctx context.Context) (*Response, error) {

	result, err := sqldb.Exec(ctx, DELETE_ALL)
	if result.RowsAffected() == 0 {
		return &Response{Message: "users not deleted"}, errors.New("users not found")
	}

	return &Response{Message: fmt.Sprintf("%v users deleted", result.RowsAffected())}, err
}

type Response struct {
	Message string
	Users   []User
}

// ==================================================================

// Encore comes with a built-in development dashboard for
// exploring your API, viewing documentation, debugging with
// distributed tracing, and more. Visit your API URL in the browser:
//
//     http://localhost:4000
//

// ==================================================================

// Next steps
//
// 1. Deploy your application to the cloud with a single command:
//
//     git push encore
//
// 2. To continue exploring Encore, check out one of these topics:
//
//    Building a Slack bot:  https://encore.dev/docs/tutorials/slack-bot
//    Building a REST API:   https://encore.dev/docs/tutorials/rest-api
//    Using SQL databases:   https://encore.dev/docs/develop/sql-database
//    Authenticating users:  https://encore.dev/docs/develop/auth
