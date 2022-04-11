package entity

import (
	"bitrade/core/constant/AdvertiseControlStatus"
	"bitrade/core/constant/AdvertiseLevel"
	"bitrade/core/constant/AdvertiseType"
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/PriceType"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *Advertise) SetId(Id int64) (result *Advertise) {
	this.Id = Id
	return this
}
func (this *Advertise) GetId() (Id int64) {
	return this.Id
}
func (this *Advertise) SetPrice(Price math.BigDecimal) (result *Advertise) {
	this.Price = Price
	return this
}
func (this *Advertise) GetPrice() (Price math.BigDecimal) {
	return this.Price
}
func (this *Advertise) SetAdvertiseType(AdvertiseType AdvertiseType.AdvertiseType) (result *Advertise) {
	this.AdvertiseType = AdvertiseType
	return this
}
func (this *Advertise) GetAdvertiseType() (AdvertiseType AdvertiseType.AdvertiseType) {
	return this.AdvertiseType
}
func (this *Advertise) SetMember(Member *Member) (result *Advertise) {
	this.Member = Member
	return this
}
func (this *Advertise) GetMember() (Member *Member) {
	return this.Member
}
func (this *Advertise) SetCreateTime(CreateTime xtime.Xtime) (result *Advertise) {
	this.CreateTime = CreateTime
	return this
}
func (this *Advertise) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *Advertise) SetUpdateTime(UpdateTime xtime.Xtime) (result *Advertise) {
	this.UpdateTime = UpdateTime
	return this
}
func (this *Advertise) GetUpdateTime() (UpdateTime xtime.Xtime) {
	return this.UpdateTime
}
func (this *Advertise) SetCoin(Coin *OtcCoin) (result *Advertise) {
	this.Coin = Coin
	return this
}
func (this *Advertise) GetCoin() (Coin *OtcCoin) {
	return this.Coin
}
func (this *Advertise) SetCountry(Country *Country) (result *Advertise) {
	this.Country = Country
	return this
}
func (this *Advertise) GetCountry() (Country *Country) {
	return this.Country
}
func (this *Advertise) SetMinLimit(MinLimit math.BigDecimal) (result *Advertise) {
	this.MinLimit = MinLimit
	return this
}
func (this *Advertise) GetMinLimit() (MinLimit math.BigDecimal) {
	return this.MinLimit
}
func (this *Advertise) SetMaxLimit(MaxLimit math.BigDecimal) (result *Advertise) {
	this.MaxLimit = MaxLimit
	return this
}
func (this *Advertise) GetMaxLimit() (MaxLimit math.BigDecimal) {
	return this.MaxLimit
}
func (this *Advertise) SetRemark(Remark string) (result *Advertise) {
	this.Remark = Remark
	return this
}
func (this *Advertise) GetRemark() (Remark string) {
	return this.Remark
}
func (this *Advertise) SetTimeLimit(TimeLimit int) (result *Advertise) {
	this.TimeLimit = TimeLimit
	return this
}
func (this *Advertise) GetTimeLimit() (TimeLimit int) {
	return this.TimeLimit
}
func (this *Advertise) SetPremiseRate(PremiseRate math.BigDecimal) (result *Advertise) {
	this.PremiseRate = PremiseRate
	return this
}
func (this *Advertise) GetPremiseRate() (PremiseRate math.BigDecimal) {
	return this.PremiseRate
}
func (this *Advertise) SetLevel(Level AdvertiseLevel.AdvertiseLevel) (result *Advertise) {
	this.Level = Level
	return this
}
func (this *Advertise) GetLevel() (Level AdvertiseLevel.AdvertiseLevel) {
	return this.Level
}
func (this *Advertise) SetPayMode(PayMode string) (result *Advertise) {
	this.PayMode = PayMode
	return this
}
func (this *Advertise) GetPayMode() (PayMode string) {
	return this.PayMode
}
func (this *Advertise) SetVersion(Version int64) (result *Advertise) {
	this.Version = Version
	return this
}
func (this *Advertise) GetVersion() (Version int64) {
	return this.Version
}
func (this *Advertise) SetStatus(Status AdvertiseControlStatus.AdvertiseControlStatus) (result *Advertise) {
	this.Status = Status
	return this
}
func (this *Advertise) GetStatus() (Status AdvertiseControlStatus.AdvertiseControlStatus) {
	return this.Status
}
func (this *Advertise) SetPriceType(PriceType PriceType.PriceType) (result *Advertise) {
	this.PriceType = PriceType
	return this
}
func (this *Advertise) GetPriceType() (PriceType PriceType.PriceType) {
	return this.PriceType
}
func (this *Advertise) SetNumber(Number math.BigDecimal) (result *Advertise) {
	this.Number = Number
	return this
}
func (this *Advertise) GetNumber() (Number math.BigDecimal) {
	return this.Number
}
func (this *Advertise) SetDealAmount(DealAmount math.BigDecimal) (result *Advertise) {
	this.DealAmount = DealAmount
	return this
}
func (this *Advertise) GetDealAmount() (DealAmount math.BigDecimal) {
	return this.DealAmount
}
func (this *Advertise) SetRemainAmount(RemainAmount math.BigDecimal) (result *Advertise) {
	this.RemainAmount = RemainAmount
	return this
}
func (this *Advertise) GetRemainAmount() (RemainAmount math.BigDecimal) {
	return this.RemainAmount
}
func (this *Advertise) SetAuto(Auto BooleanEnum.BooleanEnum) (result *Advertise) {
	this.Auto = Auto
	return this
}
func (this *Advertise) GetAuto() (Auto BooleanEnum.BooleanEnum) {
	return this.Auto
}
func (this *Advertise) SetMarketPrice(MarketPrice math.BigDecimal) (result *Advertise) {
	this.MarketPrice = MarketPrice
	return this
}
func (this *Advertise) GetMarketPrice() (MarketPrice math.BigDecimal) {
	return this.MarketPrice
}
func (this *Advertise) SetAutoword(Autoword string) (result *Advertise) {
	this.Autoword = Autoword
	return this
}
func (this *Advertise) GetAutoword() (Autoword string) {
	return this.Autoword
}
func (this *Advertise) SetLimitMoney(LimitMoney string) (result *Advertise) {
	this.LimitMoney = LimitMoney
	return this
}
func (this *Advertise) GetLimitMoney() (LimitMoney string) {
	return this.LimitMoney
}
func (this *Advertise) SetUsername(Username string) (result *Advertise) {
	this.Username = Username
	return this
}
func (this *Advertise) GetUsername() (Username string) {
	return this.Username
}
func (this *Advertise) SetCoinUnit(CoinUnit string) (result *Advertise) {
	this.CoinUnit = CoinUnit
	return this
}
func (this *Advertise) GetCoinUnit() (CoinUnit string) {
	return this.CoinUnit
}
func NewAdvertise(id int64, price math.BigDecimal, advertiseType AdvertiseType.AdvertiseType, member *Member, createTime xtime.Xtime, updateTime xtime.Xtime, coin *OtcCoin, country *Country, minLimit math.BigDecimal, maxLimit math.BigDecimal, remark string, timeLimit int, premiseRate math.BigDecimal, level AdvertiseLevel.AdvertiseLevel, payMode string, version int64, status AdvertiseControlStatus.AdvertiseControlStatus, priceType PriceType.PriceType, number math.BigDecimal, dealAmount math.BigDecimal, remainAmount math.BigDecimal, auto BooleanEnum.BooleanEnum, marketPrice math.BigDecimal, autoword string, limitMoney string, username string, coinUnit string) (ret *Advertise) {
	ret = new(Advertise)
	ret.Id = id
	ret.Price = price
	ret.AdvertiseType = advertiseType
	ret.Member = member
	ret.CreateTime = createTime
	ret.UpdateTime = updateTime
	ret.Coin = coin
	ret.Country = country
	ret.MinLimit = minLimit
	ret.MaxLimit = maxLimit
	ret.Remark = remark
	ret.TimeLimit = timeLimit
	ret.PremiseRate = premiseRate
	ret.Level = level
	ret.PayMode = payMode
	ret.Version = version
	ret.Status = status
	ret.PriceType = priceType
	ret.Number = number
	ret.DealAmount = dealAmount
	ret.RemainAmount = remainAmount
	ret.Auto = auto
	ret.MarketPrice = marketPrice
	ret.Autoword = autoword
	ret.LimitMoney = limitMoney
	ret.Username = username
	ret.CoinUnit = coinUnit
	return
}

