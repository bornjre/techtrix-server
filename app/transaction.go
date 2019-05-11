package app

import (
	"github.com/bornjre/techtrix-server/app/database"
)

var (
	TransactionBucketName = []byte("transactions")
	OrderedTransactions   = []byte("orderedransactions")
	OrderedKey            = []byte("1")
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

func GetAllTransactions() (map[string]*Transaction, error) {

	transactions := make(map[string]*Transaction)

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
		transactions[transaction.HashID] = &transaction
	}
	return transactions, nil
}

func GetOrderedTransactions() []string {
	orderedtransbyte, err := database.DB.Read(OrderedKey, OrderedTransactions)
	if err != nil {
		return nil
	}
	var orderedtrans []string

	err = database.Decode(orderedtransbyte, orderedtrans)
	if err != nil {
		return nil
	}
	return orderedtrans
}
