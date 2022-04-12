package service

import (
	"bitrade/core/constant/AppealStatus"
	"bitrade/core/constant/PageModel"
	"bitrade/core/dao"
	"bitrade/core/entity"
	"bitrade/core/pagination"
	"bitrade/core/service/Base"
	"bitrade/core/vo"
	"github.com/qauzy/chocolate/lists/arraylist"
)

func (this *AppealService) FindOne(id int64) (result *entity.Appeal, err error) {
	var appeal = this.AppealDao.FindById(id).Get()
	return appeal
}
func (this *AppealService) FindOneAppealVO(id int64) (result *vo.AppealVO, err error) {
	return this.GenerateAppealVO(this.FindOne(id))
}
func (this *AppealService) Save(appeal *entity.Appeal) (result *entity.Appeal, err error) {
	return this.AppealDao.Save(appeal)
}
func (this *AppealService) JoinFind(booleanExpressionList arraylist.List[dsl.BooleanExpression], pageModel *PageModel.PageModel) (result pagination.PageResult[vo.AppealVO], err error) {
	var qAppeal = QAppeal.Appeal
	var qBean = Projections.Fields(vo.AppealVO, qAppeal.Id.As("appealId"), qAppeal.Order.MemberName.As("advertiseCreaterUserName"), qAppeal.Order.MemberRealName.As("advertiseCreaterName"), qAppeal.Order.CustomerName.As("customerUserName"), qAppeal.Order.CustomerRealName.As("customerName"), func() {
		if qAppeal.InitiatorId == qAppeal.Order.MemberId {
			return qAppeal.Order.MemberName.As("initiatorUsername")
		} else {
			return qAppeal.Order.CustomerName.As("initiatorUsername")
		}
	}(), func() {
		if qAppeal.InitiatorId == qAppeal.Order.MemberId {
			return qAppeal.Order.MemberRealName.As("initiatorName")
		} else {
			return qAppeal.Order.CustomerRealName.As("initiatorName")
		}
	}(), func() {
		if qAppeal.InitiatorId == qAppeal.Order.MemberId {
			return qAppeal.Order.CustomerName.As("associateUsername")
		} else {
			return qAppeal.Order.MemberName.As("associateUsername")
		}
	}(), func() {
		if qAppeal.InitiatorId == qAppeal.Order.MemberId {
			return qAppeal.Order.CustomerRealName.As("associateName")
		} else {
			return qAppeal.Order.MemberRealName.As("associateName")
		}
	}(), qAppeal.Order.Commission.As("fee"), qAppeal.Order.Number, qAppeal.Order.Money, qAppeal.Order.OrderSn.As("orderSn"), qAppeal.Order.CreateTime.As("transactionTime"), qAppeal.CreateTime.As("createTime"), qAppeal.DealWithTime.As("dealTime"), qAppeal.Order.PayMode.As("payMode"), qAppeal.Order.Coin.Name.As("coinName"), qAppeal.Order.Status.As("orderStatus"), qAppeal.IsSuccess.As("isSuccess"), qAppeal.Order.AdvertiseType.As("advertiseType"), qAppeal.Status, qAppeal.Remark)
	var orderSpecifiers = pageModel.GetOrderSpecifiers()
	var jpaQuery = queryFactory.Select(qBean)
	jpaQuery.From(qAppeal)
	if booleanExpressionList != nil {
		jpaQuery.Where(booleanExpressionList.ToArray(make([]BooleanExpression, 0)))
	}
	jpaQuery.OrderBy(orderSpecifiers.ToArray(make([]OrderSpecifier, 0)))
	var list = jpaQuery.Offset((pageModel.GetPageNo() - 1) * pageModel.GetPageSize()).OrderBy(orderSpecifiers.ToArray(make([]OrderSpecifier, 0))).Limit(pageModel.GetPageSize()).Fetch()
	return PageResult(list, jpaQuery.FetchCount())
}
func (this *AppealService) GenerateAppealVO(appeal *entity.Appeal) (result *vo.AppealVO, err error) {
	var initialMember = this.MemberDao.FindById(appeal.GetInitiatorId()).Get()
	var associateMember = this.MemberDao.FindById(appeal.GetAssociateId()).Get()
	var vo = new(vo.AppealVO)
	vo.SetAppealId(BigInteger.ValueOf(appeal.GetId()))
	vo.SetAssociateName(associateMember.GetRealName())
	vo.SetAssociateUsername(associateMember.GetUsername())
	vo.SetInitiatorName(initialMember.GetRealName())
	vo.SetInitiatorUsername(initialMember.GetUsername())
	var order = appeal.GetOrder()
	vo.SetCoinName(order.GetCoin().GetName())
	vo.SetFee(order.GetCommission())
	vo.SetMoney(order.GetMoney())
	vo.SetOrderSn(order.GetOrderSn())
	vo.SetNumber(order.GetNumber())
	vo.SetOrderStatus(order.GetStatus().GetOrdinal())
	vo.SetPayMode(order.GetPayMode())
	vo.SetTransactionTime(order.GetCreateTime())
	vo.SetIsSuccess(appeal.GetIsSuccess().GetOrdinal())
	vo.SetAdvertiseType(order.GetAdvertiseType().GetOrdinal())
	vo.SetAdvertiseCreaterName(order.GetMemberRealName())
	vo.SetAdvertiseCreaterUserName(order.GetMemberName())
	vo.SetCustomerUserName(order.GetCustomerName())
	vo.SetCustomerName(order.GetCustomerRealName())
	vo.SetStatus(appeal.GetStatus().GetOrdinal())
	vo.SetCreateTime(appeal.GetCreateTime())
	vo.SetDealTime(appeal.GetDealWithTime())
	vo.SetRemark(appeal.GetRemark())
	return vo
}
func (this *AppealService) FindAll(predicate Predicate, pageable *domain.Pageable) (result domain.Page[entity.Appeal], err error) {
	return this.AppealDao.FindAll(predicate, pageable)
}
func (this *AppealService) CountAuditing() (result int64, err error) {
	return this.AppealDao.CountAllByStatus(AppealStatus.NOT_PROCESSED)
}
func NewAppealService(appealDao *dao.AppealDao, memberDao *dao.MemberDao) (ret *AppealService) {
	ret = new(AppealService)
	ret.AppealDao = appealDao
	ret.MemberDao = memberDao
	return
}

type AppealService struct {
	AppealDao dao.AppealDao
	MemberDao dao.MemberDao
	Base.BaseService
}
