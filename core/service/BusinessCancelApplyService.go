
package service

import (
	"bitrade/core/constant/CertifiedBusinessStatus"
	"bitrade/core/dao"
	"bitrade/core/entity"
	"bitrade/core/service/Base"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/chocolate/maps/hashmap"
)

func (this *BusinessCancelApplyService) SetDao(dao *dao.BusinessCancelApplyDao) (err error) {
	super.SetDao(dao)
}
func (this *BusinessCancelApplyService) FindByMemberAndStaus(member *entity.Member, status *CertifiedBusinessStatus.CertifiedBusinessStatus) (result arraylist.List[entity.BusinessCancelApply], err error) {
	return dao.FindByMemberAndStatusOrderByIdDesc(member, status)
}
func (this *BusinessCancelApplyService) FindByMember(member *entity.Member) (result arraylist.List[entity.BusinessCancelApply], err error) {
	return dao.FindByMemberOrderByIdDesc(member)
}
func (this *BusinessCancelApplyService) GetBusinessOrderStatistics(memberId int64) (result *hashmap.Map[string, interface {
}], err error) {
	return this.OrderDao.GetBusinessStatistics(memberId)
}
func (this *BusinessCancelApplyService) GetBusinessAppealStatistics(memberId int64) (result *hashmap.Map[string, interface {
}], err error) {
	var oMap = new(map.HashedMap)
	var complainantNum = this.AppealDao.GetBusinessAppealInitiatorIdStatistics(memberId)
	var defendantNum = this.AppealDao.GetBusinessAppealAssociateIdStatistics(memberId)
	oMap.Put("defendantNum", defendantNum)
	oMap.Put("complainantNum", complainantNum)
	return oMap
}
func (this *BusinessCancelApplyService) GetAdvertiserNum(memberId int64) (result int64, err error) {
	var member = this.MemberDao.FindById(memberId).Get()
	return this.AdvertiseDao.GetAdvertiseNum(member)
}
func (this *BusinessCancelApplyService) CountAuditing() (result int64, err error) {
	return dao.CountAllByStatus(CertifiedBusinessStatus.CANCEL_AUTH)
}
func NewBusinessCancelApplyService(orderDao *dao.OrderDao, appealDao *dao.AppealDao, advertiseDao *dao.AdvertiseDao, memberDao *dao.MemberDao) (ret *BusinessCancelApplyService) {
	ret = new(BusinessCancelApplyService)
	ret.OrderDao = orderDao
	ret.AppealDao = appealDao
	ret.AdvertiseDao = advertiseDao
	ret.MemberDao = memberDao
	return
}

type BusinessCancelApplyService struct {
	OrderDao     dao.OrderDao
	AppealDao    dao.AppealDao
	AdvertiseDao dao.AdvertiseDao
	MemberDao    dao.MemberDao
	Base.TopBaseService[entity.BusinessCancelApply, dao.BusinessCancelApplyDao]
}

