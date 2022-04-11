package entity

func (this *ScanMemberAddress) SetId(Id int64) (result *ScanMemberAddress) {
	this.Id = Id
	return this
}
func (this *ScanMemberAddress) GetId() (Id int64) {
	return this.Id
}
func (this *ScanMemberAddress) SetUnit(Unit string) (result *ScanMemberAddress) {
	this.Unit = Unit
	return this
}
func (this *ScanMemberAddress) GetUnit() (Unit string) {
	return this.Unit
}
func (this *ScanMemberAddress) SetRemark(Remark string) (result *ScanMemberAddress) {
	this.Remark = Remark
	return this
}
func (this *ScanMemberAddress) GetRemark() (Remark string) {
	return this.Remark
}
func (this *ScanMemberAddress) SetAddress(Address string) (result *ScanMemberAddress) {
	this.Address = Address
	return this
}
func (this *ScanMemberAddress) GetAddress() (Address string) {
	return this.Address
}
func ToScanMemberAddress(memberAddress *MemberAddress) (result *ScanMemberAddress) {
	return new(ScanMemberAddress).SetId(memberAddress.GetId()).SetAddress(memberAddress.GetAddress()).SetRemark(memberAddress.GetRemark()).SetUnit(memberAddress.GetCoin().GetUnit())
}
func NewScanMemberAddress(id int64, unit string, remark string, address string) (ret *ScanMemberAddress) {
	ret = new(ScanMemberAddress)
	ret.Id = id
	ret.Unit = unit
	ret.Remark = remark
	ret.Address = address
	return
}

type ScanMemberAddress struct {
	Id      int64  `gorm:"column:id" json:"id"`
	Unit    string `gorm:"column:unit" json:"unit"`
	Remark  string `gorm:"column:remark" json:"remark"`
	Address string `gorm:"column:address" json:"address"`
}
