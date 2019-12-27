package handler

import (
	"net/http"
	"os"
)

type spaFileSystem struct {
	root http.FileSystem
}

//Implement Open method for an instance of spaFileSystem to accepted as argument to http.FileServer()

func (sfs *spaFileSystem) Open(name string) (http.File, error) {
	f, err := sfs.root.Open(name)
	if os.IsNotExist(err) {
		return sfs.root.Open("index.html")
	}
	return f, err
}
