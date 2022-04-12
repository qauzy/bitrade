package service

import (
	"bitrade/core/dao"
	"bitrade/core/entity"
	"bitrade/core/pagination"
	"bitrade/core/service/Base"
	"bitrade/core/util"
	"github.com/qauzy/chocolate/lists/arraylist"
)

func (this *InitPlateService) FindInitPlateBySymbol(symbol string) (result *entity.InitPlate, err error) {
	return this.InitPlateDao.FindInitPlateBySymbol(symbol)
}
func (this *InitPlateService) Save(initPlate *entity.InitPlate) (result *entity.InitPlate, err error) {
	return this.InitPlateDao.Save(initPlate)
}
func (this *InitPlateService) SaveAndFlush(initPlate *entity.InitPlate) (result *entity.InitPlate, err error) {
	return this.InitPlateDao.SaveAndFlush(initPlate)
}
func (this *InitPlateService) Delete(id int64) (err error) {
	this.InitPlateDao.DeleteById(id)
}
func (this *InitPlateService) FindAllByPage(specification pagination.Criteria[entity.InitPlate], pageRequest *domain.PageRequest) (result domain.Page[entity.InitPlate], err error) {
	return this.InitPlateDao.FindAll(specification, pageRequest)
}
func (this *InitPlateService) FindByInitPlateId(id int64) (result *entity.InitPlate, err error) {
	return this.InitPlateDao.FindById(id).Get()
}
func (this *InitPlateService) FindAll() (result arraylist.List[entity.InitPlate], err error) {
	return this.InitPlateDao.FindAll()
}
func (this *InitPlateService) FindAllSymbols() (result arraylist.List[string], err error) {
	var object = this.RedisUtil.Get(SysConstant.EXCHANGE_INIT_PLATE_ALL_SYMBOLS)
	if object != nil {
		return object.(arraylist.List[string])
	} else {
		var initPlates = this.InitPlateDao.FindAll()
		var symbols = arraylist.New[string]()
		for _, initPlate := range initPlates {
			symbols.Add(initPlate.GetSymbol())
		}
		this.RedisUtil.Set(SysConstant.EXCHANGE_INIT_PLATE_ALL_SYMBOLS, symbols)
		return symbols
	}
}
func NewInitPlateService(redisUtil *util.RedisUtil, initPlateDao *dao.InitPlateDao) (ret *InitPlateService) {
	ret = new(InitPlateService)
	ret.RedisUtil = redisUtil
	ret.InitPlateDao = initPlateDao
	return
}

type InitPlateService struct {
	RedisUtil    *util.RedisUtil
	InitPlateDao dao.InitPlateDao
	Base.BaseService
}
