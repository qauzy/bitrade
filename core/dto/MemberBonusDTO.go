package dto

func (this *MemberBonusDTO) SetId(id int64) (result *MemberBonusDTO) {
	this.Id = id
	return this
}
func (this *MemberBonusDTO) GetId() (id int64) {
	return this.Id
}
func (this *MemberBonusDTO) SetMemberId(memberId int64) (result *MemberBonusDTO) {
	this.MemberId = memberId
	return this
}
func (this *MemberBonusDTO) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberBonusDTO) SetHaveTime(haveTime string) (result *MemberBonusDTO) {
	this.HaveTime = haveTime
	return this
}
func (this *MemberBonusDTO) GetHaveTime() (haveTime string) {
	return this.HaveTime
}
func (this *MemberBonusDTO) SetArriveTime(arriveTime string) (result *MemberBonusDTO) {
	this.ArriveTime = arriveTime
	return this
}
func (this *MemberBonusDTO) GetArriveTime() (arriveTime string) {
	return this.ArriveTime
}
func (this *MemberBonusDTO) SetTotal(total math.BigDecimal) (result *MemberBonusDTO) {
	this.Total = total
	return this
}
func (this *MemberBonusDTO) GetTotal() (total math.BigDecimal) {
	return this.Total
}
func (this *MemberBonusDTO) SetMemBouns(memBouns math.BigDecimal) (result *MemberBonusDTO) {
	this.MemBouns = memBouns
	return this
}
func (this *MemberBonusDTO) GetMemBouns() (memBouns math.BigDecimal) {
	return this.MemBouns
}
func (this *MemberBonusDTO) SetCoinId(coinId string) (result *MemberBonusDTO) {
	this.CoinId = coinId
	return this
}
func (this *MemberBonusDTO) GetCoinId() (coinId string) {
	return this.CoinId
}

type MemberBonusDTO struct {
	Id         int64
	MemberId   int64
	HaveTime   string
	ArriveTime string
	Total      math.BigDecimal
	MemBouns   math.BigDecimal
	CoinId     string
}
