package entity

var SerialVersionUID int64 = 1

func (this *IeoEmption) SetId(id int64) {
	this.Id = id
}
func (this *IeoEmption) GetId() (id int64) {
	return this.Id
}
func (this *IeoEmption) SetIeoName(ieoName string) {
	this.IeoName = ieoName
}
func (this *IeoEmption) GetIeoName() (ieoName string) {
	return this.IeoName
}
func (this *IeoEmption) SetPicView(picView string) {
	this.PicView = picView
}
func (this *IeoEmption) GetPicView() (picView string) {
	return this.PicView
}
func (this *IeoEmption) SetPic(pic string) {
	this.Pic = pic
}
func (this *IeoEmption) GetPic() (pic string) {
	return this.Pic
}
func (this *IeoEmption) SetSaleCoin(saleCoin string) {
	this.SaleCoin = saleCoin
}
func (this *IeoEmption) GetSaleCoin() (saleCoin string) {
	return this.SaleCoin
}
func (this *IeoEmption) SetSaleAmount(saleAmount BigDecimal) {
	this.SaleAmount = saleAmount
}
func (this *IeoEmption) GetSaleAmount() (saleAmount BigDecimal) {
	return this.SaleAmount
}
func (this *IeoEmption) SetRaiseCoin(raiseCoin string) {
	this.RaiseCoin = raiseCoin
}
func (this *IeoEmption) GetRaiseCoin() (raiseCoin string) {
	return this.RaiseCoin
}
func (this *IeoEmption) SetRatio(ratio BigDecimal) {
	this.Ratio = ratio
}
func (this *IeoEmption) GetRatio() (ratio BigDecimal) {
	return this.Ratio
}
func (this *IeoEmption) SetStartTime(startTime Date) {
	this.StartTime = startTime
}
func (this *IeoEmption) GetStartTime() (startTime Date) {
	return this.StartTime
}
func (this *IeoEmption) SetEndTime(endTime Date) {
	this.EndTime = endTime
}
func (this *IeoEmption) GetEndTime() (endTime Date) {
	return this.EndTime
}
func (this *IeoEmption) SetFee(fee BigDecimal) {
	this.Fee = fee
}
func (this *IeoEmption) GetFee() (fee BigDecimal) {
	return this.Fee
}
func (this *IeoEmption) SetExpectTime(expectTime Date) {
	this.ExpectTime = expectTime
}
func (this *IeoEmption) GetExpectTime() (expectTime Date) {
	return this.ExpectTime
}
func (this *IeoEmption) SetSuccessRatio(successRatio BigDecimal) {
	this.SuccessRatio = successRatio
}
func (this *IeoEmption) GetSuccessRatio() (successRatio BigDecimal) {
	return this.SuccessRatio
}
func (this *IeoEmption) SetLimitAmount(limitAmount BigDecimal) {
	this.LimitAmount = limitAmount
}
func (this *IeoEmption) GetLimitAmount() (limitAmount BigDecimal) {
	return this.LimitAmount
}
func (this *IeoEmption) SetHaveAmount(haveAmount BigDecimal) {
	this.HaveAmount = haveAmount
}
func (this *IeoEmption) GetHaveAmount() (haveAmount BigDecimal) {
	return this.HaveAmount
}
func (this *IeoEmption) SetHaveCoin(haveCoin string) {
	this.HaveCoin = haveCoin
}
func (this *IeoEmption) GetHaveCoin() (haveCoin string) {
	return this.HaveCoin
}
func (this *IeoEmption) SetSurplusAmount(surplusAmount BigDecimal) {
	this.SurplusAmount = surplusAmount
}
func (this *IeoEmption) GetSurplusAmount() (surplusAmount BigDecimal) {
	return this.SurplusAmount
}
func (this *IeoEmption) SetSellMode(sellMode string) {
	this.SellMode = sellMode
}
func (this *IeoEmption) GetSellMode() (sellMode string) {
	return this.SellMode
}
func (this *IeoEmption) SetSellDetail(sellDetail string) {
	this.SellDetail = sellDetail
}
func (this *IeoEmption) GetSellDetail() (sellDetail string) {
	return this.SellDetail
}
func (this *IeoEmption) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *IeoEmption) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *IeoEmption) SetCreateUser(createUser string) {
	this.CreateUser = createUser
}
func (this *IeoEmption) GetCreateUser() (createUser string) {
	return this.CreateUser
}
func (this *IeoEmption) SetUpdateTime(updateTime Date) {
	this.UpdateTime = updateTime
}
func (this *IeoEmption) GetUpdateTime() (updateTime Date) {
	return this.UpdateTime
}
func (this *IeoEmption) SetUpdateUser(updateUser string) {
	this.UpdateUser = updateUser
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
	SaleAmount    BigDecimal
	RaiseCoin     string
	Ratio         BigDecimal
	StartTime     Date
	EndTime       Date
	Fee           BigDecimal
	ExpectTime    Date
	SuccessRatio  BigDecimal
	LimitAmount   BigDecimal
	HaveAmount    BigDecimal
	HaveCoin      string
	SurplusAmount BigDecimal
	SellMode      string
	SellDetail    string
	CreateTime    Date
	CreateUser    string
	UpdateTime    Date
	UpdateUser    string
}
