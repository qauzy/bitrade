package entity

var SerialVersionUID int64 = 1

func (this *EmptionRecord) SetId(id int64) {
	this.Id = id
}
func (this *EmptionRecord) GetId() (id int64) {
	return this.Id
}
func (this *EmptionRecord) SetUserId(userId int64) {
	this.UserId = userId
}
func (this *EmptionRecord) GetUserId() (userId int64) {
	return this.UserId
}
func (this *EmptionRecord) SetUserName(userName string) {
	this.UserName = userName
}
func (this *EmptionRecord) GetUserName() (userName string) {
	return this.UserName
}
func (this *EmptionRecord) SetUserMobile(userMobile string) {
	this.UserMobile = userMobile
}
func (this *EmptionRecord) GetUserMobile() (userMobile string) {
	return this.UserMobile
}
func (this *EmptionRecord) SetPicView(picView string) {
	this.PicView = picView
}
func (this *EmptionRecord) GetPicView() (picView string) {
	return this.PicView
}
func (this *EmptionRecord) SetIeoName(ieoName string) {
	this.IeoName = ieoName
}
func (this *EmptionRecord) GetIeoName() (ieoName string) {
	return this.IeoName
}
func (this *EmptionRecord) SetIeoId(ieoId int64) {
	this.IeoId = ieoId
}
func (this *EmptionRecord) GetIeoId() (ieoId int64) {
	return this.IeoId
}
func (this *EmptionRecord) SetSaleCoin(saleCoin string) {
	this.SaleCoin = saleCoin
}
func (this *EmptionRecord) GetSaleCoin() (saleCoin string) {
	return this.SaleCoin
}
func (this *EmptionRecord) SetSaleAmount(saleAmount BigDecimal) {
	this.SaleAmount = saleAmount
}
func (this *EmptionRecord) GetSaleAmount() (saleAmount BigDecimal) {
	return this.SaleAmount
}
func (this *EmptionRecord) SetRaiseCoin(raiseCoin string) {
	this.RaiseCoin = raiseCoin
}
func (this *EmptionRecord) GetRaiseCoin() (raiseCoin string) {
	return this.RaiseCoin
}
func (this *EmptionRecord) SetRatio(ratio BigDecimal) {
	this.Ratio = ratio
}
func (this *EmptionRecord) GetRatio() (ratio BigDecimal) {
	return this.Ratio
}
func (this *EmptionRecord) SetStartTime(startTime Date) {
	this.StartTime = startTime
}
func (this *EmptionRecord) GetStartTime() (startTime Date) {
	return this.StartTime
}
func (this *EmptionRecord) SetEndTime(endTime Date) {
	this.EndTime = endTime
}
func (this *EmptionRecord) GetEndTime() (endTime Date) {
	return this.EndTime
}
func (this *EmptionRecord) SetStatus(status string) {
	this.Status = status
}
func (this *EmptionRecord) GetStatus() (status string) {
	return this.Status
}
func (this *EmptionRecord) SetReceiveAmount(receiveAmount BigDecimal) {
	this.ReceiveAmount = receiveAmount
}
func (this *EmptionRecord) GetReceiveAmount() (receiveAmount BigDecimal) {
	return this.ReceiveAmount
}
func (this *EmptionRecord) SetPayAmount(payAmount BigDecimal) {
	this.PayAmount = payAmount
}
func (this *EmptionRecord) GetPayAmount() (payAmount BigDecimal) {
	return this.PayAmount
}
func (this *EmptionRecord) SetExpectTime(expectTime Date) {
	this.ExpectTime = expectTime
}
func (this *EmptionRecord) GetExpectTime() (expectTime Date) {
	return this.ExpectTime
}
func (this *EmptionRecord) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *EmptionRecord) GetCreateTime() (createTime Date) {
	return this.CreateTime
}

type EmptionRecord struct {
	Id            int64
	UserId        int64
	UserName      string
	UserMobile    string
	PicView       string
	IeoName       string
	IeoId         int64
	SaleCoin      string
	SaleAmount    BigDecimal
	RaiseCoin     string
	Ratio         BigDecimal
	StartTime     Date
	EndTime       Date
	Status        string
	ReceiveAmount BigDecimal
	PayAmount     BigDecimal
	ExpectTime    Date
	CreateTime    Date
}
