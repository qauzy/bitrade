package vo

func (this *GiftConfigVO) SetStartTime(startTime string) (result *GiftConfigVO) {
	this.StartTime = startTime
	return this
}
func (this *GiftConfigVO) GetStartTime() (startTime string) {
	return this.StartTime
}
func (this *GiftConfigVO) SetEndTime(endTime string) (result *GiftConfigVO) {
	this.EndTime = endTime
	return this
}
func (this *GiftConfigVO) GetEndTime() (endTime string) {
	return this.EndTime
}
func (this *GiftConfigVO) SetGiftName(giftName string) (result *GiftConfigVO) {
	this.GiftName = giftName
	return this
}
func (this *GiftConfigVO) GetGiftName() (giftName string) {
	return this.GiftName
}
func (this *GiftConfigVO) SetId(id int64) (result *GiftConfigVO) {
	this.Id = id
	return this
}
func (this *GiftConfigVO) GetId() (id int64) {
	return this.Id
}

type GiftConfigVO struct {
	StartTime string
	EndTime   string
	GiftName  string
	Id        int64
}
