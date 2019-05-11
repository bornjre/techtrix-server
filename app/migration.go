package app

func InitilizeBlockchainRoot() {

	print("creating root node")

	rootTran := NewTransaction(
		"Root Node",
		"This is root node all nodes of seven kingdom",
		10000,
		defaultUser,
	)
	rootTran.HashID = rootTran.GenerateHash()

	printl("ROOTNODE:", rootTran.HashID)
	print(AddTransaction(rootTran))

	child1 := NewTransaction(
		"child one",
		" This is child one",
		1000,
		defaultUser,
	)
	child1.ParentHash = rootTran.HashID
	child1.HashID = child1.GenerateHash()
	print(AddTransaction(child1))

	child2 := NewTransaction(
		"child two",
		" This is child two",
		1000,
		defaultUser,
	)
	child2.ParentHash = rootTran.HashID
	child2.HashID = child2.GenerateHash()
	print(AddTransaction(child2))

	grandc1 := NewTransaction(
		"grandchild one",
		" This is gchild one",
		1000,
		defaultUser,
	)
	grandc1.ParentHash = child1.HashID
	grandc1.HashID = grandc1.GenerateHash()
	print(AddTransaction(grandc1))

}
