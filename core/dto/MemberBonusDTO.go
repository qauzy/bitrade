package dto

import "github.com/qauzy/math"

func (this *MemberBonusDTO) SetId(Id int64) (result *MemberBonusDTO) {
	this.Id = Id
	return this
}
func (this *MemberBonusDTO) GetId() (Id int64) {
	return this.Id
}
func (this *MemberBonusDTO) SetMemberId(MemberId int64) (result *MemberBonusDTO) {
	this.MemberId = MemberId
	return this
}
func (this *MemberBonusDTO) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *MemberBonusDTO) SetHaveTime(HaveTime string) (result *MemberBonusDTO) {
	this.HaveTime = HaveTime
	return this
}
func (this *MemberBonusDTO) GetHaveTime() (HaveTime string) {
	return this.HaveTime
}
func (this *MemberBonusDTO) SetArriveTime(ArriveTime string) (result *MemberBonusDTO) {
	this.ArriveTime = ArriveTime
	return this
}
func (this *MemberBonusDTO) GetArriveTime() (ArriveTime string) {
	return this.ArriveTime
}
func (this *MemberBonusDTO) SetTotal(Total *math.BigDecimal) (result *MemberBonusDTO) {
	this.Total = Total
	return this
}
func (this *MemberBonusDTO) GetTotal() (Total *math.BigDecimal) {
	return this.Total
}
func (this *MemberBonusDTO) SetMemBouns(MemBouns *math.BigDecimal) (result *MemberBonusDTO) {
	this.MemBouns = MemBouns
	return this
}
func (this *MemberBonusDTO) GetMemBouns() (MemBouns *math.BigDecimal) {
	return this.MemBouns
}
func (this *MemberBonusDTO) SetCoinId(CoinId string) (result *MemberBonusDTO) {
	this.CoinId = CoinId
	return this
}
func (this *MemberBonusDTO) GetCoinId() (CoinId string) {
	return this.CoinId
}

type MemberBonusDTO struct {
	Id         int64            `gorm:"column:id" json:"id"`
	MemberId   int64            `gorm:"column:member_id" json:"memberId"`
	HaveTime   string           `gorm:"column:have_time" json:"haveTime"`
	ArriveTime string           `gorm:"column:arrive_time" json:"arriveTime"`
	Total      *math.BigDecimal `gorm:"column:total" json:"total"`
	MemBouns   *math.BigDecimal `gorm:"column:mem_bouns" json:"memBouns"`
	CoinId     string           `gorm:"column:coin_id" json:"coinId"`
}
