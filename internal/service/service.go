package service

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type Service struct {
	db *r.Session
}

//New Method
func New(db *r.Session) *Service {
	return &Service{db:db}
}
