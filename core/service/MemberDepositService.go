package service

func (this *MemberDepositService) Page(predicates arraylist.List[dsl.BooleanExpression], pageModel *PageModel.PageModel) (result domain.Page[vo.MemberDepositVO], err error) {
	var query = queryFactory.Select(Projections.Fields(vo.MemberDepositVO, QMemberDeposit.MemberDeposit.Id.As("id"), QMember.Member.Username, QMember.Member.Id.As("memberId"), QMemberDeposit.MemberDeposit.Address, QMemberDeposit.MemberDeposit.Amount, QMemberDeposit.MemberDeposit.CreateTime.As("createTime"), QMemberDeposit.MemberDeposit.Unit)).From(QMember.Member, QMemberDeposit.MemberDeposit).Where(predicates.ToArray(make([]BooleanExpression, 0)))
	var orderSpecifiers = pageModel.GetOrderSpecifiers()
	query.OrderBy(orderSpecifiers.ToArray(make([]OrderSpecifier, 0)))
	var total = query.FetchCount()
	query.Offset(pageModel.GetPageSize() * (pageModel.GetPageNo() - 1)).Limit(pageModel.GetPageSize())
	var list = query.Fetch()
	return PageImpl(list, pageModel.GetPageable(), total)
}
func NewMemberDepositService(memberDepositDao *dao.MemberDepositDao) (ret *MemberDepositService) {
	ret = new(MemberDepositService)
	ret.MemberDepositDao = memberDepositDao
	return
}

type MemberDepositService struct {
	MemberDepositDao dao.MemberDepositDao
	Base.BaseService[entity.MemberDeposit]
}
