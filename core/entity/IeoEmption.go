package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *IeoEmption) SetId(Id int64) (result *IeoEmption) {
	this.Id = Id
	return this
}
func (this *IeoEmption) GetId() (Id int64) {
	return this.Id
}
func (this *IeoEmption) SetIeoName(IeoName string) (result *IeoEmption) {
	this.IeoName = IeoName
	return this
}
func (this *IeoEmption) GetIeoName() (IeoName string) {
	return this.IeoName
}
func (this *IeoEmption) SetPicView(PicView string) (result *IeoEmption) {
	this.PicView = PicView
	return this
}
func (this *IeoEmption) GetPicView() (PicView string) {
	return this.PicView
}
func (this *IeoEmption) SetPic(Pic string) (result *IeoEmption) {
	this.Pic = Pic
	return this
}
func (this *IeoEmption) GetPic() (Pic string) {
	return this.Pic
}
func (this *IeoEmption) SetSaleCoin(SaleCoin string) (result *IeoEmption) {
	this.SaleCoin = SaleCoin
	return this
}
func (this *IeoEmption) GetSaleCoin() (SaleCoin string) {
	return this.SaleCoin
}
func (this *IeoEmption) SetSaleAmount(SaleAmount math.BigDecimal) (result *IeoEmption) {
	this.SaleAmount = SaleAmount
	return this
}
func (this *IeoEmption) GetSaleAmount() (SaleAmount math.BigDecimal) {
	return this.SaleAmount
}
func (this *IeoEmption) SetRaiseCoin(RaiseCoin string) (result *IeoEmption) {
	this.RaiseCoin = RaiseCoin
	return this
}
func (this *IeoEmption) GetRaiseCoin() (RaiseCoin string) {
	return this.RaiseCoin
}
func (this *IeoEmption) SetRatio(Ratio math.BigDecimal) (result *IeoEmption) {
	this.Ratio = Ratio
	return this
}
func (this *IeoEmption) GetRatio() (Ratio math.BigDecimal) {
	return this.Ratio
}
func (this *IeoEmption) SetStartTime(StartTime xtime.Xtime) (result *IeoEmption) {
	this.StartTime = StartTime
	return this
}
func (this *IeoEmption) GetStartTime() (StartTime xtime.Xtime) {
	return this.StartTime
}
func (this *IeoEmption) SetEndTime(EndTime xtime.Xtime) (result *IeoEmption) {
	this.EndTime = EndTime
	return this
}
func (this *IeoEmption) GetEndTime() (EndTime xtime.Xtime) {
	return this.EndTime
}
func (this *IeoEmption) SetFee(Fee math.BigDecimal) (result *IeoEmption) {
	this.Fee = Fee
	return this
}
func (this *IeoEmption) GetFee() (Fee math.BigDecimal) {
	return this.Fee
}
func (this *IeoEmption) SetExpectTime(ExpectTime xtime.Xtime) (result *IeoEmption) {
	this.ExpectTime = ExpectTime
	return this
}
func (this *IeoEmption) GetExpectTime() (ExpectTime xtime.Xtime) {
	return this.ExpectTime
}
func (this *IeoEmption) SetSuccessRatio(SuccessRatio math.BigDecimal) (result *IeoEmption) {
	this.SuccessRatio = SuccessRatio
	return this
}
func (this *IeoEmption) GetSuccessRatio() (SuccessRatio math.BigDecimal) {
	return this.SuccessRatio
}
func (this *IeoEmption) SetLimitAmount(LimitAmount math.BigDecimal) (result *IeoEmption) {
	this.LimitAmount = LimitAmount
	return this
}
func (this *IeoEmption) GetLimitAmount() (LimitAmount math.BigDecimal) {
	return this.LimitAmount
}
func (this *IeoEmption) SetHaveAmount(HaveAmount math.BigDecimal) (result *IeoEmption) {
	this.HaveAmount = HaveAmount
	return this
}
func (this *IeoEmption) GetHaveAmount() (HaveAmount math.BigDecimal) {
	return this.HaveAmount
}
func (this *IeoEmption) SetHaveCoin(HaveCoin string) (result *IeoEmption) {
	this.HaveCoin = HaveCoin
	return this
}
func (this *IeoEmption) GetHaveCoin() (HaveCoin string) {
	return this.HaveCoin
}
func (this *IeoEmption) SetSurplusAmount(SurplusAmount math.BigDecimal) (result *IeoEmption) {
	this.SurplusAmount = SurplusAmount
	return this
}
func (this *IeoEmption) GetSurplusAmount() (SurplusAmount math.BigDecimal) {
	return this.SurplusAmount
}
func (this *IeoEmption) SetSellMode(SellMode string) (result *IeoEmption) {
	this.SellMode = SellMode
	return this
}
func (this *IeoEmption) GetSellMode() (SellMode string) {
	return this.SellMode
}
func (this *IeoEmption) SetSellDetail(SellDetail string) (result *IeoEmption) {
	this.SellDetail = SellDetail
	return this
}
func (this *IeoEmption) GetSellDetail() (SellDetail string) {
	return this.SellDetail
}
func (this *IeoEmption) SetCreateTime(CreateTime xtime.Xtime) (result *IeoEmption) {
	this.CreateTime = CreateTime
	return this
}
func (this *IeoEmption) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *IeoEmption) SetCreateUser(CreateUser string) (result *IeoEmption) {
	this.CreateUser = CreateUser
	return this
}
func (this *IeoEmption) GetCreateUser() (CreateUser string) {
	return this.CreateUser
}
func (this *IeoEmption) SetUpdateTime(UpdateTime xtime.Xtime) (result *IeoEmption) {
	this.UpdateTime = UpdateTime
	return this
}
func (this *IeoEmption) GetUpdateTime() (UpdateTime xtime.Xtime) {
	return this.UpdateTime
}
func (this *IeoEmption) SetUpdateUser(UpdateUser string) (result *IeoEmption) {
	this.UpdateUser = UpdateUser
	return this
}
func (this *IeoEmption) GetUpdateUser() (UpdateUser string) {
	return this.UpdateUser
}
func NewIeoEmption(id int64, ieoName string, picView string, pic string, saleCoin string, saleAmount math.BigDecimal, raiseCoin string, ratio math.BigDecimal, startTime xtime.Xtime, endTime xtime.Xtime, fee math.BigDecimal, expectTime xtime.Xtime, successRatio math.BigDecimal, limitAmount math.BigDecimal, haveAmount math.BigDecimal, haveCoin string, surplusAmount math.BigDecimal, sellMode string, sellDetail string, createTime xtime.Xtime, createUser string, updateTime xtime.Xtime, updateUser string) (ret *IeoEmption) {
	ret = new(IeoEmption)
	ret.Id = id
	ret.IeoName = ieoName
	ret.PicView = picView
	ret.Pic = pic
	ret.SaleCoin = saleCoin
	ret.SaleAmount = saleAmount
	ret.RaiseCoin = raiseCoin
	ret.Ratio = ratio
	ret.StartTime = startTime
	ret.EndTime = endTime
	ret.Fee = fee
	ret.ExpectTime = expectTime
	ret.SuccessRatio = successRatio
	ret.LimitAmount = limitAmount
	ret.HaveAmount = haveAmount
	ret.HaveCoin = haveCoin
	ret.SurplusAmount = surplusAmount
	ret.SellMode = sellMode
	ret.SellDetail = sellDetail
	ret.CreateTime = createTime
	ret.CreateUser = createUser
	ret.UpdateTime = updateTime
	ret.UpdateUser = updateUser
	return
}

