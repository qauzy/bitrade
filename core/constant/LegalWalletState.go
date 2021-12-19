package constant

type LegalWalletState int

const (
	LEGALWALLETSTATE_APPLYING LegalWalletState = iota
	LEGALWALLETSTATE_COMPLETE
	LEGALWALLETSTATE_DEFEATED
)

func (this LegalWalletState) String() string {
	switch this {
	case LEGALWALLETSTATE_APPLYING:
		return "申请中"
	case LEGALWALLETSTATE_COMPLETE:
		return "完成"
	case LEGALWALLETSTATE_DEFEATED:
		return "失败"
	}
	return ""
}
