package Promotion

import (
	"bitrade/core/constant/PageModel"
	"bitrade/core/constant/PromotionLevel"
	"bitrade/core/constant/RewardRecordType"
	"bitrade/core/constant/TransactionType"
	"bitrade/core/entity/transform"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/maps/hashmap"
	"github.com/qauzy/chocolate/sets/hashset"
)

func (this *PromotionController) PromotionRecord(ctx *gin.Context, pageModel *PageModel.PageModel, member *transform.AuthMember) (result *util.MessageResult) {
	var predicate = QMember.Member.InviterId.Eq(member.GetId())
	var page = this.MemberService.FindAll(predicate, pageModel.GetPageable())
	var list = page.GetContent()
	var list1 = list.Stream().Map(func(x interface {
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
	var messageResult = MessageResult.Success()
	var pageResult = PageResult(list1.Stream().Sorted(func(x interface {
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
func (this *PromotionController) RewardRecord(ctx *gin.Context, pageModel *PageModel.PageModel, member *transform.AuthMember) (result *util.MessageResult) {
	var predicate = QRewardRecord.RewardRecord.Member.Id.Eq(member.GetId()).And(QRewardRecord.RewardRecord.Type.Eq(RewardRecordType.PROMOTION))
	var page = this.RewardRecordService.FindAll(predicate, pageModel)
	var list = page.GetContent()
	var result = MessageResult.Success()
	var pageResult = PageResult(list.Stream().Map(func(x interface {
	}) {
		new(PromotionRewardRecord).SetAmount(x.GetAmount()).SetCreateTime(x.GetCreateTime()).SetRemark(x.GetRemark()).SetSymbol(x.GetCoin().GetUnit())
	}).Collect(Collectors.ToList()), pageModel.GetPageNo()+1, page.GetSize(), page.GetTotalElements())
	result.SetData(pageResult)
	return result
}
func (this *PromotionController) Summary(ctx *gin.Context, member *transform.AuthMember) (result *hashmap.Map[string, interface {
}]) {
	var oMap = hashset.New[interface {
	}]()
	var predicate = QRewardRecord.RewardRecord.Member.Id.Eq(member.GetId()).And(QRewardRecord.RewardRecord.Type.Eq(RewardRecordType.PROMOTION))
	oMap.Put("count", this.RewardRecordService.FindCount(predicate))
	oMap.Put("amount", this.TransactionService.FindTransactionSum(member.GetId(), TransactionType.PROMOTION_AWARD))
	return oMap
}
func NewPromotionController(memberService *service.MemberService, rewardRecordService *service.RewardRecordService, transactionService *service.MemberTransactionService) (ret *PromotionController) {
	ret = new(PromotionController)
	ret.MemberService = memberService
	ret.RewardRecordService = rewardRecordService
	ret.TransactionService = transactionService
	return
}

type PromotionController struct {
	MemberService       *service.MemberService
	RewardRecordService *service.RewardRecordService
	TransactionService  *service.MemberTransactionService
}
