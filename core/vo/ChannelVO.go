package vo

import "github.com/qauzy/math"

func (this *ChannelVO) SetMemberId(memberId int64) (result *ChannelVO) {
	this.MemberId = memberId
	return this
}
func (this *ChannelVO) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *ChannelVO) SetChannelReward(channelReward math.BigDecimal) (result *ChannelVO) {
	this.ChannelReward = channelReward
	return this
}
func (this *ChannelVO) GetChannelReward() (channelReward math.BigDecimal) {
	return this.ChannelReward
}
func (this *ChannelVO) SetChannelCount(channelCount int64) (result *ChannelVO) {
	this.ChannelCount = channelCount
	return this
}
func (this *ChannelVO) GetChannelCount() (channelCount int64) {
	return this.ChannelCount
}
func NewChannelVO(memberId int64, channelCount int64, channelReward math.BigDecimal) (this *ChannelVO) {
	this = new(ChannelVO)
	this.MemberId = memberId
	this.ChannelCount = channelCount
	this.ChannelReward = channelReward
	return
}

//func NewChannelVO() (this *ChannelVO) {
//	this = new(ChannelVO)
//	super()
//	return
//}

type ChannelVO struct {
	MemberId      int64
	ChannelReward math.BigDecimal
	ChannelCount  int64
}
