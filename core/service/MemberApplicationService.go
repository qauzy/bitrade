package service

func (this *MemberApplicationService) FindAll() (result arraylist.List[entity.MemberApplication], err error) {
	return this.MemberApplicationDao.FindAll()
}
func (this *MemberApplicationService) FindAll(predicate *types.Predicate, pageable *domain.Pageable) (result domain.Page[entity.MemberApplication], err error) {
	return this.MemberApplicationDao.FindAll(predicate, pageable)
}
func (this *MemberApplicationService) FindOne(id int64) (result *entity.MemberApplication, err error) {
	return this.MemberApplicationDao.FindById(id).Get()
}
func (this *MemberApplicationService) Save(memberApplication *entity.MemberApplication) (result *entity.MemberApplication, err error) {
	return this.MemberApplicationDao.Save(memberApplication)
}
func (this *MemberApplicationService) FindLatelyReject(member *entity.Member) (result arraylist.List[entity.MemberApplication], err error) {
	return this.MemberApplicationDao.FindMemberApplicationByMemberAndAuditStatusOrderByIdDesc(member, AuditStatus.AUDIT_DEFEATED)
}
func (this *MemberApplicationService) FindSuccessRealAuthByIdCard(idCard string) (result int, err error) {
	var list = this.MemberApplicationDao.FindSuccessMemberApplicationsByIdCard(idCard, AuditStatus.AUDIT_ING, AuditStatus.AUDIT_SUCCESS)
	return list.Size()
}
func (this *MemberApplicationService) FindSuccessRecord(idCard string) (result *entity.MemberApplication, err error) {
	var list = this.MemberApplicationDao.FindSuccessMemberApplicationsByIdCard(idCard, AuditStatus.AUDIT_ING, AuditStatus.AUDIT_SUCCESS)
	return list.Get(0)
}
func (this *MemberApplicationService) Query(predicateList arraylist.List[types.Predicate], pageNo int, pageSize int) (result pagination.PageResult[entity.MemberApplication], err error) {
	var list arraylist.List[entity.MemberApplication]
	var jpaQuery = queryFactory.SelectFrom(memberApplication)
	if predicateList != nil {
		jpaQuery.Where(predicateList.ToArray(make([]Predicate, 0)))
	}
	jpaQuery.OrderBy(memberApplication.CreateTime.Desc())
	if pageNo != nil && pageSize != nil {
		list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.Fetch()
	}
	return PageResult(list, jpaQuery.FetchCount())
}
func (this *MemberApplicationService) AuditPass(application *entity.MemberApplication) (err error) {
	var kycStatus = application.GetKycStatus()
	var member = application.GetMember()
	if kycStatus == 5 {
		//????????????
		member.SetMemberLevel(MemberLevelEnum.REALNAME)
		//?????????????????????
		member.SetIdNumber(application.GetIdCard())
		//????????????????????????
		member.SetRealName(application.GetRealName())
		//???????????????????????????
		member.SetRealNameStatus(VERIFIED)
		//kyc ???????????????
		member.SetKycStatus(1)
		member.SetApplicationTime(time.Now())
		application.SetKycStatus(1)
	}
	//kyc ??????????????????????????????
	if kycStatus == 6 {
		member.SetKycStatus(4)
		application.SetKycStatus(4)
		//????????????????????? ????????????????????? ???????????????
		exception := func() (err error) {
			if member.GetInviterId() != nil {
				var inviteMember = this.MemberDao.FindById(member.GetInviterId()).Get()
				if inviteMember.GetRealNameStatus() == RealNameStatus.VERIFIED && inviteMember.GetKycStatus() == 4 {
					this.Promotion(inviteMember)
				}
			}
			return
		}()
		if exception != nil {
			log.Infof("??????????????????????????????=%v", e)
		}
	}
	this.MemberDao.Save(member)
	//????????????
	application.SetAuditStatus(AUDIT_SUCCESS)
	this.MemberApplicationDao.Save(application)
}
func (this *MemberApplicationService) Promotion(inviteMember *entity.Member) (err error) {
	var dataDictionary = this.DictionaryService.FindByBond(SysConstant.INTEGRATION_GIVING_ONE_INVITE)
	//???????????????????????????
	this.IncreaseIntegration(inviteMember, dataDictionary)
	if inviteMember.GetInviterId() != nil {
		var inviteMember2 = this.MemberDao.FindById(inviteMember.GetInviterId()).Get()
		if inviteMember2.GetRealNameStatus() == RealNameStatus.VERIFIED {
			this.PromotionLevelTwo(inviteMember2)
		}
	}
}
func (this *MemberApplicationService) PromotionLevelTwo(inviteMember2 *entity.Member) (err error) {
	var dataDictionary = this.DictionaryService.FindByBond(SysConstant.INTEGRATION_GIVING_TWO_INVITE)
	//???????????????????????????
	this.IncreaseIntegration(inviteMember2, dataDictionary)
}
func (this *MemberApplicationService) IncreaseIntegration(member *entity.Member, dataDictionary *entity.DataDictionary) (err error) {
	if dataDictionary != nil {
		if StringUtils.IsNumeric(dataDictionary.GetValue()) {
			var integrationAmount = Long.ParseLong(dataDictionary.GetValue())
			member.SetIntegration(member.GetIntegration() + integrationAmount)
			// ??????????????????????????????????????????
			member.SetGeneralizeTotal(member.GetGeneralizeTotal() + integrationAmount)
			//??????????????????????????????????????????
			var grade = this.GradeService.FindOne(member.GetMemberGradeId())
			if grade.GetId() != 5 && grade.GetId() != 6 {
				var integration = member.GetIntegration()
				if grade.GetGradeBound() < integration {
					member.SetMemberGradeId(member.GetMemberGradeId() + 1)
				}
			}
			this.MemberDao.Save(member)
			var integrationRecord = new(entity.IntegrationRecord)
			integrationRecord.SetAmount(integrationAmount)
			integrationRecord.SetMemberId(member.GetId())
			integrationRecord.SetCreateTime(time.Now())
			integrationRecord.SetType(IntegrationRecordType.PROMOTION_GIVING)
			this.IntegrationRecordService.Save(integrationRecord)
		}
	}
}
func (this *MemberApplicationService) CountAuditing() (result int64, err error) {
	return this.MemberApplicationDao.CountAllByAuditStatus(AuditStatus.AUDIT_ING)
}
func (this *MemberApplicationService) AuditNotPass(application *entity.MemberApplication) (err error) {
	var member = application.GetMember()
	var kycStatus = application.GetKycStatus()
	if kycStatus == 5 {
		application.SetKycStatus(2)
		member.SetKycStatus(2)
		//???????????????????????????
		member.SetRealNameStatus(NOT_CERTIFIED)
		//????????????
		application.SetAuditStatus(AUDIT_DEFEATED)
	}
	if kycStatus == 6 {
		application.SetKycStatus(3)
		member.SetKycStatus(3)
	}
	this.MemberDao.Save(member)
	this.MemberApplicationDao.Save(application)
}
func (this *MemberApplicationService) FindMemberApplicationByKycStatusIn(kycStatus arraylist.List[int], member *entity.Member) (result *entity.MemberApplication, err error) {
	return this.MemberApplicationDao.FindMemberApplicationByKycStatusInAndMember(kycStatus, member)
}
func NewMemberApplicationService(memberApplicationDao *dao.MemberApplicationDao, memberDao *dao.MemberDao, memberWalletService MemberWalletService, integrationRecordService IntegrationRecordService, gradeService MemberGradeService, dictionaryService DataDictionaryService) (ret *MemberApplicationService) {
	ret = new(MemberApplicationService)
	ret.MemberApplicationDao = memberApplicationDao
	ret.MemberDao = memberDao
	ret.MemberWalletService = memberWalletService
	ret.IntegrationRecordService = integrationRecordService
	ret.GradeService = gradeService
	ret.DictionaryService = dictionaryService
	return
}

type MemberApplicationService struct {
	MemberApplicationDao     dao.MemberApplicationDao
	MemberDao                dao.MemberDao
	MemberWalletService      *MemberWalletService
	IntegrationRecordService *IntegrationRecordService
	GradeService             *MemberGradeService
	DictionaryService        *DataDictionaryService
	Base.BaseService
}
