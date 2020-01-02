package service

import (
	"context"
	"github.com/hako/branca"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"log"
	"strings"
	"time"
)

var (
	tokenLifespan            = time.Hour * 24 * 14
	verificationCodeLifespan = time.Minute * 15
)
// DevLoginOutput response.
type DevLoginOutput struct {
	User      User      `json:"user"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

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

func (s *Service) codec() *branca.Branca {
	cdc := branca.NewBranca(s.tokenKey)
	cdc.SetTTL(uint32(tokenLifespan.Seconds()))
	return cdc
}
