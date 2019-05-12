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
	return

}
