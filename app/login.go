package app

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	ErrorLoginIncorrect = errors.New("Incorrect login details")
	ErrorAPI            = errors.New("could not generate api")
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var logininfo map[string]string
	err = deJson(body, &logininfo) //todo check fields
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	password := "hard123"
	username := "admin"
	log.Print(logininfo)

	if !(logininfo["password"] == password && logininfo["username"] == username) {
		http.Error(w, ErrorLoginIncorrect.Error(), http.StatusBadRequest)
		return
	}
	api := "46e4a1a76444f60d338e36121de0a32a"

	if api == "" {
		http.Error(w, ErrorAPI.Error(), http.StatusBadRequest)
		return
	}

	generateResp(w, map[string]string{"api": api}, nil)

}
