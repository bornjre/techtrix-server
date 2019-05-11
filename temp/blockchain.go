package app

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

var ErrBCinit = errors.New("Block already initilized")

type Block struct {
	Parent *Block
	//date,  transactionid
	Data string // TODO convert to proper transaction type
	Hash string
}

func NewBlock(parent *Block, data string) *Block {
	return &Block{
		Parent: parent,
		Data:   data,
	}
}

func (h *Block) GenerateHash() string {

	hasher := sha256.New()
	if h.Parent != nil {
		hasher.Write([]byte(h.Parent.Hash + h.Data))
	} else {
		hasher.Write([]byte(h.Data))
	}

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (h *Block) VerifyHash() bool {
	return h.GenerateHash() == h.Hash
}

func (h *Block) template() {
}

type BlockChain struct {
	// TODO mutex
	// lock sync.mutex
	Hashstore map[string]*Block
	//TODO helper structure of for easy parsing data
	_hashindexes []string
	_dataindexes []string
	_dirty       bool
	// TODO end

	hashlist []string
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		Hashstore: make(map[string]*Block),
	}
}

func (c *BlockChain) getLatestBlock() *Block {
	fmt.Printf("%+v", c.hashlist)
	_hashid := c.hashlist[len(c.hashlist)-1]
	printl(len(c.hashlist))
	block, ok := c.Hashstore[_hashid]
	if !ok {
		ErrorHappened("No latest block to get")
	}
	return block
}

func (c *BlockChain) initBlock(data string) error {
	// TODO init data type depend upon on our tansaction type

	if len(c.Hashstore) != 0 && len(c.hashlist) != 0 {
		printl("???")
		return ErrBCinit
	}

	nb := NewBlock(nil, data)
	if nb == nil {
		ErrorHappened("could not init")
	}
	nb.Hash = nb.GenerateHash()

	c.hashlist = append(c.hashlist, nb.Hash)
	c.Hashstore[nb.Hash] = nb
	return nil

}

func (c *BlockChain) addBlock(data string) *Block {
	oldblock := c.getLatestBlock()
	newblock := NewBlock(oldblock, data)

	newblock.Hash = newblock.GenerateHash()

	c.hashlist = append(c.hashlist, newblock.Hash)
	c.Hashstore[newblock.Hash] = newblock
	return newblock

}

func (c *BlockChain) VerifyBlockChain() bool {

	for key, index := range c.hashlist {

		block, ok := c.Hashstore[index]
		if !ok {
			printl("Error at key:", key)
			return false
		}

		if !block.VerifyHash() {
			return false
		}

	}
	return true
}

func (c *BlockChain) printBlockChain() {

	b, err := json.MarshalIndent(c.Hashstore, "", "    ")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v", string(b))
}

func (c *BlockChain) getBlock(hashid string) *Block {
	b, ok := c.Hashstore[hashid]
	if ok {
		return b
	}
	return nil
}

func (c *BlockChain) generateCache() {
}

func (c *BlockChain) template() {
}
