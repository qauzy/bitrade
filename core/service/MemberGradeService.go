package service

func (this *MemberGradeService) FindAll() (result arraylist.List[entity.MemberGrade], err error) {
	return this.Dao.FindAll()
}
func (this *MemberGradeService) FindOne(id int64) (result *entity.MemberGrade, err error) {
	//        Object object = redisUtil.get(SysConstant.CUSTOMER_INTEGRATION_GRADE+id);
	//        MemberGrade memberGrade ;
	//        if(object==null){
	//            memberGrade = dao.findById(id).get();
	//            redisUtil.set(SysConstant.CUSTOMER_INTEGRATION_GRADE+id,memberGrade);
	//        }else {
	//            memberGrade = (MemberGrade)object;
	//        }
	return this.Dao.FindById(id).Get()
}
func (this *MemberGradeService) Save(memberGrade *entity.MemberGrade) (result *entity.MemberGrade, err error) {
	if memberGrade.GetId() != nil {
		this.RedisUtil.Delete(SysConstant.CUSTOMER_INTEGRATION_GRADE + memberGrade.GetId())
	}
	return this.Dao.Save(memberGrade)
}
func NewMemberGradeService(dao *dao.MemberGradeDao, redisUtil *util.RedisUtil) (ret *MemberGradeService) {
	ret = new(MemberGradeService)
	ret.Dao = dao
	ret.RedisUtil = redisUtil
	return
}

type MemberGradeService struct {
	Dao       dao.MemberGradeDao
	RedisUtil *util.RedisUtil
	Base.BaseService[entity.MemberGrade]
}
