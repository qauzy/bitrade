package LegalWalletState

type LegalWalletState int

const (
	APPLYING LegalWalletState = iota
	COMPLETE
	DEFEATED
)

func (this LegalWalletState) String() string {
	switch this {
	case APPLYING:
		return "申请中"
	case COMPLETE:
		return "完成"
	case DEFEATED:
		return "失败"
	}
	return ""
}
