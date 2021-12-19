package dto

import "bitrade/core/entity"

func (this *MemberDTO) SetMember(member entity.Member) (result *MemberDTO) {
	this.Member = member
	return this
}
func (this *MemberDTO) GetMember() (member entity.Member) {
	return this.Member
}
func (this *MemberDTO) SetList(list []entity.MemberWallet) (result *MemberDTO) {
	this.List = list
	return this
}
func (this *MemberDTO) GetList() (list []entity.MemberWallet) {
	return this.List
}

type MemberDTO struct {
	Member entity.Member
	List   []entity.MemberWallet
}
