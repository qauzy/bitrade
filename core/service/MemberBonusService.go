package service

func (this *MemberBonusService) GetBonusByMemberId(memberId int64) (result arraylist.List[dto.MemberBonusDTO], err error) {
	return this.MemberBonusDao.GetBonusByMemberId(memberId)
}
func (this *MemberBonusService) GetBonusAmountByMemberId(memberId int64) (result *math.BigDecimal, err error) {
	return this.MemberBonusDao.GetBonusAmountByMemberId(memberId)
}
func (this *MemberBonusService) Save(memberBonusDTO *dto.MemberBonusDTO) (result *dto.MemberBonusDTO, err error) {
	return this.MemberBonusDao.Save(memberBonusDTO)
}
func (this *MemberBonusService) GetBonusByMemberIdPage(memberId int64, pageNo int, pageSize int) (result domain.Page[dto.MemberBonusDTO], err error) {
	var orders = Criteria.SortStatic("id.desc")
	var pageRequest = PageRequest.Of(pageNo, pageSize, orders)
	var criteria = new(pagination.Criteria)
	criteria.Add(Restrictions.Eq("memberId", memberId, false))
	return this.MemberBonusDao.FindAll(criteria, pageRequest)
}
func (this *MemberBonusService) GetMemberBounsPage(pageNo int, pageSize int) (result domain.Page[dto.MemberBonusDTO], err error) {
	var orders = Criteria.SortStatic("id.desc")
	var pageRequest = PageRequest.Of(pageNo, pageSize, orders)
	return this.MemberBonusDao.FindAll(pageRequest)
}
func NewMemberBonusService(memberBonusDao *dao.MemberBonusDao) (ret *MemberBonusService) {
	ret = new(MemberBonusService)
	ret.MemberBonusDao = memberBonusDao
	return
}

type MemberBonusService struct {
	MemberBonusDao dao.MemberBonusDao
	Base.BaseService[dto.MemberBonusDTO]
}
