package entity

func (this *ScanMemberAddress) SetId(id int64) (result *ScanMemberAddress) {
	this.Id = id
	return this
}
func (this *ScanMemberAddress) GetId() (id int64) {
	return this.Id
}
func (this *ScanMemberAddress) SetUnit(unit string) (result *ScanMemberAddress) {
	this.Unit = unit
	return this
}
func (this *ScanMemberAddress) GetUnit() (unit string) {
	return this.Unit
}
func (this *ScanMemberAddress) SetRemark(remark string) (result *ScanMemberAddress) {
	this.Remark = remark
	return this
}
func (this *ScanMemberAddress) GetRemark() (remark string) {
	return this.Remark
}
func (this *ScanMemberAddress) SetAddress(address string) (result *ScanMemberAddress) {
	this.Address = address
	return this
}
func (this *ScanMemberAddress) GetAddress() (address string) {
	return this.Address
}
func ToScanMemberAddress(memberAddress *MemberAddress) (result *ScanMemberAddress) {
	return new(ScanMemberAddress).SetId(memberAddress.GetId()).SetAddress(memberAddress.GetAddress()).SetRemark(memberAddress.GetRemark()).SetUnit(memberAddress.GetCoin().GetUnit())
}

type ScanMemberAddress struct {
	Id      int64
	Unit    string
	Remark  string
	Address string
}
