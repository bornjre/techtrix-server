package database

type Transaction struct {
	HashID          string          `json:hash_id`
	TransactionInfo TransactionInfo `json:transaction_info`
}

type TransactionInfo struct {
	subject    string `json:subject`
	body       string `json:body`
	created_at int64  `json:created_at`
}

var (
	TransactionBucketName = []byte("transactions")
)

func AddTransaction(trans *Transaction) error {

	transbyte, err := Encode(traans)
	if err != nil {
		return err
	}

	err = DB.Create([]byte(trans.HashID), transbyte, TransactionBucketName)
	if err != nil {
		return err
	}
	return nil
}

func GetTransaction(hash_id string) (*Transaction, error) {
	transaction := &Transaction{}
	transbyte, err := DB.Read([]byte(hash_id), TransactionBucketName)

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
