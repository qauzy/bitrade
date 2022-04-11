package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *EmptionRecord) SetId(Id int64) (result *EmptionRecord) {
	this.Id = Id
	return this
}
func (this *EmptionRecord) GetId() (Id int64) {
	return this.Id
}
func (this *EmptionRecord) SetUserId(UserId int64) (result *EmptionRecord) {
	this.UserId = UserId
	return this
}
func (this *EmptionRecord) GetUserId() (UserId int64) {
	return this.UserId
}
func (this *EmptionRecord) SetUserName(UserName string) (result *EmptionRecord) {
	this.UserName = UserName
	return this
}
func (this *EmptionRecord) GetUserName() (UserName string) {
	return this.UserName
}
func (this *EmptionRecord) SetUserMobile(UserMobile string) (result *EmptionRecord) {
	this.UserMobile = UserMobile
	return this
}
func (this *EmptionRecord) GetUserMobile() (UserMobile string) {
	return this.UserMobile
}
func (this *EmptionRecord) SetPicView(PicView string) (result *EmptionRecord) {
	this.PicView = PicView
	return this
}
func (this *EmptionRecord) GetPicView() (PicView string) {
	return this.PicView
}
func (this *EmptionRecord) SetIeoName(IeoName string) (result *EmptionRecord) {
	this.IeoName = IeoName
	return this
}
func (this *EmptionRecord) GetIeoName() (IeoName string) {
	return this.IeoName
}
func (this *EmptionRecord) SetIeoId(IeoId int64) (result *EmptionRecord) {
	this.IeoId = IeoId
	return this
}
func (this *EmptionRecord) GetIeoId() (IeoId int64) {
	return this.IeoId
}
func (this *EmptionRecord) SetSaleCoin(SaleCoin string) (result *EmptionRecord) {
	this.SaleCoin = SaleCoin
	return this
}
func (this *EmptionRecord) GetSaleCoin() (SaleCoin string) {
	return this.SaleCoin
}
func (this *EmptionRecord) SetSaleAmount(SaleAmount math.BigDecimal) (result *EmptionRecord) {
	this.SaleAmount = SaleAmount
	return this
}
func (this *EmptionRecord) GetSaleAmount() (SaleAmount math.BigDecimal) {
	return this.SaleAmount
}
func (this *EmptionRecord) SetRaiseCoin(RaiseCoin string) (result *EmptionRecord) {
	this.RaiseCoin = RaiseCoin
	return this
}
func (this *EmptionRecord) GetRaiseCoin() (RaiseCoin string) {
	return this.RaiseCoin
}
func (this *EmptionRecord) SetRatio(Ratio math.BigDecimal) (result *EmptionRecord) {
	this.Ratio = Ratio
	return this
}
func (this *EmptionRecord) GetRatio() (Ratio math.BigDecimal) {
	return this.Ratio
}
func (this *EmptionRecord) SetStartTime(StartTime xtime.Xtime) (result *EmptionRecord) {
	this.StartTime = StartTime
	return this
}
func (this *EmptionRecord) GetStartTime() (StartTime xtime.Xtime) {
	return this.StartTime
}
func (this *EmptionRecord) SetEndTime(EndTime xtime.Xtime) (result *EmptionRecord) {
	this.EndTime = EndTime
	return this
}
func (this *EmptionRecord) GetEndTime() (EndTime xtime.Xtime) {
	return this.EndTime
}
func (this *EmptionRecord) SetStatus(Status string) (result *EmptionRecord) {
	this.Status = Status
	return this
}
func (this *EmptionRecord) GetStatus() (Status string) {
	return this.Status
}
func (this *EmptionRecord) SetReceiveAmount(ReceiveAmount math.BigDecimal) (result *EmptionRecord) {
	this.ReceiveAmount = ReceiveAmount
	return this
}
func (this *EmptionRecord) GetReceiveAmount() (ReceiveAmount math.BigDecimal) {
	return this.ReceiveAmount
}
func (this *EmptionRecord) SetPayAmount(PayAmount math.BigDecimal) (result *EmptionRecord) {
	this.PayAmount = PayAmount
	return this
}
func (this *EmptionRecord) GetPayAmount() (PayAmount math.BigDecimal) {
	return this.PayAmount
}
func (this *EmptionRecord) SetExpectTime(ExpectTime xtime.Xtime) (result *EmptionRecord) {
	this.ExpectTime = ExpectTime
	return this
}
func (this *EmptionRecord) GetExpectTime() (ExpectTime xtime.Xtime) {
	return this.ExpectTime
}
func (this *EmptionRecord) SetCreateTime(CreateTime xtime.Xtime) (result *EmptionRecord) {
	this.CreateTime = CreateTime
	return this
}
func (this *EmptionRecord) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func NewEmptionRecord(id int64, userId int64, userName string, userMobile string, picView string, ieoName string, ieoId int64, saleCoin string, saleAmount math.BigDecimal, raiseCoin string, ratio math.BigDecimal, startTime xtime.Xtime, endTime xtime.Xtime, status string, receiveAmount math.BigDecimal, payAmount math.BigDecimal, expectTime xtime.Xtime, createTime xtime.Xtime) (ret *EmptionRecord) {
	ret = new(EmptionRecord)
	ret.Id = id
	ret.UserId = userId
	ret.UserName = userName
	ret.UserMobile = userMobile
	ret.PicView = picView
	ret.IeoName = ieoName
	ret.IeoId = ieoId
	ret.SaleCoin = saleCoin
	ret.SaleAmount = saleAmount
	ret.RaiseCoin = raiseCoin
	ret.Ratio = ratio
	ret.StartTime = startTime
	ret.EndTime = endTime
	ret.Status = status
	ret.ReceiveAmount = receiveAmount
	ret.PayAmount = payAmount
	ret.ExpectTime = expectTime
	ret.CreateTime = createTime
	return
}

type EmptionRecord struct {
	Id            int64           `gorm:"column:id" json:"id"`
	UserId        int64           `gorm:"column:user_id" json:"userId"`
	UserName      string          `gorm:"column:user_name" json:"userName"`
	UserMobile    string          `gorm:"column:user_mobile" json:"userMobile"`
	PicView       string          `gorm:"column:pic_view" json:"picView"`
	IeoName       string          `gorm:"column:ieo_name" json:"ieoName"`
	IeoId         int64           `gorm:"column:ieo_id" json:"ieoId"`
	SaleCoin      string          `gorm:"column:sale_coin" json:"saleCoin"`
	SaleAmount    math.BigDecimal `gorm:"column:sale_amount" json:"saleAmount"`
	RaiseCoin     string          `gorm:"column:raise_coin" json:"raiseCoin"`
	Ratio         math.BigDecimal `gorm:"column:ratio" json:"ratio"`
	StartTime     xtime.Xtime     `gorm:"column:start_time" json:"startTime"`
	EndTime       xtime.Xtime     `gorm:"column:end_time" json:"endTime"`
	Status        string          `gorm:"column:status" json:"status"`
	ReceiveAmount math.BigDecimal `gorm:"column:receive_amount" json:"receiveAmount"`
	PayAmount     math.BigDecimal `gorm:"column:pay_amount" json:"payAmount"`
	ExpectTime    xtime.Xtime     `gorm:"column:expect_time" json:"expectTime"`
	CreateTime    xtime.Xtime     `gorm:"column:create_time" json:"createTime"`
}
