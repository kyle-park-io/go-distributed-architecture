package node

const (
	// action
	SaveMsg = iota
	ReturnMsg
	SaveTx
	ReturnTx
	VerifyTx
	// result
	Success
	Fail
	ValidTx
	InvalidTx
	// admin
	ADMIN_NODE = 9999
)
