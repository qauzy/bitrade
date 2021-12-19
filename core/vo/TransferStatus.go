package vo

type TransferStatus int

const (
	TRANSFERSTATUS_Recharge TransferStatus = iota
	TRANSFERSTATUS_Withdraw
	TRANSFERSTATUS_BUY
	TRANSFERSTATUS_SELL
)

func (this TransferStatus) String() string {
	switch this {
	}
	return ""
}
