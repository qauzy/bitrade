package entity

func (this *ScanMemberAddress) SetId(id int64) {
	this.Id = id
}
func (this *ScanMemberAddress) GetId() (id int64) {
	return this.Id
}
func (this *ScanMemberAddress) SetUnit(unit string) {
	this.Unit = unit
}
func (this *ScanMemberAddress) GetUnit() (unit string) {
	return this.Unit
}
func (this *ScanMemberAddress) SetRemark(remark string) {
	this.Remark = remark
}
func (this *ScanMemberAddress) GetRemark() (remark string) {
	return this.Remark
}
func (this *ScanMemberAddress) SetAddress(address string) {
	this.Address = address
}
func (this *ScanMemberAddress) GetAddress() (address string) {
	return this.Address
}
func ToScanMemberAddress(memberAddress MemberAddress) (result ScanMemberAddress) {
	return ScanMemberAddress.Builder().Id(memberAddress.GetId()).Address(memberAddress.GetAddress()).Remark(memberAddress.GetRemark()).Unit(memberAddress.GetCoin().GetUnit()).Build()
}

type ScanMemberAddress struct {
	Id      int64
	Unit    string
	Remark  string
	Address string
}