type Advertise struct {
	Id            int64                                         `gorm:"column:id" json:"id"`
	Price         math.BigDecimal                               `gorm:"column:price" json:"price"`
	AdvertiseType AdvertiseType.AdvertiseType                   `gorm:"column:advertise_type" json:"advertiseType"`
	Member        *Member                                       `gorm:"column:member" json:"member"`
	CreateTime    xtime.Xtime                                   `gorm:"column:create_time" json:"createTime"`
	UpdateTime    xtime.Xtime                                   `gorm:"column:update_time" json:"updateTime"`
	Coin          *OtcCoin                                      `gorm:"column:coin" json:"coin"`
	Country       *Country                                      `gorm:"column:country" json:"country"`
	MinLimit      math.BigDecimal                               `gorm:"column:min_limit" json:"minLimit"`
	MaxLimit      math.BigDecimal                               `gorm:"column:max_limit" json:"maxLimit"`
	Remark        string                                        `gorm:"column:remark" json:"remark"`
	TimeLimit     int                                           `gorm:"column:time_limit" json:"timeLimit"`
	PremiseRate   math.BigDecimal                               `gorm:"column:premise_rate" json:"premiseRate"`
	Level         AdvertiseLevel.AdvertiseLevel                 `gorm:"column:level" json:"level"`
	PayMode       string                                        `gorm:"column:pay_mode" json:"payMode"`
	Version       int64                                         `gorm:"column:version" json:"version"`
	Status        AdvertiseControlStatus.AdvertiseControlStatus `gorm:"column:status" json:"status"`
	PriceType     PriceType.PriceType                           `gorm:"column:price_type" json:"priceType"`
	Number        math.BigDecimal                               `gorm:"column:number" json:"number"`
	DealAmount    math.BigDecimal                               `gorm:"column:deal_amount" json:"dealAmount"`
	RemainAmount  math.BigDecimal                               `gorm:"column:remain_amount" json:"remainAmount"`
	Auto          BooleanEnum.BooleanEnum                       `gorm:"column:auto" json:"auto"`
	MarketPrice   math.BigDecimal                               `gorm:"column:market_price" json:"marketPrice"`
	Autoword      string                                        `gorm:"column:autoword" json:"autoword"`
	LimitMoney    string                                        `gorm:"column:limit_money" json:"limitMoney"`
	Username      string                                        `gorm:"column:username" json:"username"`
	CoinUnit      string                                        `gorm:"column:coin_unit" json:"coinUnit"`
}
