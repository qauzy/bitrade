package transform

import (
	"bitrade/core/constant/AdvertiseControlStatus"
	"bitrade/core/constant/AdvertiseType"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *MemberAdvertise) SetId(Id int64) (result *MemberAdvertise) {
	this.Id = Id
	return this
}
func (this *MemberAdvertise) GetId() (Id int64) {
	return this.Id
}
func (this *MemberAdvertise) SetAdvertiseType(AdvertiseType AdvertiseType.AdvertiseType) (result *MemberAdvertise) {
	this.AdvertiseType = AdvertiseType
	return this
}
func (this *MemberAdvertise) GetAdvertiseType() (AdvertiseType AdvertiseType.AdvertiseType) {
	return this.AdvertiseType
}
func (this *MemberAdvertise) SetMinLimit(MinLimit math.BigDecimal) (result *MemberAdvertise) {
	this.MinLimit = MinLimit
	return this
}
func (this *MemberAdvertise) GetMinLimit() (MinLimit math.BigDecimal) {
	return this.MinLimit
}
func (this *MemberAdvertise) SetMaxLimit(MaxLimit math.BigDecimal) (result *MemberAdvertise) {
	this.MaxLimit = MaxLimit
	return this
}
func (this *MemberAdvertise) GetMaxLimit() (MaxLimit math.BigDecimal) {
	return this.MaxLimit
}
func (this *MemberAdvertise) SetStatus(Status AdvertiseControlStatus.AdvertiseControlStatus) (result *MemberAdvertise) {
	this.Status = Status
	return this
}
func (this *MemberAdvertise) GetStatus() (Status AdvertiseControlStatus.AdvertiseControlStatus) {
	return this.Status
}
func (this *MemberAdvertise) SetRemainAmount(RemainAmount math.BigDecimal) (result *MemberAdvertise) {
	this.RemainAmount = RemainAmount
	return this
}
func (this *MemberAdvertise) GetRemainAmount() (RemainAmount math.BigDecimal) {
	return this.RemainAmount
}
func (this *MemberAdvertise) SetCoinUnit(CoinUnit string) (result *MemberAdvertise) {
	this.CoinUnit = CoinUnit
	return this
}
func (this *MemberAdvertise) GetCoinUnit() (CoinUnit string) {
	return this.CoinUnit
}
func (this *MemberAdvertise) SetCreateTime(CreateTime xtime.Xtime) (result *MemberAdvertise) {
	this.CreateTime = CreateTime
	return this
}
func (this *MemberAdvertise) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *MemberAdvertise) SetCountry(Country *entity.Country) (result *MemberAdvertise) {
	this.Country = Country
	return this
}
func (this *MemberAdvertise) GetCountry() (Country *entity.Country) {
	return this.Country
}
func ToMemberAdvertise(x *entity.Advertise) (result *MemberAdvertise) {
	return new(MemberAdvertise).SetId(x.GetId()).SetAdvertiseType(x.GetAdvertiseType()).SetCoinUnit(x.GetCoin().GetUnit()).SetCreateTime(x.GetCreateTime()).SetMinLimit(x.GetMinLimit()).SetMaxLimit(x.GetMaxLimit()).SetRemainAmount(x.GetRemainAmount()).SetStatus(x.GetStatus()).SetCountry(x.GetCountry())
}
func NewMemberAdvertise(id int64, advertiseType AdvertiseType.AdvertiseType, minLimit math.BigDecimal, maxLimit math.BigDecimal, status AdvertiseControlStatus.AdvertiseControlStatus, remainAmount math.BigDecimal, coinUnit string, createTime xtime.Xtime, country *entity.Country) (ret *MemberAdvertise) {
	ret = new(MemberAdvertise)
	ret.Id = id
	ret.AdvertiseType = advertiseType
	ret.MinLimit = minLimit
	ret.MaxLimit = maxLimit
	ret.Status = status
	ret.RemainAmount = remainAmount
	ret.CoinUnit = coinUnit
	ret.CreateTime = createTime
	ret.Country = country
	return
}

type MemberAdvertise struct {
	Id            int64                                         `gorm:"column:id" json:"id"`
	AdvertiseType AdvertiseType.AdvertiseType                   `gorm:"column:advertise_type" json:"advertiseType"`
	MinLimit      math.BigDecimal                               `gorm:"column:min_limit" json:"minLimit"`
	MaxLimit      math.BigDecimal                               `gorm:"column:max_limit" json:"maxLimit"`
	Status        AdvertiseControlStatus.AdvertiseControlStatus `gorm:"column:status" json:"status"`
	RemainAmount  math.BigDecimal                               `gorm:"column:remain_amount" json:"remainAmount"`
	CoinUnit      string                                        `gorm:"column:coin_unit" json:"coinUnit"`
	CreateTime    xtime.Xtime                                   `gorm:"column:create_time" json:"createTime"`
	Country       *entity.Country                               `gorm:"column:country" json:"country"`
}
