
package entity

func (this *RewardActivitySetting) SetId(id int64) {
	this.Id = id
}
func (this *RewardActivitySetting) GetId() (id int64) {
	return this.Id
}
func (this *RewardActivitySetting) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *RewardActivitySetting) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *RewardActivitySetting) SetStatus(status BooleanEnum) {
	this.Status = status
}
func (this *RewardActivitySetting) GetStatus() (status BooleanEnum) {
	return this.Status
}
func (this *RewardActivitySetting) SetType(type ActivityRewardType) {
	this.Type = type
}
func (this *RewardActivitySetting) GetType() (type ActivityRewardType) {
	return this.Type
}
func (this *RewardActivitySetting) SetInfo(info string) {
	this.Info = info
}
func (this *RewardActivitySetting) GetInfo() (info string) {
	return this.Info
}
func (this *RewardActivitySetting) SetUpdateTime(updateTime Date) {
	this.UpdateTime = updateTime
}
func (this *RewardActivitySetting) GetUpdateTime() (updateTime Date) {
	return this.UpdateTime
}
func (this *RewardActivitySetting) SetAdmin(admin Admin) {
	this.Admin = admin
}
func (this *RewardActivitySetting) GetAdmin() (admin Admin) {
	return this.Admin
}

type RewardActivitySetting struct {
	Id         int64
	Coin       Coin
	Status     BooleanEnum
	Type       ActivityRewardType
	Info       string
	UpdateTime Date
	Admin      Admin
}

