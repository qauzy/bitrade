package LegalWalletState

var (
	APPLYING = LegalWalletState{"申请中", 0}
	COMPLETE = LegalWalletState{"完成", 1}
	DEFEATED = LegalWalletState{"失败", 2}
)

func (this *LegalWalletState) Ordinal() (result int) {
	return this.ordinal
}
func NewLegalWalletState(CnName string) (this *LegalWalletState) {
	this = new(LegalWalletState)
	this.CnName = this.CnName
	return
}

type LegalWalletState struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
