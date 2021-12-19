package vo

import (
	"github.com/qauzy/math"
	"time"
)

func (this *RegisterPromotionVO) SetId(id math.BigInteger) (result *RegisterPromotionVO) {
	this.Id = id
	return this
}
func (this *RegisterPromotionVO) GetId() (id math.BigInteger) {
	return this.Id
}
func (this *RegisterPromotionVO) SetPresentee(presentee string) (result *RegisterPromotionVO) {
	this.Presentee = presentee
	return this
}
func (this *RegisterPromotionVO) GetPresentee() (presentee string) {
	return this.Presentee
}
func (this *RegisterPromotionVO) SetPresenteeRealName(presenteeRealName string) (result *RegisterPromotionVO) {
	this.PresenteeRealName = presenteeRealName
	return this
}
func (this *RegisterPromotionVO) GetPresenteeRealName() (presenteeRealName string) {
	return this.PresenteeRealName
}
func (this *RegisterPromotionVO) SetPresenteeEmail(presenteeEmail string) (result *RegisterPromotionVO) {
	this.PresenteeEmail = presenteeEmail
	return this
}
func (this *RegisterPromotionVO) GetPresenteeEmail() (presenteeEmail string) {
	return this.PresenteeEmail
}
func (this *RegisterPromotionVO) SetPresenteePhone(presenteePhone string) (result *RegisterPromotionVO) {
	this.PresenteePhone = presenteePhone
	return this
}
func (this *RegisterPromotionVO) GetPresenteePhone() (presenteePhone string) {
	return this.PresenteePhone
}
func (this *RegisterPromotionVO) SetPromotionTime(promotionTime time.Time) (result *RegisterPromotionVO) {
	this.PromotionTime = promotionTime
	return this
}
func (this *RegisterPromotionVO) GetPromotionTime() (promotionTime time.Time) {
	return this.PromotionTime
}
func (this *RegisterPromotionVO) SetPromotionLevel(promotionLevel int) (result *RegisterPromotionVO) {
	this.PromotionLevel = promotionLevel
	return this
}
func (this *RegisterPromotionVO) GetPromotionLevel() (promotionLevel int) {
	return this.PromotionLevel
}

type RegisterPromotionVO struct {
	Id                math.BigInteger
	Presentee         string
	PresenteeRealName string
	PresenteeEmail    string
	PresenteePhone    string
	PromotionTime     time.Time
	PromotionLevel    int
}