type IeoEmption struct {
	Id            int64           `gorm:"column:id" json:"id"`
	IeoName       string          `gorm:"column:ieo_name" json:"ieoName"`
	PicView       string          `gorm:"column:pic_view" json:"picView"`
	Pic           string          `gorm:"column:pic" json:"pic"`
	SaleCoin      string          `gorm:"column:sale_coin" json:"saleCoin"`
	SaleAmount    math.BigDecimal `gorm:"column:sale_amount" json:"saleAmount"`
	RaiseCoin     string          `gorm:"column:raise_coin" json:"raiseCoin"`
	Ratio         math.BigDecimal `gorm:"column:ratio" json:"ratio"`
	StartTime     xtime.Xtime     `gorm:"column:start_time" json:"startTime"`
	EndTime       xtime.Xtime     `gorm:"column:end_time" json:"endTime"`
	Fee           math.BigDecimal `gorm:"column:fee" json:"fee"`
	ExpectTime    xtime.Xtime     `gorm:"column:expect_time" json:"expectTime"`
	SuccessRatio  math.BigDecimal `gorm:"column:success_ratio" json:"successRatio"`
	LimitAmount   math.BigDecimal `gorm:"column:limit_amount" json:"limitAmount"`
	HaveAmount    math.BigDecimal `gorm:"column:have_amount" json:"haveAmount"`
	HaveCoin      string          `gorm:"column:have_coin" json:"haveCoin"`
	SurplusAmount math.BigDecimal `gorm:"column:surplus_amount" json:"surplusAmount"`
	SellMode      string          `gorm:"column:sell_mode" json:"sellMode"`
	SellDetail    string          `gorm:"column:sell_detail" json:"sellDetail"`
	CreateTime    xtime.Xtime     `gorm:"column:create_time" json:"createTime"`
	CreateUser    string          `gorm:"column:create_user" json:"createUser"`
	UpdateTime    xtime.Xtime     `gorm:"column:update_time" json:"updateTime"`
	UpdateUser    string          `gorm:"column:update_user" json:"updateUser"`
}
