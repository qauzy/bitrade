
package entity

func (this *RewardPromotionSetting) SetId(id int64) {
	this.Id = id
}
func (this *RewardPromotionSetting) GetId() (id int64) {
	return this.Id
}
func (this *RewardPromotionSetting) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *RewardPromotionSetting) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *RewardPromotionSetting) SetStatus(status BooleanEnum) {
	this.Status = status
}
func (this *RewardPromotionSetting) GetStatus() (status BooleanEnum) {
	return this.Status
}
func (this *RewardPromotionSetting) SetType(type PromotionRewardType) {
	this.Type = type
}
func (this *RewardPromotionSetting) GetType() (type PromotionRewardType) {
	return this.Type
}
func (this *RewardPromotionSetting) SetInfo(info string) {
	this.Info = info
}
func (this *RewardPromotionSetting) GetInfo() (info string) {
	return this.Info
}
func (this *RewardPromotionSetting) SetEffectiveTime(effectiveTime int) {
	this.EffectiveTime = effectiveTime
}
func (this *RewardPromotionSetting) GetEffectiveTime() (effectiveTime int) {
	return this.EffectiveTime
}
func (this *RewardPromotionSetting) SetUpdateTime(updateTime Date) {
	this.UpdateTime = updateTime
}
func (this *RewardPromotionSetting) GetUpdateTime() (updateTime Date) {
	return this.UpdateTime
}
func (this *RewardPromotionSetting) SetAdmin(admin Admin) {
	this.Admin = admin
}
func (this *RewardPromotionSetting) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *RewardPromotionSetting) SetOne(one BigDecimal) {
	this.One = one
}
func (this *RewardPromotionSetting) GetOne() (one BigDecimal) {
	return this.One
}
func (this *RewardPromotionSetting) SetTwo(two BigDecimal) {
	this.Two = two
}
func (this *RewardPromotionSetting) GetTwo() (two BigDecimal) {
	return this.Two
}

type RewardPromotionSetting struct {
	Id            int64
	Coin          Coin
	Status        BooleanEnum
	Type          PromotionRewardType
	Info          string
	EffectiveTime int
	UpdateTime    Date
	Admin         Admin
	One           BigDecimal
	Two           BigDecimal
}

