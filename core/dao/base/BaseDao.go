package base

type BaseDao interface {
	Save(m *entity.Base) (result *entity.Base, err error)
	FindById(id int64) (result *entity.Base, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.Base, err error)
}
type baseDao struct {
	*db.DB
}

func NewBaseDao(db *db.DB) (dao BaseDao) {
	dao = &baseDao{db}
	return
}
func (this *baseDao) Save(m *entity.Base) (result *entity.Base, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *baseDao) FindById(id int64) (result *entity.Base, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *baseDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Base{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *baseDao) FindAll(qp *types.QueryParam) (result []*entity.Base, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
