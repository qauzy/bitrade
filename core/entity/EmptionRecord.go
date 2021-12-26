package entity

func (this *EmptionRecord) SetId(id int64) (result *EmptionRecord) {
	this.Id = id
	return this
}
func (this *EmptionRecord) GetId() (id int64) {
	return this.Id
}
func (this *EmptionRecord) SetUserId(userId int64) (result *EmptionRecord) {
	this.UserId = userId
	return this
}
func (this *EmptionRecord) GetUserId() (userId int64) {
	return this.UserId
}
func (this *EmptionRecord) SetUserName(userName string) (result *EmptionRecord) {
	this.UserName = userName
	return this
}
func (this *EmptionRecord) GetUserName() (userName string) {
	return this.UserName
}
func (this *EmptionRecord) SetUserMobile(userMobile string) (result *EmptionRecord) {
	this.UserMobile = userMobile
	return this
}
func (this *EmptionRecord) GetUserMobile() (userMobile string) {
	return this.UserMobile
}
func (this *EmptionRecord) SetPicView(picView string) (result *EmptionRecord) {
	this.PicView = picView
	return this
}
func (this *EmptionRecord) GetPicView() (picView string) {
	return this.PicView
}
func (this *EmptionRecord) SetIeoName(ieoName string) (result *EmptionRecord) {
	this.IeoName = ieoName
	return this
}
func (this *EmptionRecord) GetIeoName() (ieoName string) {
	return this.IeoName
}
func (this *EmptionRecord) SetIeoId(ieoId int64) (result *EmptionRecord) {
	this.IeoId = ieoId
	return this
}
func (this *EmptionRecord) GetIeoId() (ieoId int64) {
	return this.IeoId
}
func (this *EmptionRecord) SetSaleCoin(saleCoin string) (result *EmptionRecord) {
	this.SaleCoin = saleCoin
	return this
}
func (this *EmptionRecord) GetSaleCoin() (saleCoin string) {
	return this.SaleCoin
}
func (this *EmptionRecord) SetSaleAmount(saleAmount math.BigDecimal) (result *EmptionRecord) {
	this.SaleAmount = saleAmount
	return this
}
func (this *EmptionRecord) GetSaleAmount() (saleAmount math.BigDecimal) {
	return this.SaleAmount
}
func (this *EmptionRecord) SetRaiseCoin(raiseCoin string) (result *EmptionRecord) {
	this.RaiseCoin = raiseCoin
	return this
}
func (this *EmptionRecord) GetRaiseCoin() (raiseCoin string) {
	return this.RaiseCoin
}
func (this *EmptionRecord) SetRatio(ratio math.BigDecimal) (result *EmptionRecord) {
	this.Ratio = ratio
	return this
}
func (this *EmptionRecord) GetRatio() (ratio math.BigDecimal) {
	return this.Ratio
}
func (this *EmptionRecord) SetStartTime(startTime time.Time) (result *EmptionRecord) {
	this.StartTime = startTime
	return this
}
func (this *EmptionRecord) GetStartTime() (startTime time.Time) {
	return this.StartTime
}
func (this *EmptionRecord) SetEndTime(endTime time.Time) (result *EmptionRecord) {
	this.EndTime = endTime
	return this
}
func (this *EmptionRecord) GetEndTime() (endTime time.Time) {
	return this.EndTime
}
func (this *EmptionRecord) SetStatus(status string) (result *EmptionRecord) {
	this.Status = status
	return this
}
func (this *EmptionRecord) GetStatus() (status string) {
	return this.Status
}
func (this *EmptionRecord) SetReceiveAmount(receiveAmount math.BigDecimal) (result *EmptionRecord) {
	this.ReceiveAmount = receiveAmount
	return this
}
func (this *EmptionRecord) GetReceiveAmount() (receiveAmount math.BigDecimal) {
	return this.ReceiveAmount
}
func (this *EmptionRecord) SetPayAmount(payAmount math.BigDecimal) (result *EmptionRecord) {
	this.PayAmount = payAmount
	return this
}
func (this *EmptionRecord) GetPayAmount() (payAmount math.BigDecimal) {
	return this.PayAmount
}
func (this *EmptionRecord) SetExpectTime(expectTime time.Time) (result *EmptionRecord) {
	this.ExpectTime = expectTime
	return this
}
func (this *EmptionRecord) GetExpectTime() (expectTime time.Time) {
	return this.ExpectTime
}
func (this *EmptionRecord) SetCreateTime(createTime time.Time) (result *EmptionRecord) {
	this.CreateTime = createTime
	return this
}
func (this *EmptionRecord) GetCreateTime() (createTime time.Time) {
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
	SaleAmount    math.BigDecimal
	RaiseCoin     string
	Ratio         math.BigDecimal
	StartTime     time.Time
	EndTime       time.Time
	Status        string
	ReceiveAmount math.BigDecimal
	PayAmount     math.BigDecimal
	ExpectTime    time.Time
	CreateTime    time.Time
}
