package service

func (this *BusinessAuthApplyService) Page(predicate *types.Predicate, pageModel *PageModel.PageModel) (result domain.Page[entity.BusinessAuthApply], err error) {
	return this.BusinessAuthApplyDao.FindAll(predicate, pageModel.GetPageable())
}
func (this *BusinessAuthApplyService) FindByMember(member *entity.Member) (result arraylist.List[entity.BusinessAuthApply], err error) {
	return this.BusinessAuthApplyDao.FindByMemberOrderByIdDesc(member)
}
func (this *BusinessAuthApplyService) FindOne(id int64) (result *entity.BusinessAuthApply, err error) {
	return this.BusinessAuthApplyDao.FindById(id).Get()
}
func (this *BusinessAuthApplyService) Create(businessAuthApply *entity.BusinessAuthApply) (err error) {
	this.BusinessAuthApplyDao.Save(businessAuthApply)
}
func (this *BusinessAuthApplyService) Update(businessAuthApply *entity.BusinessAuthApply) (err error) {
	this.BusinessAuthApplyDao.Save(businessAuthApply)
}
func (this *BusinessAuthApplyService) FindByMemberAndCertifiedBusinessStatus(member *entity.Member, certifiedBusinessStatus *CertifiedBusinessStatus.CertifiedBusinessStatus) (result arraylist.List[entity.BusinessAuthApply], err error) {
	return this.BusinessAuthApplyDao.FindByMemberAndCertifiedBusinessStatusOrderByIdDesc(member, certifiedBusinessStatus)
}
func (this *BusinessAuthApplyService) Save(businessAuthApply *entity.BusinessAuthApply) (result *entity.BusinessAuthApply, err error) {
	return this.BusinessAuthApplyDao.Save(businessAuthApply)
}
func (this *BusinessAuthApplyService) Page(predicate *types.Predicate, pageable *domain.Pageable) (result domain.Page[entity.BusinessAuthApply], err error) {
	return this.BusinessAuthApplyDao.FindAll(predicate, pageable)
}
func (this *BusinessAuthApplyService) Detail(id int64) (result *util.MessageResult, err error) {
	var qBusinessAuthApply = QBusinessAuthApply.BusinessAuthApply
	var query = queryFactory.Select(Projections.Fields(entity.BusinessAuthApplyDetailVO, qBusinessAuthApply.Id.As("id"), qBusinessAuthApply.CertifiedBusinessStatus.As("status"), qBusinessAuthApply.Amount.As("amount"), qBusinessAuthApply.AuthInfo.As("authInfo"), qBusinessAuthApply.Member.RealName.As("realName"), qBusinessAuthApply.Detail.As("detail"), qBusinessAuthApply.AuditingTime.As("checkTime"))).From(qBusinessAuthApply)
	query.Where(qBusinessAuthApply.Id.Eq(id))
	var vo = query.FetchOne()
	var result *util.MessageResult
	var jsonStr = vo.GetAuthInfo()
	log.Infof("认证信息 jsonStr = %v", jsonStr)
	if StringUtils.IsEmpty(jsonStr) {
		result = MessageResult.Error("认证相关信息不存在")
		result.SetData(vo)
		return result
	}
	exception := func() (err error) {
		var json = JSONObject.ParseObject(jsonStr)
		vo.SetInfo(json)
		result = MessageResult.Success("认证详情")
		result.SetData(vo)
		return result
		return
	}()
	if exception != nil {
		log.Infof("认证信息格式异常:%v", e)
		result = MessageResult.Error("认证信息格式异常")
		return result
	}
}
func (this *BusinessAuthApplyService) CountAuditing() (result int64, err error) {
	return this.BusinessAuthApplyDao.CountAllByCertifiedBusinessStatus(CertifiedBusinessStatus.AUDITING)
}
func NewBusinessAuthApplyService(businessAuthApplyDao *dao.BusinessAuthApplyDao) (ret *BusinessAuthApplyService) {
	ret = new(BusinessAuthApplyService)
	ret.BusinessAuthApplyDao = businessAuthApplyDao
	return
}

type BusinessAuthApplyService struct {
	BusinessAuthApplyDao dao.BusinessAuthApplyDao
	Base.BaseService
}
