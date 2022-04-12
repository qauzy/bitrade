package service

func (this *TransferAddressService) SetDao(dao *dao.TransferAddressDao) (err error) {
	super.SetDao(dao)
}
func (this *TransferAddressService) FindByUnit(unit string) (result arraylist.List[entity.TransferAddress], err error) {
	var coin = this.CoinDao.FindByUnit(unit)
	return dao.FindAllByStatusAndCoin(CommonStatus.NORMAL, coin)
}
func (this *TransferAddressService) FindByCoin(coin *entity.Coin) (result arraylist.List[entity.TransferAddress], err error) {
	return dao.FindAllByStatusAndCoin(CommonStatus.NORMAL, coin)
}
func (this *TransferAddressService) FindOnlyOne(coin *entity.Coin, address string) (result *entity.TransferAddress, err error) {
	return dao.FindByAddressAndCoin(address, coin)
}
func NewTransferAddressService(coinDao *dao.CoinDao) (ret *TransferAddressService) {
	ret = new(TransferAddressService)
	ret.CoinDao = coinDao
	return
}

type TransferAddressService struct {
	CoinDao dao.CoinDao
	Base.TopBaseService[entity.TransferAddress, dao.TransferAddressDao]
}
