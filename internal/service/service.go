package service

import (
	"github.com/hako/branca"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"net/url"
)

type Service struct {
	db *r.Session
	origin url.URL
	tokenKey string
}

func (s *Service) codec() *branca.Branca {
	cdc := branca.NewBranca(s.tokenKey)
	//cdc.SetTTL(uint32(tokenLifespan.Seconds()))
	return cdc
}

//New Method
func New(db *r.Session, origin url.URL, tokenKey string) *Service {
	return &Service{
		db:db,
		origin:origin,
		tokenKey:tokenKey,
	}
}
