package service

import (
	"crypto/rand"
	"log"
	"io"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)
//Passcodes is used to save passcode on db
type Passcodes struct {
	ID   string `gorethink:"id,omitempty"`
	Code string `gorethink:"code,omitempty"`
}

//GeneratePasscodes will create 5 random passcode and store to db
func (s *Service) GeneratePasscodes(passcodes []string) ([]string, error) {
	//Store the passcodes in the db
	codes := randomPasscodes()
	for _, v := range codes {
		log.Printf("Generated passcodes are : %s", v)
		var data = map[string]interface{}{
			"Code":  v,
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

func (s *Service) GetPassCodes() ([]string, error) {
	return nil, nil
}

func randomPasscodes() ([]string) {
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