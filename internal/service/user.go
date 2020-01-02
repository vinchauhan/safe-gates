package service

import (
	"context"
	"errors"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"log"
	"path"
	"regexp"
	"strings"
)

var (
	rxEmail    = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	rxUsername = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{0,17}$`)
	avatarsDir = path.Join("web", "static", "img", "avatars")
)

var (
	// ErrInvalidUserID denotes an invalid user id; that is not uuid.
	ErrInvalidUserID = errors.New("invalid user id")
	// ErrInvalidEmail denotes an invalid email address.
	ErrInvalidEmail = errors.New("invalid email")
	// ErrInvalidUsername denotes an invalid username.
	ErrInvalidUsername = errors.New("invalid username")
	ErrUserNameExists = errors.New("username exists")
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID string `json:"id,omitempty"`
	Email string `json:"email"`
	Username string `json:"username"`
}

//CreateUser is used to Create a new User
func (s *Service) CreateUser(ctx context.Context, email , username string) error {
	var outUser User
	email = strings.TrimSpace(email)
	if !rxEmail.MatchString(email) {
		return ErrInvalidEmail
	}

	username = strings.TrimSpace(username)
	if !rxUsername.MatchString(username) {
		return ErrInvalidUsername
	}

	var userData = map[string]interface{}{
		"id": username,
		"email": email,
		"username": username,
	}
	//In RethinkDB we need to do two operations to see if the user exists first and if not then create.

	//Call the DB to check if username exists.
	cur, err := r.DB("test").Table("users").Filter(func(uu r.Term) r.Term {
		return uu.Field("username").Eq(username)
	}).Run(s.db)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close()
	for cur.Next(&outUser) {
		// If row was found
		return ErrUserNameExists
	}
	if cur.Err() != nil {
		log.Fatal(err)
	}

	//If username not found then store it in the db
	r.DB("test").Table("users").Insert(userData).Run(s.db)
	return nil
}
