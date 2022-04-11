package transform

import (
	"bitrade/core/constant/AdvertiseType"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *ScanAdvertise) SetMemberName(MemberName string) (result *ScanAdvertise) {
	this.MemberName = MemberName
	return this
}
func (this *ScanAdvertise) GetMemberName() (MemberName string) {
	return this.MemberName
}
func (this *ScanAdvertise) SetAvatar(Avatar string) (result *ScanAdvertise) {
	this.Avatar = Avatar
	return this
}
func (this *ScanAdvertise) GetAvatar() (Avatar string) {
	return this.Avatar
}
func (this *ScanAdvertise) SetAdvertiseId(AdvertiseId int64) (result *ScanAdvertise) {
	this.AdvertiseId = AdvertiseId
	return this
}
func (this *ScanAdvertise) GetAdvertiseId() (AdvertiseId int64) {
	return this.AdvertiseId
}
func (this *ScanAdvertise) SetTransactions(Transactions int) (result *ScanAdvertise) {
	this.Transactions = Transactions
	return this
}
func (this *ScanAdvertise) GetTransactions() (Transactions int) {
	return this.Transactions
}
func (this *ScanAdvertise) SetPrice(Price math.BigDecimal) (result *ScanAdvertise) {
	this.Price = Price
	return this
}
func (this *ScanAdvertise) GetPrice() (Price math.BigDecimal) {
	return this.Price
}
func (this *ScanAdvertise) SetMinLimit(MinLimit math.BigDecimal) (result *ScanAdvertise) {
	this.MinLimit = MinLimit
	return this
}
func (this *ScanAdvertise) GetMinLimit() (MinLimit math.BigDecimal) {
	return this.MinLimit
}
func (this *ScanAdvertise) SetMaxLimit(MaxLimit math.BigDecimal) (result *ScanAdvertise) {
	this.MaxLimit = MaxLimit
	return this
}
func (this *ScanAdvertise) GetMaxLimit() (MaxLimit math.BigDecimal) {
	return this.MaxLimit
}
func (this *ScanAdvertise) SetRemainAmount(RemainAmount math.BigDecimal) (result *ScanAdvertise) {
	this.RemainAmount = RemainAmount
	return this
}
func (this *ScanAdvertise) GetRemainAmount() (RemainAmount math.BigDecimal) {
	return this.RemainAmount
}
func (this *ScanAdvertise) SetCreateTime(CreateTime xtime.Xtime) (result *ScanAdvertise) {
	this.CreateTime = CreateTime
	return this
}
func (this *ScanAdvertise) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *ScanAdvertise) SetPayMode(PayMode string) (result *ScanAdvertise) {
	this.PayMode = PayMode
	return this
}
func (this *ScanAdvertise) GetPayMode() (PayMode string) {
	return this.PayMode
}
func (this *ScanAdvertise) SetCoinId(CoinId int64) (result *ScanAdvertise) {
	this.CoinId = CoinId
	return this
}
func (this *ScanAdvertise) GetCoinId() (CoinId int64) {
	return this.CoinId
}
func (this *ScanAdvertise) SetUnit(Unit string) (result *ScanAdvertise) {
	this.Unit = Unit
	return this
}
func (this *ScanAdvertise) GetUnit() (Unit string) {
	return this.Unit
}
func (this *ScanAdvertise) SetCoinName(CoinName string) (result *ScanAdvertise) {
	this.CoinName = CoinName
	return this
}
func (this *ScanAdvertise) GetCoinName() (CoinName string) {
	return this.CoinName
}
func (this *ScanAdvertise) SetCoinNameCn(CoinNameCn string) (result *ScanAdvertise) {
	this.CoinNameCn = CoinNameCn
	return this
}
func (this *ScanAdvertise) GetCoinNameCn() (CoinNameCn string) {
	return this.CoinNameCn
}
func (this *ScanAdvertise) SetLevel(Level int) (result *ScanAdvertise) {
	this.Level = Level
	return this
}
func (this *ScanAdvertise) GetLevel() (Level int) {
	return this.Level
}
func (this *ScanAdvertise) SetAdvertiseType(AdvertiseType AdvertiseType.AdvertiseType) (result *ScanAdvertise) {
	this.AdvertiseType = AdvertiseType
	return this
}
func (this *ScanAdvertise) GetAdvertiseType() (AdvertiseType AdvertiseType.AdvertiseType) {
	return this.AdvertiseType
}
func (this *ScanAdvertise) SetAdvType(AdvType int) (result *ScanAdvertise) {
	this.AdvType = AdvType
	return this
}
func (this *ScanAdvertise) GetAdvType() (AdvType int) {
	return this.AdvType
}
func (this *ScanAdvertise) SetPremiseRate(PremiseRate math.BigDecimal) (result *ScanAdvertise) {
	this.PremiseRate = PremiseRate
	return this
}
func (this *ScanAdvertise) GetPremiseRate() (PremiseRate math.BigDecimal) {
	return this.PremiseRate
}
func NewScanAdvertise(memberName string, avatar string, advertiseId int64, transactions int, price math.BigDecimal, minLimit math.BigDecimal, maxLimit math.BigDecimal, remainAmount math.BigDecimal, createTime xtime.Xtime, payMode string, coinId int64, unit string, coinName string, coinNameCn string, level int, advertiseType AdvertiseType.AdvertiseType, advType int, premiseRate math.BigDecimal) (ret *ScanAdvertise) {
	ret = new(ScanAdvertise)
	ret.MemberName = memberName
	ret.Avatar = avatar
	ret.AdvertiseId = advertiseId
	ret.Transactions = transactions
	ret.Price = price
	ret.MinLimit = minLimit
	ret.MaxLimit = maxLimit
	ret.RemainAmount = remainAmount
	ret.CreateTime = createTime
	ret.PayMode = payMode
	ret.CoinId = coinId
	ret.Unit = unit
	ret.CoinName = coinName
	ret.CoinNameCn = coinNameCn
	ret.Level = level
	ret.AdvertiseType = advertiseType
	ret.AdvType = advType
	ret.PremiseRate = premiseRate
	return
}

