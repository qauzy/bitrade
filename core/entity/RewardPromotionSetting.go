package entity

func (this *RewardPromotionSetting) SetId(id int64) (result *RewardPromotionSetting) {
	this.Id = id
	return this
}
func (this *RewardPromotionSetting) GetId() (id int64) {
	return this.Id
}
func (this *RewardPromotionSetting) SetCoin(coin Coin) (result *RewardPromotionSetting) {
	this.Coin = coin
	return this
}
func (this *RewardPromotionSetting) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *RewardPromotionSetting) SetStatus(status constant.BooleanEnum) (result *RewardPromotionSetting) {
	this.Status = status
	return this
}
func (this *RewardPromotionSetting) GetStatus() (status constant.BooleanEnum) {
	return this.Status
}
func (this *RewardPromotionSetting) SetType(oType constant.PromotionRewardType) (result *RewardPromotionSetting) {
	this.Type = oType
	return this
}
func (this *RewardPromotionSetting) GetType() (oType constant.PromotionRewardType) {
	return this.Type
}
func (this *RewardPromotionSetting) SetInfo(info string) (result *RewardPromotionSetting) {
	this.Info = info
	return this
}
func (this *RewardPromotionSetting) GetInfo() (info string) {
	return this.Info
}
func (this *RewardPromotionSetting) SetEffectiveTime(effectiveTime int) (result *RewardPromotionSetting) {
	this.EffectiveTime = effectiveTime
	return this
}
func (this *RewardPromotionSetting) GetEffectiveTime() (effectiveTime int) {
	return this.EffectiveTime
}
func (this *RewardPromotionSetting) SetUpdateTime(updateTime time.Time) (result *RewardPromotionSetting) {
	this.UpdateTime = updateTime
	return this
}
func (this *RewardPromotionSetting) GetUpdateTime() (updateTime time.Time) {
	return this.UpdateTime
}
func (this *RewardPromotionSetting) SetAdmin(admin Admin) (result *RewardPromotionSetting) {
	this.Admin = admin
	return this
}
func (this *RewardPromotionSetting) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *RewardPromotionSetting) SetOne(one math.BigDecimal) (result *RewardPromotionSetting) {
	this.One = one
	return this
}
func (this *RewardPromotionSetting) GetOne() (one math.BigDecimal) {
	return this.One
}
func (this *RewardPromotionSetting) SetTwo(two math.BigDecimal) (result *RewardPromotionSetting) {
	this.Two = two
	return this
}
func (this *RewardPromotionSetting) GetTwo() (two math.BigDecimal) {
	return this.Two
}

type RewardPromotionSetting struct {
	Id            int64
	Coin          Coin
	Status        constant.BooleanEnum
	Type          constant.PromotionRewardType
	Info          string
	EffectiveTime int
	UpdateTime    time.Time
	Admin         Admin
	One           math.BigDecimal
	Two           math.BigDecimal
}
