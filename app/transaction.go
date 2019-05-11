package app

import (
	"github.com/bornjre/techtrix-server/app/database"
)

var (
	TransactionBucketName = []byte("transactions")
)

func AddTransaction(trans *Transaction) error {

	transbyte, err := database.Encode(trans)
	if err != nil {
		return err
	}

	err = database.DB.Create([]byte(trans.HashID), transbyte, TransactionBucketName)
	if err != nil {
		return err
	}
	return nil
}

func GetTransaction(hash_id string) (*Transaction, error) {
	transaction := &Transaction{}
	transbyte, err := database.DB.Read([]byte(hash_id), TransactionBucketName)

	if err != nil {
		return nil, err
	}

	err = database.Decode(transbyte, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func GetAllTransactions() ([]*Transaction, error) {

	transactions := []*Transaction{}

	transactionsbyte, err := database.DB.ReadAll(TransactionBucketName)
	if err != nil {
		return nil, err
	}

	for _, val := range transactionsbyte {
		var transaction Transaction
		err = database.Decode(val, &transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}
