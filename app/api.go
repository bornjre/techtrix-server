package app

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func addTransaction(w http.ResponseWriter, r *http.Request) {

	_ = getUserFromSession()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		print("body read error")
		errorResp(w, errors.New("body read error"))
		return
	}
	defer r.Body.Close()
	var transactioninfo map[string]string
	err = json.Unmarshal(body, &transactioninfo)
	if err != nil {
		print("body unmarshel error")
		return
	}

	addBlock(transactioninfo)

}

func getTransactions(w http.ResponseWriter, r *http.Request) {

	trans, err := GetAllTransactions()
	if err != nil {
		errorResp(w, err)
		return
	}
	generateResp(w, trans, nil)
}

func getTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hashid, ok := vars["hashid"]
	if !ok {
		errorResp(w, errors.New("no hashid"))
	}
	trans, err := GetTransaction(hashid)
	if err != nil {
		errorResp(w, err)
		return
	}
	generateResp(w, trans, nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["str"]
	if ok {
		errorResp(w, errors.New("invalid request send search str"))
		return
	}
	/*
		trans, err := GetAllTransactions()
		if err != nil {
			errorResp(w, err)
			return
		} */
	print(str)
	//search
	// TODO
}

func publish(w http.ResponseWriter, r *http.Request) {}

func subscribe(w http.ResponseWriter, r *http.Request) {}
