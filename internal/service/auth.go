package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/hako/branca"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"log"
	"strings"
	"time"
)

// KeyAuthUserID to use in context.
const KeyAuthUserID = ctxkey("auth_user_id")

const (
	tokenLifespan            = time.Hour * 24 * 14
	verificationCodeLifespan = time.Minute * 15
)

var (
	ErrInvalidToken = errors.New("invalid token")
	// ErrExpiredToken denotes that the token already expired.
	ErrExpiredToken = errors.New("expired token")
	ErrUnauthenticated = errors.New("access denied")
)
// DevLoginOutput response.
type DevLoginOutput struct {
	User      User      `json:"user"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type ctxkey string

func (s *Service) DevLogin(ctx context.Context, email string) (DevLoginOutput, error)  {
	 var out DevLoginOutput
	 var userModel User

	email = strings.TrimSpace(email)
	if !rxEmail.MatchString(email) {
		return out, ErrInvalidEmail
	}

	//Call the DB to check if passcode exists.
	cur, err := r.DB("test").Table("users").Filter(func(uu r.Term) r.Term {
		return uu.Field("email").Eq(email)
	}).Run(s.db)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close()
	for cur.Next(&userModel) {
		// If row was found - Create a token
		out.Token, err = s.codec().EncodeToString(userModel.ID)
		out.User = userModel
		out.ExpiresAt = time.Now().Add(tokenLifespan)
		return out, nil
	}
	if cur.Err() != nil {
		log.Fatal(err)
	}

	return out, ErrUserNotFound
}

// AuthUserIDFromToken decodes the token into a user ID.
func (s *Service) AuthUserIDFromToken(token string) (string, error) {
	uid, err := s.codec().DecodeToString(token)
	if err != nil {
		// We check error string because branca doesn't export errors.
		msg := err.Error()
		if msg == "invalid base62 token" || msg == "invalid token version" {
			return "", ErrInvalidToken
		}
		if msg == "token is expired" {
			return "", ErrExpiredToken
		}
		return "", fmt.Errorf("could not decode token: %v", err)
	}

	//if !rxUUID.MatchString(uid) {
	//	return "", ErrInvalidUserID
	//}
	return uid, nil
}

func (s *Service) codec() *branca.Branca {
	cdc := branca.NewBranca(s.tokenKey)
	cdc.SetTTL(uint32(tokenLifespan.Seconds()))
	return cdc
}
