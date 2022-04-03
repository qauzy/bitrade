package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type MemberApiKeyDao interface {
	FindMemberApiKeyByApiKey(apiKey string) (result *entity.MemberApiKey, err error)
	FindAllByMemberId(memberId int64) (result arraylist.List[entity.MemberApiKey], err error)
	FindMemberApiKeyByMemberIdAndId(memberId int64, id int64) (result *entity.MemberApiKey, err error)
	Del(id int64) (result int, err error)
	Save(m *entity.MemberApiKey) (result *entity.MemberApiKey, err error)
	FindById(id int64) (result *entity.MemberApiKey, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberApiKey], err error)
}
type memberApiKeyDao struct {
	*db.DB
}

func NewMemberApiKeyDao(db *db.DB) (dao MemberApiKeyDao) {
	dao = &memberApiKeyDao{db}
	return
}
func (this *memberApiKeyDao) FindMemberApiKeyByApiKey(apiKey string) (result *entity.MemberApiKey, err error) {
	err = this.DBRead().Where("api_key = ?", apiKey).First(&result).Error
	return
}
func (this *memberApiKeyDao) FindAllByMemberId(memberId int64) (result arraylist.List[entity.MemberApiKey], err error) {
	err = this.DBRead().Where("member_id = ?", memberId).Find(&result).Error
	return
}
func (this *memberApiKeyDao) FindMemberApiKeyByMemberIdAndId(memberId int64, id int64) (result *entity.MemberApiKey, err error) {
	err = this.DBRead().Where("member_id = ?", memberId).Where("id = ?", id).First(&result).Error
	return
}
func (this *memberApiKeyDao) Del(id int64) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("delete from MemberApiKey where id=?", id)
	err = eng.Error
	return
}
func (this *memberApiKeyDao) Save(m *entity.MemberApiKey) (result *entity.MemberApiKey, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberApiKeyDao) FindById(id int64) (result *entity.MemberApiKey, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberApiKeyDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberApiKey{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberApiKeyDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberApiKey], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
