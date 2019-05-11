package app

import (
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)

func init() {
	InitilizeBlockchainRoot()
}

type Transaction struct {
	HashID          string          `json:hash_id`
	TransactionInfo TransactionInfo `json:transaction_info`
	ParentHash      string          `json:parenthash`
}

type TransactionInfo struct {
	Subject   string `json:subject`
	Body      string `json:body`
	CreatedAt int64  `json:created_at`
	Amount    int64  `json:amount`
	User      string `json:user`
}

func (tf *TransactionInfo) getBytes() []byte {
	return []byte(tf.Subject + tf.Body + string(tf.CreatedAt) + string(tf.Amount) + tf.User)
}

func NewTransaction(subject string, body string, amount int64, user string) *Transaction {

	return &Transaction{
		TransactionInfo: TransactionInfo{

			Subject:   subject,
			Body:      body,
			User:      user,
			CreatedAt: time.Now().Unix(),
			Amount:    amount,
		}}
}

func (t *Transaction) GenerateHash() string {
	hasher := sha256.New()
	if t.ParentHash == "" {
		// Means ROOT NODE
		hasher.Write(t.TransactionInfo.getBytes())
	} else {
		w := append(t.TransactionInfo.getBytes(), []byte(t.ParentHash)...)
		hasher.Write(w)
	}

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (t *Transaction) VerifyHash() bool {
	return t.GenerateHash() == t.HashID
}

func InitilizeBlockchainRoot() {

	transaction := NewTransaction(
		"Root Node",
		"This is root node all nodes of seven kingdom",
		10000,
		defaultUser,
	)
	AddTransaction(transaction)

}

func addBlock(transinfo map[string]string) {

	// TODO check if ok
	subject := transinfo["subject"]
	body := transinfo["body"]
	amount := transinfo["amount"]
	amountint, err := strconv.Atoi(amount)
	if err != nil {
		print(err)
		return
	}

	transaction := NewTransaction(subject, body, int64(amountint), defaultUser)
	transaction.HashID = transaction.GenerateHash()
	AddTransaction(transaction)

}

func VerifyBlockChain() bool {
	alltrans, err := GetAllTransactions()
	if err != nil {
		print("database error")
		return false
	}

	for _, value := range alltrans {
		//root node
		if value.HashID == "" {
			return value.VerifyHash()
		}

		_, ok := alltrans[value.ParentHash]
		if !ok {
			return false
		}

		return value.VerifyHash()

	}

	return false
}
