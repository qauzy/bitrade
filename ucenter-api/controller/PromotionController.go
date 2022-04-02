package controller

import (
	"bitrade/core/constant/PromotionLevel"
	"bitrade/core/constant/RewardRecordType"
	"bitrade/core/constant/TransactionType"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/pagination"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/qauzy/util/lists/arraylist"
	"github.com/qauzy/util/maps/hashmap"
	"github.com/qauzy/util/sets/hashset"
)

func (this *PromotionController) PromotionRecord(pageModel *PageModel.PageModel, member *transform.AuthMember) (result *util.MessageResult) {
	var predicate *types.Predicate = QMember.Member.InviterId.Eq(member.GetId())
	var page *domain.Page = this.MemberService.FindAll(predicate, pageModel.GetPageable())
	var list *arraylist.List[entity.Member] = page.GetContent()
	var list1 *arraylist.List[entity.PromotionMember] = list.Stream().Map(func(x interface {
	}) {
		new(PromotionMember).SetCreateTime(x.GetRegistrationTime()).SetLevel(PromotionLevel.ONE).SetUsername(x.GetUsername())
	}).Collect(Collectors.ToList())
	if list.Size() > 0 {
		list.Stream().ForEach(func(x interface {
		}) {
			if x.GetPromotionCode() != nil {
				list1.AddAll(this.MemberService.FindPromotionMember(x.GetId()).Stream().Map(func(y interface {
				}) {
					new(PromotionMember).SetCreateTime(y.GetRegistrationTime()).SetLevel(PromotionLevel.TWO).SetUsername(y.GetUsername())
				}).Collect(Collectors.ToList()))
			}
		})
	}
	var messageResult *util.MessageResult = MessageResult.Success()
	var pageResult *pagination.PageResult = PageResult(list1.Stream().Sorted(func(x interface {
	}, y interface {
	}) {
		if x.GetCreateTime().After(y.GetCreateTime()) {
			return -1
		} else {
			return 1
		}
	}).Collect(Collectors.ToList()), pageModel.GetPageNo()+1, page.GetSize(), page.GetTotalElements())
	messageResult.SetData(pageResult)
	return messageResult
}
func (this *PromotionController) RewardRecord(pageModel *PageModel.PageModel, member *transform.AuthMember) (result *util.MessageResult) {
	var predicate *types.Predicate = QRewardRecord.RewardRecord.Member.Id.Eq(member.GetId()).And(QRewardRecord.RewardRecord.Type.Eq(RewardRecordType.PROMOTION))
	var page *domain.Page = this.RewardRecordService.FindAll(predicate, pageModel)
	var list *arraylist.List[entity.RewardRecord] = page.GetContent()
	var result *util.MessageResult = MessageResult.Success()
	var pageResult *pagination.PageResult = PageResult(list.Stream().Map(func(x interface {
	}) {
		new(PromotionRewardRecord).SetAmount(x.GetAmount()).SetCreateTime(x.GetCreateTime()).SetRemark(x.GetRemark()).SetSymbol(x.GetCoin().GetUnit())
	}).Collect(Collectors.ToList()), pageModel.GetPageNo()+1, page.GetSize(), page.GetTotalElements())
	result.SetData(pageResult)
	return result
}
func (this *PromotionController) Summary(member *transform.AuthMember) (result *hashmap.Map[string, interface {
}]) {
	var oMap = hashset.New[interface {
	}]()
	var predicate *types.Predicate = QRewardRecord.RewardRecord.Member.Id.Eq(member.GetId()).And(QRewardRecord.RewardRecord.Type.Eq(RewardRecordType.PROMOTION))
	oMap.Put("count", this.RewardRecordService.FindCount(predicate))
	oMap.Put("amount", this.TransactionService.FindTransactionSum(member.GetId(), TransactionType.PROMOTION_AWARD))
	return oMap
}

type PromotionController struct {
	MemberService       *service.MemberService
	RewardRecordService *service.RewardRecordService
	TransactionService  *service.MemberTransactionService
}
