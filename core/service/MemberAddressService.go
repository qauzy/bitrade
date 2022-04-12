package service

func (this *MemberAddressService) AddMemberAddress(memberId int64, address string, unit string, remark string) (result *util.MessageResult, err error) {
	var coin = this.CoinDao.FindByUnit(unit)
	if coin == nil || coin.GetCanWithdraw().Equals(BooleanEnum.IS_FALSE) {
		return MessageResult.Error(600, "The currency does not support withdrawals")
	}
	var memberAddress = new(entity.MemberAddress)
	memberAddress.SetAddress(address)
	memberAddress.SetCoin(coin)
	memberAddress.SetMemberId(memberId)
	memberAddress.SetRemark(remark)
	var memberAddress1 = this.MemberAddressDao.SaveAndFlush(memberAddress)
	if memberAddress1 != nil {
		return MessageResult.Success()
	} else {
		return MessageResult.Error("failed")
	}
}
func (this *MemberAddressService) DeleteMemberAddress(memberId int64, addressId int64) (result *util.MessageResult, err error) {
	var is = this.MemberAddressDao.DeleteMemberAddress(time.Now(), addressId, memberId)
	if is > 0 {
		return MessageResult.Success()
	} else {
		return MessageResult.Error("failed")
	}
}
func (this *MemberAddressService) PageQuery(pageNo int, pageSize int, id int64, unit string) (result domain.Page[entity.MemberAddress], err error) {
	var orders = Criteria.SortStatic("id.desc")
	var pageRequest = PageRequest.Of(pageNo, pageSize, orders)
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Eq("memberId", id, false))
	specification.Add(Restrictions.Eq("status", CommonStatus.NORMAL, false))
	specification.Add(Restrictions.Eq("coin.unit", unit, false))
	return this.MemberAddressDao.FindAll(specification, pageRequest)
}
func (this *MemberAddressService) QueryAddress(userId int64, coinId string) (result arraylist.List[*hashmap.Map[string, string]], err error) {
	exception := func() (err error) {
		return Model("member_address").Field(" remark,address").Where("member_id=? and coin_id=? and status=?", userId, coinId, CommonStatus.NORMAL.Ordinal()).Select()
		return
	}()
	if exception != nil {
		e.PrintStackTrace()
	}
	return arraylist.New[*hashmap.Map[string, string]]()
}
func (this *MemberAddressService) FindByMemberIdAndAddress(userId int64, address string) (result arraylist.List[entity.MemberAddress], err error) {
	return this.MemberAddressDao.FindAllByMemberIdAndAddressAndStatus(userId, address, CommonStatus.NORMAL)
}
func (this *MemberAddressService) FindByMemberIdAndCoinAndAddress(userId int64, coin *entity.Coin, address string, status *CommonStatus.CommonStatus) (result arraylist.List[entity.MemberAddress], err error) {
	return this.MemberAddressDao.FindByMemberIdAndCoinAndAddressAndStatus(userId, coin, address, status)
}
func NewMemberAddressService(memberAddressDao *dao.MemberAddressDao, coinDao *dao.CoinDao) (ret *MemberAddressService) {
	ret = new(MemberAddressService)
	ret.MemberAddressDao = memberAddressDao
	ret.CoinDao = coinDao
	return
}

type MemberAddressService struct {
	MemberAddressDao dao.MemberAddressDao
	CoinDao          dao.CoinDao
	Base.BaseService
}
