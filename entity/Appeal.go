package entity

func (this *Appeal) SetId(id int64) {
	this.Id = id
}
func (this *Appeal) GetId() (id int64) {
	return this.Id
}
func (this *Appeal) SetOrder(order Order) {
	this.Order = order
}
func (this *Appeal) GetOrder() (order Order) {
	return this.Order
}
func (this *Appeal) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *Appeal) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *Appeal) SetDealWithTime(dealWithTime Date) {
	this.DealWithTime = dealWithTime
}
func (this *Appeal) GetDealWithTime() (dealWithTime Date) {
	return this.DealWithTime
}
func (this *Appeal) SetRemark(remark string) {
	this.Remark = remark
}
func (this *Appeal) GetRemark() (remark string) {
	return this.Remark
}
func (this *Appeal) SetInitiatorId(initiatorId int64) {
	this.InitiatorId = initiatorId
}
func (this *Appeal) GetInitiatorId() (initiatorId int64) {
	return this.InitiatorId
}
func (this *Appeal) SetAssociateId(associateId int64) {
	this.AssociateId = associateId
}
func (this *Appeal) GetAssociateId() (associateId int64) {
	return this.AssociateId
}
func (this *Appeal) SetIsSuccess(isSuccess BooleanEnum) {
	this.IsSuccess = isSuccess
}
func (this *Appeal) GetIsSuccess() (isSuccess BooleanEnum) {
	return this.IsSuccess
}
func (this *Appeal) SetStatus(status AppealStatus) {
	this.Status = status
}
func (this *Appeal) GetStatus() (status AppealStatus) {
	return this.Status
}
func (this *Appeal) SetAdmin(admin Admin) {
	this.Admin = admin
}
func (this *Appeal) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *Appeal) SetImgUrls(imgUrls string) {
	this.ImgUrls = imgUrls
}
func (this *Appeal) GetImgUrls() (imgUrls string) {
	return this.ImgUrls
}

type Appeal struct {
	Id           int64
	Order        Order
	CreateTime   Date
	DealWithTime Date
	Remark       string
	InitiatorId  int64
	AssociateId  int64
	IsSuccess    BooleanEnum
	Status       AppealStatus
	Admin        Admin
	ImgUrls      string
}
