package app

import "fmt"

func Run() {

	bc := NewBlockChain()

	bc.initBlock("Mother of all blocks")
	child1 := bc.addBlock("Children1")
	bc.addBlock("grandchild")
	bc.printBlockChain()
	fmt.Println("Checking blockchain..")
	fmt.Println(bc.VerifyBlockChain())

	fmt.Println("Intermission")
	child1.Hash = "ASASBBDBDBBDBBD"

	fmt.Println("Is block still correct ?")
	print(bc.VerifyBlockChain())

}
