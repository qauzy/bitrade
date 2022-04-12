package service

func (this *MemberService) QueryWhereOrPage(booleanExpressionList arraylist.List[dsl.BooleanExpression], pageNo int, pageSize int) (result pagination.PageResult[entity.Member], err error) {
	var list arraylist.List[entity.Member]
	var jpaQuery = queryFactory.SelectFrom(QMember.Member).Where(booleanExpressionList.ToArray(make([]BooleanExpression, 0)))
	jpaQuery.OrderBy(QMember.Member.Id.Desc())
	if pageNo != nil && pageSize != nil {
		list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.Fetch()
	}
	return PageResult(list, jpaQuery.FetchCount())
}
func (this *MemberService) Save(member *entity.Member) (result *entity.Member, err error) {
	return this.MemberDao.SaveAndFlush(member)
}
func (this *MemberService) FindAll(predicate *types.Predicate) (result arraylist.List[entity.Member], err error) {
	//return Collections.;
	var iterable = this.MemberDao.FindAll(predicate)
	var list = IteratorUtils.ToList(iterable.Iterator())
	return list
}
func (this *MemberService) SaveAndFlush(member *entity.Member) (result *entity.Member, err error) {
	return this.MemberDao.SaveAndFlush(member)
}
func (this *MemberService) LoginWithToken(token string, ip string, device string) (result *entity.Member, err error) {
	if StringUtils.IsBlank(token) {
		return nil
	}
	//Member mr = memberDao.findMemberByTokenAndTokenExpireTimeAfter(token,new Date());
	var mr = this.MemberDao.FindMemberByToken(token)
	return mr
}
func (this *MemberService) Login(username string, password string) (result *entity.Member, err error) {
	var member = this.MemberDao.FindMemberByMobilePhoneOrEmail(username, username)
	if member == nil {
		return AuthenticationException("账号或密码错误")
	} else if member.GetLoginLock() != nil && member.GetLoginLock().Equals(BooleanEnum.IS_TRUE) {
		return AuthenticationException("多次输入错误的密码，账号已锁定，请联系客服解锁或找回密码解锁")
	} else if !Md5.Md5Digest(password + member.GetSalt()).ToLowerCase().Equals(member.GetPassword()) {
		log.Info("账号或密码错误")
		return nil
	} else if member.GetStatus().Equals(CommonStatus.ILLEGAL) {
		return AuthenticationException("该帐号已经被禁用，请联系客服")
	}
	return member
}
func (this *MemberService) FindOne(id int64) (result *entity.Member, err error) {
	return this.MemberDao.FindById(id).Get()
}
func (this *MemberService) FindAll() (result arraylist.List[entity.Member], err error) {
	return this.MemberDao.FindAll()
}
func (this *MemberService) Count() (result int64, err error) {
	return this.MemberDao.Count()
}
func (this *MemberService) FindAll(predicate *types.Predicate, pageModel *PageModel.PageModel) (result domain.Page[entity.Member], err error) {
	return this.MemberDao.FindAll(predicate, pageModel.GetPageable())
}
func (this *MemberService) FindPromotionMember(id int64) (result arraylist.List[entity.Member], err error) {
	return this.MemberDao.FindAllByInviterId(id)
}
func (this *MemberService) Page(pageNo int, pageSize int, status *CommonStatus.CommonStatus) (result domain.Page[entity.Member], err error) {
	//排序方式 (需要倒序 这样    Criteria.sort("id","createTime.desc") ) //参数实体类为字段名
	var orders = Criteria.SortStatic("id")
	//分页参数
	var pageRequest = PageRequest.Of(pageNo, pageSize, orders)
	//查询条件
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Eq("status", status, false))
	return this.MemberDao.FindAll(specification, pageRequest)
}
func (this *MemberService) Page(pageNo int, pageSize int) (result domain.Page[entity.Member], err error) {
	//排序方式 (需要倒序 这样    Criteria.sort("id","createTime.desc") ) //参数实体类为字段名
	var orders = Criteria.SortStatic("id")
	//分页参数
	var pageRequest = PageRequest.Of(pageNo, pageSize, orders)
	return this.MemberDao.FindAll(pageRequest)
}
func (this *MemberService) EmailIsExist(email string) (result bool, err error) {
	var list = this.MemberDao.GetAllByEmailEquals(email)
	if list.Size() > 0 {
		return true
	} else {
		return false
	}
}
func (this *MemberService) UsernameIsExist(username string) (result bool, err error) {
	if this.MemberDao.GetAllByUsernameEquals(username).Size() > 0 {
		return true
	} else {
		return false
	}
}
func (this *MemberService) PhoneIsExist(phone string) (result bool, err error) {
	if this.MemberDao.GetAllByMobilePhoneEquals(phone).Size() > 0 {
		return true
	} else {
		return false
	}
}
func (this *MemberService) FindByUsername(username string) (result *entity.Member, err error) {
	return this.MemberDao.FindByUsername(username)
}
func (this *MemberService) FindByEmail(email string) (result *entity.Member, err error) {
	return this.MemberDao.FindMemberByEmail(email)
}
func (this *MemberService) FindByPhone(phone string) (result *entity.Member, err error) {
	return this.MemberDao.FindMemberByMobilePhone(phone)
}
func (this *MemberService) FindByPhoneOrEmail(str string) (result *entity.Member, err error) {
	return this.MemberDao.FindMemberByMobilePhoneOrEmail(str, str)
}
func (this *MemberService) FindAll(predicate *types.Predicate, pageable *domain.Pageable) (result domain.Page[entity.Member], err error) {
	return this.MemberDao.FindAll(predicate, pageable)
}
func (this *MemberService) FindUserNameById(id int64) (result string, err error) {
	return this.MemberDao.FindUserNameById(id)
}
func (this *MemberService) SignInIncident(member *entity.Member, memberWallet *entity.MemberWallet, sign sign) (err error) {
	member.SetSignInAbility(false)
	//失去签到能力
	memberWallet.SetBalance(BigDecimalUtils.Add(memberWallet.GetBalance(), sign.GetAmount()))
	//签到收益
	// 签到记录
	this.SignRecordDao.Save(cn.Ztuo.Bitrade.Entity.MemberSignRecord(member, sign))
	//账单明细
	var memberTransaction = cn.Ztuo.Bitrade.Entity.MemberTransaction()
	memberTransaction.SetMemberId(member.GetId())
	memberTransaction.SetAmount(sign.GetAmount())
	memberTransaction.SetType(TransactionType.ACTIVITY_AWARD)
	memberTransaction.SetSymbol(sign.GetCoin().GetUnit())
	this.TransactionDao.Save(memberTransaction)
}
func (this *MemberService) ResetSignIn() (err error) {
	this.MemberDao.ResetSignIn()
}
func (this *MemberService) UpdateCertifiedBusinessStatusByIdList(idList arraylist.List[int64]) (err error) {
	this.MemberDao.UpdateCertifiedBusinessStatusByIdList(idList, CertifiedBusinessStatus.DEPOSIT_LESS)
}
func (this *MemberService) GetChannelCount(memberIds arraylist.List[int64]) (result arraylist.List[vo.ChannelVO], err error) {
	var list = this.MemberDao.GetChannelCount(memberIds)
	var channelVOList = arraylist.New[vo.ChannelVO]()
	if list != nil && list.Size() > 0 {
		for _, objs := range list {
			var memberId = objs[0].(int)
			var channelCount = objs[1].(int)
			var channelReward = objs[2].(int)
			var channelVO = ChannelVO(memberId.LongValue(), channelCount.IntValue(), BigDecimal(channelReward.DoubleValue()))
			channelVOList.Add(channelVO)
		}
	}
	return channelVOList
}
func (this *MemberService) Lock(username string) (err error) {
	this.MemberDao.UpdateLoginLock(username, BooleanEnum.IS_TRUE)
}
func (this *MemberService) UnLock(username string) (err error) {
	this.MemberDao.UpdateLoginLock(username, BooleanEnum.IS_FALSE)
}
func (this *MemberService) SaveWallet(coinId string, memberId int64, balance *math.BigDecimal) (result int, err error) {
	return this.MemberDao.SaveWallet(coinId, memberId, balance)
}
func NewMemberService(memberDao *dao.MemberDao, signRecordDao *dao.MemberSignRecordDao, transactionDao *dao.MemberTransactionDao) (ret *MemberService) {
	ret = new(MemberService)
	ret.MemberDao = memberDao
	ret.SignRecordDao = signRecordDao
	ret.TransactionDao = transactionDao
	return
}

type MemberService struct {
	MemberDao      dao.MemberDao
	SignRecordDao  dao.MemberSignRecordDao
	TransactionDao dao.MemberTransactionDao
	Base.BaseService
}
