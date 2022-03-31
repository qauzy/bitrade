package MemberRegisterType

var (
	BACKSTAGE              = MemberRegisterType{"后台", 0}
	AUTONOMOUSLY           = MemberRegisterType{"自主", 1}
	AUTONOMOUSLY_RECOMMEND = MemberRegisterType{"自主(推荐)", 2}
)

func (this *MemberRegisterType) Ordinal() (result int) {
	return this.ordinal
}
func (this *MemberRegisterType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type MemberRegisterType struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
