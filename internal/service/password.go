package service

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/vinchauhan/two-f-gates/internal/util"
)

func (s *Service) GeneratePasscodes(passcodes []string) error {
	//Store the passcodes in the db
	txn := s.db.NewTransaction(true)
	defer txn.Discard()

	//Loop through the slice
	for k, v := range passcodes {
		err := txn.Set(util.IntToBytes(k), []byte(v))
		if err != nil {
			return err
		}
	}

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetPassCodes() ([]string, error) {
	var passCodes []string
	txn := s.db.NewTransaction(true)
	defer txn.Discard()

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return passCodes, err
	}
	return passCodes, nil
}
