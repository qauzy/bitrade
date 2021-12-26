package entity

import (
	"bitrade/core/constant/AppealStatus"
	"bitrade/core/constant/BooleanEnum"
	"time"
)

func (this *Appeal) SetId(id int64) (result *Appeal) {
	this.Id = id
	return this
}
func (this *Appeal) GetId() (id int64) {
	return this.Id
}
func (this *Appeal) SetOrder(order Order) (result *Appeal) {
	this.Order = order
	return this
}
func (this *Appeal) GetOrder() (order Order) {
	return this.Order
}
func (this *Appeal) SetCreateTime(createTime time.Time) (result *Appeal) {
	this.CreateTime = createTime
	return this
}
func (this *Appeal) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *Appeal) SetDealWithTime(dealWithTime time.Time) (result *Appeal) {
	this.DealWithTime = dealWithTime
	return this
}
func (this *Appeal) GetDealWithTime() (dealWithTime time.Time) {
	return this.DealWithTime
}
func (this *Appeal) SetRemark(remark string) (result *Appeal) {
	this.Remark = remark
	return this
}
func (this *Appeal) GetRemark() (remark string) {
	return this.Remark
}
func (this *Appeal) SetInitiatorId(initiatorId int64) (result *Appeal) {
	this.InitiatorId = initiatorId
	return this
}
func (this *Appeal) GetInitiatorId() (initiatorId int64) {
	return this.InitiatorId
}
func (this *Appeal) SetAssociateId(associateId int64) (result *Appeal) {
	this.AssociateId = associateId
	return this
}
func (this *Appeal) GetAssociateId() (associateId int64) {
	return this.AssociateId
}
func (this *Appeal) SetIsSuccess(isSuccess BooleanEnum.BooleanEnum) (result *Appeal) {
	this.IsSuccess = isSuccess
	return this
}
func (this *Appeal) GetIsSuccess() (isSuccess BooleanEnum.BooleanEnum) {
	return this.IsSuccess
}
func (this *Appeal) SetStatus(status AppealStatus.AppealStatus) (result *Appeal) {
	this.Status = status
	return this
}
func (this *Appeal) GetStatus() (status AppealStatus.AppealStatus) {
	return this.Status
}
func (this *Appeal) SetAdmin(admin Admin) (result *Appeal) {
	this.Admin = admin
	return this
}
func (this *Appeal) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *Appeal) SetImgUrls(imgUrls string) (result *Appeal) {
	this.ImgUrls = imgUrls
	return this
}
func (this *Appeal) GetImgUrls() (imgUrls string) {
	return this.ImgUrls
}

type Appeal struct {
	Id           int64
	Order        Order
	CreateTime   time.Time
	DealWithTime time.Time
	Remark       string
	InitiatorId  int64
	AssociateId  int64
	IsSuccess    BooleanEnum.BooleanEnum
	Status       AppealStatus.AppealStatus
	Admin        Admin
	ImgUrls      string
}
