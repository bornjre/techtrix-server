package app

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/bornjre/techtrix-server/app/database"

	"github.com/bornjre/techtrix-server/app/config"
)

var subscribersService *subscribers

func init() {
	print("initilizing blockchain.., ")
	if config.IsFirstRun() {
		InitilizeBlockchainRoot()

	}
	subscribersService = NewSubscriberService()
	subscribersService.RunService()
}

type Transaction struct {
	HashID          string           `json:hash_id`
	TransactionInfo *TransactionInfo `json:transaction_info`
	ParentHash      string           `json:parenthash`
}

type TransactionInfo struct {
	Subject   string `json:subject`
	Body      string `json:body`
	CreatedAt int64  `json:created_at`
	Amount    int64  `json:amount`
	User      string `json:user`
}

func (tf *TransactionInfo) getBytes() []byte {
	b, err := json.Marshal(tf)
	if err != nil {
		printl("MARSEL ERR", err)
	}
	return b
}

func NewTransaction(subject string, body string, amount int64, user string) *Transaction {

	return &Transaction{
		TransactionInfo: &TransactionInfo{

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
	printl(t.GenerateHash(), ">>", t.HashID)
	return t.GenerateHash() == t.HashID
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
	//get parent
	_, err = GetTransaction(parenthash)
	if err != nil {
		printl("no parent :( ", err)
		return err
	}

	transaction.ParentHash = parenthash
	transaction.HashID = transaction.GenerateHash()

	err = AddTransaction(transaction)
	if err != nil {
		print(err)
		return err
	}
	printl("NODEADDED:", transaction.HashID)
	subscribersService.BroadcastChan <- fmt.Sprintf("%s : %s", transaction.HashID, transaction.ParentHash)
	return nil
}

func updateBlock(transinfo map[string]string) {
	hashid, ok := transinfo["hashid"]
	if !ok {
		return
	}

	trans, err := GetTransaction(hashid)
	if err != nil {
		printl(err)
		return
	}
	b, err := database.Encode(transinfo)
	if err != nil {
		printl(err)
		return
	}
	t := &TransactionInfo{}
	err = database.Decode(b, t)
	if err != nil {
		printl(err)
		return
	}
	trans.TransactionInfo = t
	//AddTransaction(trans)
	transbyte, err := database.Encode(trans)
	if err != nil {
		printl(err)
		return
	}

	err = database.DB.Update([]byte(trans.HashID), transbyte, TransactionBucketName)

	if err != nil {
		printl(err)
		return
	}

}

func VerifyBlockChain() (bool, error) {
	alltrans, err := GetAllTransactions()
	if err != nil {
		printl("database error")
		return false, err
	}

	for _, value := range alltrans {
		//root node
		printl("HERE")
		fmt.Printf("%+v", value)
		fmt.Printf("%+v", value.TransactionInfo)
		printl("HERE")
		fmt.Println("parent:", value.ParentHash)
		printl("==============")
		if value.ParentHash == "" {
			printl("INSIDE>>>>>>")
			if value.VerifyHash() {
				printl("one passed")
				continue
			} else {
				return false, fmt.Errorf(" %s is modified: %v", value.HashID, value)
			}
		} else {

			_, ok := alltrans[value.ParentHash]
			if !ok {
				return false, fmt.Errorf(" %s 's parent has error : %v", value.HashID, value)
			}
			if value.VerifyHash() {
				printl("one passed 2")
				continue
			} else {
				return false, fmt.Errorf(" %s is modified: %v", value.HashID, value)
			}

		}
	}

	return true, nil
}
