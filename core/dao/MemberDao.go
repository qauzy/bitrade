package dao

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CertifiedBusinessStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/math"
	"time"
)

type MemberDao interface {
	GetAllByEmailEquals(email string) (result []entity.Member, err error)
	GetAllByUsernameEquals(username string) (result []entity.Member, err error)
	GetAllByMobilePhoneEquals(phone string) (result []entity.Member, err error)
	FindByUsername(username string) (result entity.Member, err error)
	FindMemberByTokenAndTokenExpireTimeAfter(token string, date time.Time) (result entity.Member, err error)
	FindMemberByToken(token string) (result entity.Member, err error)
	FindMemberByMobilePhoneOrEmail(phone string, email string) (result entity.Member, err error)
	CountByRegistrationTimeBetween(startTime time.Time, endTime time.Time) (result int, err error)
	FindMemberByPromotionCode(code string) (result entity.Member, err error)
	FindMemberByEmail(email string) (result entity.Member, err error)
	FindMemberByMobilePhone(mobilePhone string) (result entity.Member, err error)
	FindAllByInviterId(id int64) (result []entity.Member, err error)
	FindUserNameById(id int64) (result string, err error)
	ResetSignIn() (err error)
	UpdateCertifiedBusinessStatusByIdList(idList []int64, status CertifiedBusinessStatus.CertifiedBusinessStatus) (err error)
	GetRegistrationNum(date string) (result int, err error)
	GetBussinessNum(date string) (result int, err error)
	GetApplicationNum(date string) (result int, err error)
	GetStartRegistrationDate() (result time.Time, err error)
	UpdateChannelId(memberId int64, channelId int64) (result int, err error)
	GetChannelCount(memberIds []int64) (result [][]interface{}, err error)
	UpdateLoginLock(userName string, loginLock BooleanEnum.BooleanEnum) (result int, err error)
	SaveWallet(coinId string, memberId int64, balance math.BigDecimal) (result int64, err error)
	Save(m *entity.Member) (result *entity.Member, err error)
	FindById(id int64) (result *entity.Member, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.Member, err error)
}
type memberDao struct {
	*db.DB
}

func NewMemberDao(db *db.DB) (dao MemberDao) {
	dao = &memberDao{db}
	return
}
func (this *memberDao) GetAllByEmailEquals(email string) (result []entity.Member, err error) {
	return
}
func (this *memberDao) GetAllByUsernameEquals(username string) (result []entity.Member, err error) {
	return
}
func (this *memberDao) GetAllByMobilePhoneEquals(phone string) (result []entity.Member, err error) {
	return
}
func (this *memberDao) FindByUsername(username string) (result entity.Member, err error) {
	err = this.DBRead().Where("username = ?", username).First(&result).Error
	return
}
func (this *memberDao) FindMemberByTokenAndTokenExpireTimeAfter(token string, date time.Time) (result entity.Member, err error) {
	err = this.DBRead().Where("token = ?", token).Where("token_expire_time_after = ?", date).First(&result).Error
	return
}
func (this *memberDao) FindMemberByToken(token string) (result entity.Member, err error) {
	err = this.DBRead().Where("token = ?", token).First(&result).Error
	return
}
func (this *memberDao) FindMemberByMobilePhoneOrEmail(phone string, email string) (result entity.Member, err error) {
	err = this.DBRead().Where("mobile_phone_or_email = ?", phone).First(&result).Error
	return
}
func (this *memberDao) CountByRegistrationTimeBetween(startTime time.Time, endTime time.Time) (result int, err error) {
	return
}
func (this *memberDao) FindMemberByPromotionCode(code string) (result entity.Member, err error) {
	err = this.DBRead().Where("promotion_code = ?", code).First(&result).Error
	return
}
func (this *memberDao) FindMemberByEmail(email string) (result entity.Member, err error) {
	err = this.DBRead().Where("email = ?", email).First(&result).Error
	return
}
func (this *memberDao) FindMemberByMobilePhone(mobilePhone string) (result entity.Member, err error) {
	err = this.DBRead().Where("mobile_phone = ?", mobilePhone).First(&result).Error
	return
}
func (this *memberDao) FindAllByInviterId(id int64) (result []entity.Member, err error) {
	err = this.DBRead().Where("inviter_id = ?", id).Find(&result).Error
	return
}
func (this *memberDao) FindUserNameById(id int64) (result string, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberDao) ResetSignIn() (err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update Member set signInAbility = true ")
	err = eng.Error
	return
}
func (this *memberDao) UpdateCertifiedBusinessStatusByIdList(idList []int64, status CertifiedBusinessStatus.CertifiedBusinessStatus) (err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update Member set certified_business_status = ? where id in (?) and certified_business_status=2", idList, status)
	err = eng.Error
	return
}
func (this *memberDao) GetRegistrationNum(date string) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select count(id) from member where date_format(registration_time,'%Y-%m-%d') = ?", date)
	err = eng.Error
	return
}
func (this *memberDao) GetBussinessNum(date string) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select count(id) from member where date_format(certified_business_check_time,'%Y-%m-%d') = ?", date)
	err = eng.Error
	return
}
func (this *memberDao) GetApplicationNum(date string) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select count(a.id) from member a , member_application b where a.id = b.member_id and b.audit_status = 2 and date_format(b.update_time,'%Y-%m-%d') = ?", date)
	err = eng.Error
	return
}
func (this *memberDao) GetStartRegistrationDate() (result time.Time, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select min(a.registrationTime) as date from Member a")
	err = eng.Error
	return
}
func (this *memberDao) UpdateChannelId(memberId int64, channelId int64) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update Member set channelId = ? where id = ?", memberId, channelId)
	err = eng.Error
	return
}
func (this *memberDao) GetChannelCount(memberIds []int64) (result [][]interface{}, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select m.channel_id as memberId,count(m.id) as channelCount,IFNULL((select sum(amount) from member_transaction where type=16 and member_id=m.id),0) as channelReward from member m where channel_id in (?) GROUP BY m.channel_id", memberIds)
	err = eng.Error
	return
}
func (this *memberDao) UpdateLoginLock(userName string, loginLock BooleanEnum.BooleanEnum) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update Member m set m.loginLock= ? where m.mobilePhone= ? or m.email = ?", userName, loginLock)
	err = eng.Error
	return
}
func (this *memberDao) SaveWallet(coinId string, memberId int64, balance math.BigDecimal) (result int64, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("INSERT INTO member_wallet(coin_id,member_id,balance,version,frozen_balance) VALUES (?,?,?,"
	"0,0)", coinId, memberId, balance)
	err = eng.Error
	return
}
func (this *memberDao) Save(m *entity.Member) (result *entity.Member, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberDao) FindById(id int64) (result *entity.Member, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Member{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberDao) FindAll(qp *types.QueryParam) (result []*entity.Member, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
