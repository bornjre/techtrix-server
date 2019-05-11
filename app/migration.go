package app

func InitilizeBlockchainRoot() {

	print("creating root node")

	rootTran := NewTransaction(
		"Bharatpur Metropolitan",
		"total budget amount allocated for bharatpur metropolitan city",
		10000000,
		defaultUser,
	)
	rootTran.HashID = rootTran.GenerateHash()
	printl("ROOTNODE:", rootTran.HashID)
	print(AddTransaction(rootTran))

	child1 := NewTransaction(
		"General Infrastructure",
		"Budget allocated for various infrastructure",
		3000000,
		defaultUser,
	)
	child1.ParentHash = rootTran.HashID
	child1.HashID = child1.GenerateHash()
	print(AddTransaction(child1))

	child2 := NewTransaction(
		"Education",
		"Budget allocated for Education sector",
		100000,
		defaultUser,
	)
	child2.ParentHash = rootTran.HashID
	child2.HashID = child2.GenerateHash()
	print(AddTransaction(child2))

	grandc1 := NewTransaction(
		"Employee salary",
		"Employee salary under education office",
		100000,
		defaultUser,
	)
	grandc1.ParentHash = child1.HashID
	grandc1.HashID = grandc1.GenerateHash()
	print(AddTransaction(grandc1))

	grandc2 := NewTransaction(
		"Remuneration",
		"Remuneration..",
		100000,
		defaultUser,
	)
	grandc2.ParentHash = child1.HashID
	grandc2.HashID = grandc2.GenerateHash()
	print(AddTransaction(grandc2))

}
