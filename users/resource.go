package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"encore.dev/storage/sqldb"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// defaultDB connects to the "postgres" database.
var db = sqldb.Named("postgres")

//encore:service
type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FullName  string    `json:"full_name"`
	Active    bool      `json:"active"`
}

//encore:api public path=/user/get/:id
func (*User) Get(ctx context.Context, id string) (*User, error) {

	var user User
	err := db.QueryRow(ctx, GET, id).Scan(&user.ID, &user.FullName, &user.Active)
	return &user, err
}

//encore:api public path=/user/get
func (*User) GetAll(ctx context.Context) (*Response, error) {

	var users []User

	rows, err := db.Query(ctx, GET_ALL)

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
func (*User) Add(ctx context.Context, payload User) (*User, error) {
	var user User
	err := db.QueryRow(ctx, INSERT, payload.FullName, payload.Active).Scan(&user.ID, &user.FullName, &user.Active)
	return &user, err
}

//encore:api public path=/user/delete/:id
func (*User) Delete(ctx context.Context, id string) (*Response, error) {

	result, err := db.Exec(ctx, DELETE, id)
	if result.RowsAffected() == 0 {
		return &Response{Message: "user not deleted"}, errors.New("user not found")
	}
	return &Response{Message: "user deleted"}, err
}

//encore:api public path=/user/delete
func (*User) DeleteAll(ctx context.Context) (*Response, error) {

	result, err := db.Exec(ctx, DELETE_ALL)
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
