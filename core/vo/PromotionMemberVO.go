package vo

func (this *PromotionMemberVO) SetId(id int64) (result *PromotionMemberVO) {
	this.Id = id
	return this
}
func (this *PromotionMemberVO) GetId() (id int64) {
	return this.Id
}
func (this *PromotionMemberVO) SetUsername(username string) (result *PromotionMemberVO) {
	this.Username = username
	return this
}
func (this *PromotionMemberVO) GetUsername() (username string) {
	return this.Username
}
func (this *PromotionMemberVO) SetRealName(realName string) (result *PromotionMemberVO) {
	this.RealName = realName
	return this
}
func (this *PromotionMemberVO) GetRealName() (realName string) {
	return this.RealName
}
func (this *PromotionMemberVO) SetEmail(email string) (result *PromotionMemberVO) {
	this.Email = email
	return this
}
func (this *PromotionMemberVO) GetEmail() (email string) {
	return this.Email
}
func (this *PromotionMemberVO) SetMobilePhone(mobilePhone string) (result *PromotionMemberVO) {
	this.MobilePhone = mobilePhone
	return this
}
func (this *PromotionMemberVO) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *PromotionMemberVO) SetPromotionCode(promotionCode string) (result *PromotionMemberVO) {
	this.PromotionCode = promotionCode
	return this
}
func (this *PromotionMemberVO) GetPromotionCode() (promotionCode string) {
	return this.PromotionCode
}
func (this *PromotionMemberVO) SetPromotionNum(promotionNum int) (result *PromotionMemberVO) {
	this.PromotionNum = promotionNum
	return this
}
func (this *PromotionMemberVO) GetPromotionNum() (promotionNum int) {
	return this.PromotionNum
}
func (this *PromotionMemberVO) SetReward(reward map[string]string) (result *PromotionMemberVO) {
	this.Reward = reward
	return this
}
func (this *PromotionMemberVO) GetReward() (reward map[string]string) {
	return this.Reward
}

type PromotionMemberVO struct {
	Id            int64
	Username      string
	RealName      string
	Email         string
	MobilePhone   string
	PromotionCode string
	PromotionNum  int
	Reward        map[string]string
}
