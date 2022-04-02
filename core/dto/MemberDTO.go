package dto

import (
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

func (this *MemberDTO) SetMember(Member *entity.Member) (result *MemberDTO) {
	this.Member = Member
	return this
}
func (this *MemberDTO) GetMember() (Member *entity.Member) {
	return this.Member
}
func (this *MemberDTO) SetList(List *arraylist.List[entity.MemberWallet]) (result *MemberDTO) {
	this.List = List
	return this
}
func (this *MemberDTO) GetList() (List *arraylist.List[entity.MemberWallet]) {
	return this.List
}

type MemberDTO struct {
	Member *entity.Member                       `gorm:"column:member" json:"member"`
	List   *arraylist.List[entity.MemberWallet] `gorm:"column:list" json:"list"`
}
