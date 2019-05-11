package database

type Transaction struct {
	HashID          string          `json:hash_id`
	TransactionInfo TransactionInfo `json:transaction_info`
}

type TransactionInfo struct {
	Subject   string `json:subject`
	Body      string `json:body`
	CreatedAt int64  `json:created_at`
	Amount    int64  `json:amount`
	User      string `json:user`
}

var (
	TransactionBucketName = []byte("transactions")
)

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
