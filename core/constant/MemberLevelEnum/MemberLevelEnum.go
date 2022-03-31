package MemberLevelEnum

var (
	GENERAL        = MemberLevelEnum{"普通", 0}
	REALNAME       = MemberLevelEnum{"实名", 1}
	IDENTIFICATION = MemberLevelEnum{"认证商家", 2}
)

func (this *MemberLevelEnum) Ordinal() (result int) {
	return this.ordinal
}
func (this *MemberLevelEnum) GetOrdinal() (result int) {
	return this.Ordinal()
}

type MemberLevelEnum struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
