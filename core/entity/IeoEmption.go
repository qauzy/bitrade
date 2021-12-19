package entity

import (
	"github.com/qauzy/math"
	"time"
)

func (this *IeoEmption) SetId(id int64) (result *IeoEmption) {
	this.Id = id
	return this
}
func (this *IeoEmption) GetId() (id int64) {
	return this.Id
}
func (this *IeoEmption) SetIeoName(ieoName string) (result *IeoEmption) {
	this.IeoName = ieoName
	return this
}
func (this *IeoEmption) GetIeoName() (ieoName string) {
	return this.IeoName
}
func (this *IeoEmption) SetPicView(picView string) (result *IeoEmption) {
	this.PicView = picView
	return this
}
func (this *IeoEmption) GetPicView() (picView string) {
	return this.PicView
}
func (this *IeoEmption) SetPic(pic string) (result *IeoEmption) {
	this.Pic = pic
	return this
}
func (this *IeoEmption) GetPic() (pic string) {
	return this.Pic
}
func (this *IeoEmption) SetSaleCoin(saleCoin string) (result *IeoEmption) {
	this.SaleCoin = saleCoin
	return this
}
func (this *IeoEmption) GetSaleCoin() (saleCoin string) {
	return this.SaleCoin
}
func (this *IeoEmption) SetSaleAmount(saleAmount math.BigDecimal) (result *IeoEmption) {
	this.SaleAmount = saleAmount
	return this
}
func (this *IeoEmption) GetSaleAmount() (saleAmount math.BigDecimal) {
	return this.SaleAmount
}
func (this *IeoEmption) SetRaiseCoin(raiseCoin string) (result *IeoEmption) {
	this.RaiseCoin = raiseCoin
	return this
}
func (this *IeoEmption) GetRaiseCoin() (raiseCoin string) {
	return this.RaiseCoin
}
func (this *IeoEmption) SetRatio(ratio math.BigDecimal) (result *IeoEmption) {
	this.Ratio = ratio
	return this
}
func (this *IeoEmption) GetRatio() (ratio math.BigDecimal) {
	return this.Ratio
}
func (this *IeoEmption) SetStartTime(startTime time.Time) (result *IeoEmption) {
	this.StartTime = startTime
	return this
}
func (this *IeoEmption) GetStartTime() (startTime time.Time) {
	return this.StartTime
}
func (this *IeoEmption) SetEndTime(endTime time.Time) (result *IeoEmption) {
	this.EndTime = endTime
	return this
}
func (this *IeoEmption) GetEndTime() (endTime time.Time) {
	return this.EndTime
}
func (this *IeoEmption) SetFee(fee math.BigDecimal) (result *IeoEmption) {
	this.Fee = fee
	return this
}
func (this *IeoEmption) GetFee() (fee math.BigDecimal) {
	return this.Fee
}
func (this *IeoEmption) SetExpectTime(expectTime time.Time) (result *IeoEmption) {
	this.ExpectTime = expectTime
	return this
}
func (this *IeoEmption) GetExpectTime() (expectTime time.Time) {
	return this.ExpectTime
}
func (this *IeoEmption) SetSuccessRatio(successRatio math.BigDecimal) (result *IeoEmption) {
	this.SuccessRatio = successRatio
	return this
}
func (this *IeoEmption) GetSuccessRatio() (successRatio math.BigDecimal) {
	return this.SuccessRatio
}
func (this *IeoEmption) SetLimitAmount(limitAmount math.BigDecimal) (result *IeoEmption) {
	this.LimitAmount = limitAmount
	return this
}
func (this *IeoEmption) GetLimitAmount() (limitAmount math.BigDecimal) {
	return this.LimitAmount
}
func (this *IeoEmption) SetHaveAmount(haveAmount math.BigDecimal) (result *IeoEmption) {
	this.HaveAmount = haveAmount
	return this
}
func (this *IeoEmption) GetHaveAmount() (haveAmount math.BigDecimal) {
	return this.HaveAmount
}
func (this *IeoEmption) SetHaveCoin(haveCoin string) (result *IeoEmption) {
	this.HaveCoin = haveCoin
	return this
}
func (this *IeoEmption) GetHaveCoin() (haveCoin string) {
	return this.HaveCoin
}
func (this *IeoEmption) SetSurplusAmount(surplusAmount math.BigDecimal) (result *IeoEmption) {
	this.SurplusAmount = surplusAmount
	return this
}
func (this *IeoEmption) GetSurplusAmount() (surplusAmount math.BigDecimal) {
	return this.SurplusAmount
}
func (this *IeoEmption) SetSellMode(sellMode string) (result *IeoEmption) {
	this.SellMode = sellMode
	return this
}
func (this *IeoEmption) GetSellMode() (sellMode string) {
	return this.SellMode
}
func (this *IeoEmption) SetSellDetail(sellDetail string) (result *IeoEmption) {
	this.SellDetail = sellDetail
	return this
}
func (this *IeoEmption) GetSellDetail() (sellDetail string) {
	return this.SellDetail
}
func (this *IeoEmption) SetCreateTime(createTime time.Time) (result *IeoEmption) {
	this.CreateTime = createTime
	return this
}
func (this *IeoEmption) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *IeoEmption) SetCreateUser(createUser string) (result *IeoEmption) {
	this.CreateUser = createUser
	return this
}
func (this *IeoEmption) GetCreateUser() (createUser string) {
	return this.CreateUser
}
func (this *IeoEmption) SetUpdateTime(updateTime time.Time) (result *IeoEmption) {
	this.UpdateTime = updateTime
	return this
}
func (this *IeoEmption) GetUpdateTime() (updateTime time.Time) {
	return this.UpdateTime
}
func (this *IeoEmption) SetUpdateUser(updateUser string) (result *IeoEmption) {
	this.UpdateUser = updateUser
	return this
}
func (this *IeoEmption) GetUpdateUser() (updateUser string) {
	return this.UpdateUser
}

type IeoEmption struct {
	Id            int64
	IeoName       string
	PicView       string
	Pic           string
	SaleCoin      string
	SaleAmount    math.BigDecimal
	RaiseCoin     string
	Ratio         math.BigDecimal
	StartTime     time.Time
	EndTime       time.Time
	Fee           math.BigDecimal
	ExpectTime    time.Time
	SuccessRatio  math.BigDecimal
	LimitAmount   math.BigDecimal
	HaveAmount    math.BigDecimal
	HaveCoin      string
	SurplusAmount math.BigDecimal
	SellMode      string
	SellDetail    string
	CreateTime    time.Time
	CreateUser    string
	UpdateTime    time.Time
	UpdateUser    string
}
