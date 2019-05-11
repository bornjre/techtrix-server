package database

import (
	"crypto/sha256"
	"time"
)

type Transaction struct {
	HashID          string          `json:hashid`
	TransactionInfo TransactionInfo `json:transaction_info`
}

type TransactionInfo struct {
	Subject   string `json:subject`
	Body      string `json:body`
	CreatedAt string `json:created_at`
	Amount    int64  `json:amount`
	User      string `json:user`
}

var (
	TransactionBucketName = []byte("transactions")
)

func NewTransaction(hashid string,subject string, body string, amount int64, user string,) *Transaction {

	if err != nil {
		return nil
	}
	return &TransactionInfo{
		HashID:    hashid,
		Subject : subject,
		Body : body,
		User:      user,
		CreatedAt: time.Now().Unix(),
		Amount: amount
	}
}

func AddTransaction(trans *Transaction) error {

	transbyte, err := Encode(trans)
	if err != nil {
		return err
	}

	err = DB.Create([]byte(trans.HashID), transbyte, TransactionBucketName)
	if err != nil {
		return err
	}
	return nil
}

func GetTransaction(hashid string) (*Transaction, error) {
	transaction := &Transaction{}
	transbyte, err := DB.Read([]byte(hashid), TransactionBucketName)

	if err != nil {
		return nil, err
	}

	err = Decode(transbyte, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func GetAllTransactions() ([]*Transaction, error) {

	transactions := []*Transaction{}

	transactionsbyte, err := DB.ReadAll(TransactionBucketName)
	if err != nil {
		return nil, err
	}

	for _, val := range transactionsbyte {
		var transaction Transaction
		err = Decode(val, &transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}
