package service

func (this *InviteManagementService) LookAll(inviteManagementVO *vo.InviteManagementVO) (result Page[entity.Member], err error) {
	var releaseBalance = new(pagination.Criteria)
	var sort = releaseBalance.Sort("registrationTime.desc")
	// PageNum 当前页 PageSize 每页多少条
	var pageRequest = PageRequest.Of(inviteManagementVO.GetPageNo()-1, inviteManagementVO.GetPageSize(), sort)
	return this.Dao.FindAll(releaseBalance, pageRequest)
}
func (this *InviteManagementService) QueryCondition(imVO *vo.InviteManagementVO) (result Page[entity.Member], err error) {
	var criteria = new(pagination.Criteria)
	var sort = criteria.Sort("registrationTime.desc")
	if StringUtils.IsNotEmpty(imVO.GetRealName()) {
		criteria.Add(Restrictions.Eq("realName", imVO.GetRealName(), false))
	}
	if StringUtils.IsNotEmpty(imVO.GetMobilePhone()) {
		criteria.Add(Restrictions.Eq("mobilePhone", imVO.GetMobilePhone(), false))
	}
	if StringUtils.IsNotEmpty(imVO.GetEmail()) {
		criteria.Add(Restrictions.Eq("email", imVO.GetEmail(), false))
	}
	var pageRequest = PageRequest.Of(imVO.GetPageNo()-1, imVO.GetPageSize(), sort)
	return this.Dao.FindAll(criteria, pageRequest)
}
func (this *InviteManagementService) QueryId(inviteManagementVO *vo.InviteManagementVO) (result Page[entity.Member], err error) {
	var member = this.Dao.FindById(inviteManagementVO.GetId()).Get()
	// 获取当前用户的id  根据id去查询邀请者ID 为这个id 的所有的数据
	var memberList = arraylist.New[entity.Member]()
	var firstMemberList = this.Dao.FindAllByInviterId(member.GetId())
	for i := 0; i < firstMemberList.Size(); i += 1 {
		memberList.Add(firstMemberList.Get(i))
	}
	// 存储一级数据.
	var set = hashset.New[interface {
	}]()
	for i := 0; i < firstMemberList.Size(); i += 1 {
		set.Add(firstMemberList.Get(i).GetId())
	}
	for _, setId := range set {
		var list = this.Dao.FindAllByInviterId(setId.(int64))
		for i := 0; i < list.Size(); i += 1 {
			memberList.Add(list.Get(i))
		}
	}
	// 对查询出来的list进行ID倒序排序
	ListSort(memberList)
	var pageable = PageRequest.Of(inviteManagementVO.GetPageNumber()-1, inviteManagementVO.GetPageSize())
	var pageData = this.GetPageData(pageable, memberList)
	return pageData
}
func (this *InviteManagementService) GetPageData(pageable Pageable, maps arraylist.List[entity.Member]) (result PageImpl[entity.Member], err error) {
	var page PageImpl
	exception := func() (err error) {
		// 每页起始
		var startValue = pageable.GetPageNumber() * pageable.GetPageSize()
		// 每页终止
		var endValue = (pageable.GetPageNumber() + 1) * pageable.GetPageSize()
		if endValue > maps.Size() {
			endValue = maps.Size()
		}
		var dataList = Lists.NewArrayList()
		for i := startValue; i < endValue; i += 1 {
			dataList.Add(maps.Get(i))
		}
		var count = 0
		if dataList != nil && dataList.Size() > 0 {
			count = Long.ValueOf(maps.Size())
		}
		page = PageImpl(dataList, pageable, count)
		return
	}()
	if exception != nil {
		log.Infof("分页异常了 getPageData=%v", e)
	}
	return page
}
func ListSort(list arraylist.List[entity.Member]) (err error) {
	Collections.Sort(list, func(o1 interface {
	}, o2 interface {
	}) {
		var format = SimpleDateFormat("yyyy-MM-dd HH:mm:ss")
		exception := func() (err error) {
			var dt1 = format.Parse(String.ValueOf(o1.GetRegistrationTime()))
			var dt2 = format.Parse(String.ValueOf(o2.GetRegistrationTime()))
			if dt1.GetTime() < dt2.GetTime() {
				return 1
			} else if dt1.GetTime() > dt2.GetTime() {
				return -1
			} else {
				return 0
			}
			return
		}()
		if exception != nil {
			log.Infof("倒序时间错误 ListSort=%v", e)
		}
		return 0
	})
}
func NewInviteManagementService(dao *dao.MemberDao) (ret *InviteManagementService) {
	ret = new(InviteManagementService)
	ret.Dao = dao
	return
}

type InviteManagementService struct {
	Dao dao.MemberDao
	Base.BaseService
}