type ScanAdvertise struct {
	MemberName    string                      `gorm:"column:member_name" json:"memberName"`
	Avatar        string                      `gorm:"column:avatar" json:"avatar"`
	AdvertiseId   int64                       `gorm:"column:advertise_id" json:"advertiseId"`
	Transactions  int                         `gorm:"column:transactions" json:"transactions"`
	Price         math.BigDecimal             `gorm:"column:price" json:"price"`
	MinLimit      math.BigDecimal             `gorm:"column:min_limit" json:"minLimit"`
	MaxLimit      math.BigDecimal             `gorm:"column:max_limit" json:"maxLimit"`
	RemainAmount  math.BigDecimal             `gorm:"column:remain_amount" json:"remainAmount"`
	CreateTime    xtime.Xtime                 `gorm:"column:create_time" json:"createTime"`
	PayMode       string                      `gorm:"column:pay_mode" json:"payMode"`
	CoinId        int64                       `gorm:"column:coin_id" json:"coinId"`
	Unit          string                      `gorm:"column:unit" json:"unit"`
	CoinName      string                      `gorm:"column:coin_name" json:"coinName"`
	CoinNameCn    string                      `gorm:"column:coin_name_cn" json:"coinNameCn"`
	Level         int                         `gorm:"column:level" json:"level"`
	AdvertiseType AdvertiseType.AdvertiseType `gorm:"column:advertise_type" json:"advertiseType"`
	AdvType       int                         `gorm:"column:adv_type" json:"advType"`
	PremiseRate   math.BigDecimal             `gorm:"column:premise_rate" json:"premiseRate"`
}
