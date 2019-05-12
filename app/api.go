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

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		printl("body read error")
		errorResp(w, errors.New("body read error"))
		return
	}
	defer r.Body.Close()
	var transactioninfo map[string]string
	err = json.Unmarshal(body, &transactioninfo)
	if err != nil {
		printl("body unmarshel error")
		return
	}

	addBlock(transactioninfo)

}

func updateTransaction(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		printl("body read error")
		errorResp(w, errors.New("body read error"))
		return
	}
	defer r.Body.Close()
	var transactioninfo map[string]string
	err = json.Unmarshal(body, &transactioninfo)
	if err != nil {
		printl("body unmarshel error")
		return
	}

	updateBlock(transactioninfo)

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
	printl(str)
	//search
	// TODO
}

func verify(w http.ResponseWriter, r *http.Request) {
	ok, err := VerifyBlockChain()
	if ok {
		generateResp(w, "Blockchain is verified", nil)
		return
	}

	errorResp(w, err)

}

func publish(w http.ResponseWriter, r *http.Request) {}

func subscribe(w http.ResponseWriter, r *http.Request) {
	printl("websocket!!")

}
