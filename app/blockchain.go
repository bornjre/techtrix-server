package app

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/bornjre/techtrix-server/app/config"
)

func init() {
	print("initilizing blockchain.., ")
	if config.IsFirstRun() {
		InitilizeBlockchainRoot()

	}
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
	concatstr := tf.Subject + tf.Body + string(tf.CreatedAt) + strconv.Itoa(int(tf.Amount)) + tf.User
	print("bytes from body")
	print(concatstr)
	return []byte(concatstr)
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
	h := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	printl("    ")

	printl(h)
	printl("    ")

	return h
}

func (t *Transaction) VerifyHash() bool {
	return t.GenerateHash() == t.HashID
}

func InitilizeBlockchainRoot() {

	print("creating root node")

	transaction := NewTransaction(
		"Root Node",
		"This is root node all nodes of seven kingdom",
		10000,
		defaultUser,
	)
	transaction.HashID = transaction.GenerateHash()

	printl("ROOTNODE:", transaction.HashID)
	err := AddTransaction(transaction)
	print(err)
}

func addBlock(transinfo map[string]string) error {

	// TODO check if ok
	subject := transinfo["subject"]
	body := transinfo["body"]
	amount := transinfo["amount"]
	parenthash := transinfo["parent"]
	amountint, err := strconv.Atoi(amount)
	if err != nil {
		print(err)
		return err
	}

	transaction := NewTransaction(subject, body, int64(amountint), defaultUser)
	transaction.HashID = transaction.GenerateHash()

	//get parent
	_, err = GetTransaction(parenthash)
	if err != nil {
		printl("no parent :( ", err)
		return err
	}

	transaction.ParentHash = parenthash

	err = AddTransaction(transaction)
	if err != nil {
		print(err)
		return err
	}
	printl("NODEADDED:", transaction.HashID)
	return nil
}

func VerifyBlockChain() (bool, error) {
	alltrans, err := GetAllTransactions()
	if err != nil {
		printl("database error")
		return false, err
	}

	for _, value := range alltrans {
		//root node
		if value.HashID == "" {
			if value.VerifyHash() {
				continue
			} else {
				return false, fmt.Errorf(" %s is modified: %v", value.HashID, value)
			}
		}

		_, ok := alltrans[value.ParentHash]
		if !ok {
			return false, fmt.Errorf(" %s 's parent has error : %v", value.HashID, value)
		}
		if value.VerifyHash() {
			continue
		} else {
			return false, fmt.Errorf(" %s is modified: %v", value.HashID, value)
		}

	}

	return true, nil
}
