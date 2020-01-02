package service

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"net/url"
)

type Service struct {
	db *r.Session
	origin url.URL
	tokenKey string
}

//New Method
func New(db *r.Session, origin url.URL, tokenKey string) *Service {
	return &Service{
		db:db,
		origin:origin,
		tokenKey:tokenKey,
	}
}
