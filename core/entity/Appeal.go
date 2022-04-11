package entity

import (
	"bitrade/core/constant/AppealStatus"
	"bitrade/core/constant/BooleanEnum"
	"github.com/qauzy/chocolate/xtime"
)

func (this *Appeal) SetId(Id int64) (result *Appeal) {
	this.Id = Id
	return this
}
func (this *Appeal) GetId() (Id int64) {
	return this.Id
}
func (this *Appeal) SetOrder(Order *Order) (result *Appeal) {
	this.Order = Order
	return this
}
func (this *Appeal) GetOrder() (Order *Order) {
	return this.Order
}
func (this *Appeal) SetCreateTime(CreateTime xtime.Xtime) (result *Appeal) {
	this.CreateTime = CreateTime
	return this
}
func (this *Appeal) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *Appeal) SetDealWithTime(DealWithTime xtime.Xtime) (result *Appeal) {
	this.DealWithTime = DealWithTime
	return this
}
func (this *Appeal) GetDealWithTime() (DealWithTime xtime.Xtime) {
	return this.DealWithTime
}
func (this *Appeal) SetRemark(Remark string) (result *Appeal) {
	this.Remark = Remark
	return this
}
func (this *Appeal) GetRemark() (Remark string) {
	return this.Remark
}
func (this *Appeal) SetInitiatorId(InitiatorId int64) (result *Appeal) {
	this.InitiatorId = InitiatorId
	return this
}
func (this *Appeal) GetInitiatorId() (InitiatorId int64) {
	return this.InitiatorId
}
func (this *Appeal) SetAssociateId(AssociateId int64) (result *Appeal) {
	this.AssociateId = AssociateId
	return this
}
func (this *Appeal) GetAssociateId() (AssociateId int64) {
	return this.AssociateId
}
func (this *Appeal) SetIsSuccess(IsSuccess BooleanEnum.BooleanEnum) (result *Appeal) {
	this.IsSuccess = IsSuccess
	return this
}
func (this *Appeal) GetIsSuccess() (IsSuccess BooleanEnum.BooleanEnum) {
	return this.IsSuccess
}
func (this *Appeal) SetStatus(Status AppealStatus.AppealStatus) (result *Appeal) {
	this.Status = Status
	return this
}
func (this *Appeal) GetStatus() (Status AppealStatus.AppealStatus) {
	return this.Status
}
func (this *Appeal) SetAdmin(Admin *Admin) (result *Appeal) {
	this.Admin = Admin
	return this
}
func (this *Appeal) GetAdmin() (Admin *Admin) {
	return this.Admin
}
func (this *Appeal) SetImgUrls(ImgUrls string) (result *Appeal) {
	this.ImgUrls = ImgUrls
	return this
}
func (this *Appeal) GetImgUrls() (ImgUrls string) {
	return this.ImgUrls
}
func NewAppeal(id int64, order *Order, createTime xtime.Xtime, dealWithTime xtime.Xtime, remark string, initiatorId int64, associateId int64, isSuccess BooleanEnum.BooleanEnum, status AppealStatus.AppealStatus, admin *Admin, imgUrls string) (ret *Appeal) {
	ret = new(Appeal)
	ret.Id = id
	ret.Order = order
	ret.CreateTime = createTime
	ret.DealWithTime = dealWithTime
	ret.Remark = remark
	ret.InitiatorId = initiatorId
	ret.AssociateId = associateId
	ret.IsSuccess = isSuccess
	ret.Status = status
	ret.Admin = admin
	ret.ImgUrls = imgUrls
	return
}

type Appeal struct {
	Id           int64                     `gorm:"column:id" json:"id"`
	Order        *Order                    `gorm:"column:order" json:"order"`
	CreateTime   xtime.Xtime               `gorm:"column:create_time" json:"createTime"`
	DealWithTime xtime.Xtime               `gorm:"column:deal_with_time" json:"dealWithTime"`
	Remark       string                    `gorm:"column:remark" json:"remark"`
	InitiatorId  int64                     `gorm:"column:initiator_id" json:"initiatorId"`
	AssociateId  int64                     `gorm:"column:associate_id" json:"associateId"`
	IsSuccess    BooleanEnum.BooleanEnum   `gorm:"column:is_success" json:"isSuccess"`
	Status       AppealStatus.AppealStatus `gorm:"column:status" json:"status"`
	Admin        *Admin                    `gorm:"column:admin" json:"admin"`
	ImgUrls      string                    `gorm:"column:img_urls" json:"imgUrls"`
}
