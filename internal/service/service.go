package service

import "github.com/dgraph-io/badger"

type Service struct {
	db *badger.DB
}

func New(db *badger.DB) *Service {
	return &Service{db:db}
}
