package service

import (
	"context"
	"crypto/rand"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"io"
	"log"
)

//Passcodes is used to save passcode on db
type Passcode struct {
	ID   string `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Username string `json:"username,omitempty"`
}

//GeneratePasscodes will create 5 random passcode and store to db
func (s *Service) GeneratePasscodes(ctx context.Context, username string) ([]string, error) {
	//Store the passcodes in the db
	codes := randomPasscodes()
	for _, v := range codes {
		log.Printf("Generated passcodes are : %s", v)
		var data = map[string]interface{}{
			"username": username,
			"Code": v,
		}
		log.Printf("Inserting code %s", data)
		//Save then to database and then
		_, err := r.DB("test").Table("passcodes").Insert(data).RunWrite(s.db)
		if err != nil {
			log.Fatal(err)
		}
	}

	return codes, nil
}

//GetPassCodes is used to list all pascodes to the UI
func (s *Service) GetPassCodes(ctx context.Context, username string) ([]Passcode, error) {
	log.Printf("Getting passcodes for username : %s", username)
	var out []Passcode
	cur, err := r.DB("test").Table("passcodes").Filter(func(uu r.Term) r.Term {
		return uu.Field("username").Eq(username)
	}).Run(s.db)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close()

	if err != nil {
		return out, err
	}

	defer cur.Close()
	err = cur.All(&out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func randomPasscodes() []string {
	var passcodes []string
	for i := 0; i < 5; i++ {
		s := encodeToString(6)
		passcodes = append(passcodes, s)
	}
	return passcodes
}

func encodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

//ValidatePasscode will validate the passcode on db
func (s *Service) ValidatePasscode(passcode string) (bool, error) {

	var out Passcodes
	var outBool = false
	//Call the DB to check if passcode exists.
	cur, err := r.DB("test").Table("passcodes").Filter(func(uu r.Term) r.Term {
		return uu.Field("Code").Eq(passcode)
	}).Run(s.db)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close()
	for cur.Next(&out) {
		// If row was found
		outBool = true
	}
	if cur.Err() != nil {
		log.Fatal(err)
	}

	return outBool, nil

}
