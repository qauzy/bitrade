package vo

func (this *IeoEmptionVO) SetStartTime(startTime string) (result *IeoEmptionVO) {
	this.StartTime = startTime
	return this
}
func (this *IeoEmptionVO) GetStartTime() (startTime string) {
	return this.StartTime
}
func (this *IeoEmptionVO) SetEndTime(endTime string) (result *IeoEmptionVO) {
	this.EndTime = endTime
	return this
}
func (this *IeoEmptionVO) GetEndTime() (endTime string) {
	return this.EndTime
}
func (this *IeoEmptionVO) SetStatus(status string) (result *IeoEmptionVO) {
	this.Status = status
	return this
}
func (this *IeoEmptionVO) GetStatus() (status string) {
	return this.Status
}
func (this *IeoEmptionVO) SetId(id int64) (result *IeoEmptionVO) {
	this.Id = id
	return this
}
func (this *IeoEmptionVO) GetId() (id int64) {
	return this.Id
}

type IeoEmptionVO struct {
	StartTime string
	EndTime   string
	Status    string
	Id        int64
}
