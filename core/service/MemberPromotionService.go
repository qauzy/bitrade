package service

func (this *MemberPromotionService) Save(memberPromotion *entity.MemberPromotion) (result *entity.MemberPromotion, err error) {
	return this.MemberPromotionDao.Save(memberPromotion)
}
func (this *MemberPromotionService) GetPromotionDetails(memberId int64, pageModel *PageModel.PageModel) (result domain.Page[vo.RegisterPromotionVO], err error) {
	var headSql = StringBuilder("select a.id id ,a.username presentee,a.email presenteeEmail, ").Append("a.mobile_phone presenteePhone,a.real_name presenteeRealName,a.registration_time promotionTime, ").Append("b.level promotionLevel ")
	var endSql = StringBuilder(" from member a join member_promotion b on a.id = b.invitees_id and b.inviter_id = " + memberId)
	var countHead = StringBuilder("select count(*) ")
	var page = createNativePageQuery(countHead.Append(endSql), headSql.Append(endSql), pageModel, Transformers.AliasToBean(vo.RegisterPromotionVO))
	/*  RewardPromotionSetting setting = rewardPromotionSettingService.findByType(PromotionRewardType.REGISTER) ;

	    BigDecimal one = JSONObject.parseObject(setting.getInfo()).getBigDecimal("one");
	    BigDecimal two = JSONObject.parseObject(setting.getInfo()).getBigDecimal("two");

	    for(RegisterPromotionVO vo:page.getContent()){
	        vo.setUnit(setting.getCoin().getUnit());
	        if(vo.getPromotionLevel()== PromotionLevel.ONE.getOrdinal()){
	            vo.setReward(one);
	        }else if(vo.getPromotionLevel()== PromotionLevel.TWO.getOrdinal()){
	            vo.setReward(two);
	        }else{
	            vo.setReward(BigDecimal.ZERO);
	        }
	    }*/
	return page
}
func NewMemberPromotionService(entityManager *persistence.EntityManager, memberPromotionDao *dao.MemberPromotionDao, rewardPromotionSettingService RewardPromotionSettingService) (ret *MemberPromotionService) {
	ret = new(MemberPromotionService)
	ret.EntityManager = entityManager
	ret.MemberPromotionDao = memberPromotionDao
	ret.RewardPromotionSettingService = rewardPromotionSettingService
	return
}

type MemberPromotionService struct {
	EntityManager                 *persistence.EntityManager
	MemberPromotionDao            dao.MemberPromotionDao
	RewardPromotionSettingService *RewardPromotionSettingService
	Base.BaseService
}
