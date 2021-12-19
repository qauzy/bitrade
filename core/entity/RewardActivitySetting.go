package entity

func (this *RewardActivitySetting) SetId(id int64) (result *RewardActivitySetting) {
	this.Id = id
	return this
}
func (this *RewardActivitySetting) GetId() (id int64) {
	return this.Id
}
func (this *RewardActivitySetting) SetCoin(coin Coin) (result *RewardActivitySetting) {
	this.Coin = coin
	return this
}
func (this *RewardActivitySetting) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *RewardActivitySetting) SetStatus(status constant.BooleanEnum) (result *RewardActivitySetting) {
	this.Status = status
	return this
}
func (this *RewardActivitySetting) GetStatus() (status constant.BooleanEnum) {
	return this.Status
}
func (this *RewardActivitySetting) SetType(oType constant.ActivityRewardType) (result *RewardActivitySetting) {
	this.Type = oType
	return this
}
func (this *RewardActivitySetting) GetType() (oType constant.ActivityRewardType) {
	return this.Type
}
func (this *RewardActivitySetting) SetInfo(info string) (result *RewardActivitySetting) {
	this.Info = info
	return this
}
func (this *RewardActivitySetting) GetInfo() (info string) {
	return this.Info
}
func (this *RewardActivitySetting) SetUpdateTime(updateTime time.Time) (result *RewardActivitySetting) {
	this.UpdateTime = updateTime
	return this
}
func (this *RewardActivitySetting) GetUpdateTime() (updateTime time.Time) {
	return this.UpdateTime
}
func (this *RewardActivitySetting) SetAdmin(admin Admin) (result *RewardActivitySetting) {
	this.Admin = admin
	return this
}
func (this *RewardActivitySetting) GetAdmin() (admin Admin) {
	return this.Admin
}

type RewardActivitySetting struct {
	Id         int64
	Coin       Coin
	Status     constant.BooleanEnum
	Type       constant.ActivityRewardType
	Info       string
	UpdateTime time.Time
	Admin      Admin
}
