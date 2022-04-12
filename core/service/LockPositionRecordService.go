package service

func (this *LockPositionRecordService) FindById(id int64) (result *entity.LockPositionRecord, err error) {
	return this.LockPositionRecordDao.FindById(id).Get()
}
func (this *LockPositionRecordService) FindAll(predicate *types.Predicate, pageModel *PageModel.PageModel) (result domain.Page[entity.LockPositionRecord], err error) {
	return this.LockPositionRecordDao.FindAll(predicate, pageModel.GetPageable())
}
func (this *LockPositionRecordService) Save(lockPositionRecord *entity.LockPositionRecord) (result *entity.LockPositionRecord, err error) {
	return this.LockPositionRecordDao.Save(lockPositionRecord)
}
func (this *LockPositionRecordService) UnlockByTime(now xtime.Xtime) (err error) {
	var lockPositionRecordList = this.LockPositionRecordDao.FindByStatusAndUnlockTime(CommonStatus.NORMAL, now)
	if lockPositionRecordList != nil && lockPositionRecordList.Size() > 0 {
		var ids = arraylist.New[int64]()
		var memberTransactionList = arraylist.New[entity.MemberTransaction]()
		for _, lockPositionRecord := range lockPositionRecordList {
			this.MemberWalletDao.ThawBalance(lockPositionRecord.GetWalletId(), lockPositionRecord.GetAmount())
			var memberTransaction = new(entity.MemberTransaction)
			memberTransaction.SetMemberId(lockPositionRecord.GetMemberId())
			memberTransaction.SetAmount(lockPositionRecord.GetAmount())
			memberTransaction.SetSymbol(lockPositionRecord.GetCoin().GetUnit())
			memberTransaction.SetFee(decimal.Zero)
			memberTransaction.SetType(TransactionType.UNLOCK_POSITION)
			memberTransaction.SetAddress("")
			memberTransactionList.Add(memberTransaction)
			ids.Add(lockPositionRecord.GetId())
		}
		this.LockPositionRecordDao.UnlockByIds(ids, CommonStatus.ILLEGAL)
		this.MemberTransactionDao.SaveAll(memberTransactionList)
	}
}
func (this *LockPositionRecordService) LockPosition(memberWallet *entity.MemberWallet, amount *math.BigDecimal, member *entity.Member, reason string, unlockTime xtime.Xtime) (result *util.MessageResult, err error) {
	this.MemberWalletDao.FreezeBalance(memberWallet.GetId(), amount)
	var lockPositionRecord = new(entity.LockPositionRecord)
	lockPositionRecord.SetMemberId(memberWallet.GetMemberId())
	lockPositionRecord.SetCoin(memberWallet.GetCoin())
	lockPositionRecord.SetCreateTime(time.Now())
	lockPositionRecord.SetMemberName(member.GetUsername())
	lockPositionRecord.SetStatus(CommonStatus.NORMAL)
	lockPositionRecord.SetReason(reason)
	lockPositionRecord.SetUnlockTime(unlockTime)
	lockPositionRecord.SetAmount(amount)
	lockPositionRecord.SetWalletId(memberWallet.GetId())
	this.LockPositionRecordDao.Save(lockPositionRecord)
	var memberTransaction = new(entity.MemberTransaction)
	memberTransaction.SetMemberId(member.GetId())
	memberTransaction.SetAmount(amount.Negate())
	memberTransaction.SetSymbol(memberWallet.GetCoin().GetUnit())
	memberTransaction.SetFee(decimal.Zero)
	memberTransaction.SetType(TransactionType.LOCK_POSITION)
	memberTransaction.SetAddress("")
	this.MemberTransactionDao.Save(memberTransaction)
	return MessageResult.Success()
}
func (this *LockPositionRecordService) Unlock(lockPositionRecord *entity.LockPositionRecord) (result *util.MessageResult, err error) {
	this.LockPositionRecordDao.UnlockById(lockPositionRecord.GetId(), CommonStatus.ILLEGAL)
	this.MemberWalletDao.ThawBalance(lockPositionRecord.GetWalletId(), lockPositionRecord.GetAmount())
	var memberTransaction = new(entity.MemberTransaction)
	memberTransaction.SetMemberId(lockPositionRecord.GetMemberId())
	memberTransaction.SetAmount(lockPositionRecord.GetAmount())
	memberTransaction.SetSymbol(lockPositionRecord.GetCoin().GetUnit())
	memberTransaction.SetFee(decimal.Zero)
	memberTransaction.SetType(TransactionType.UNLOCK_POSITION)
	memberTransaction.SetAddress("")
	this.MemberTransactionDao.Save(memberTransaction)
	return MessageResult.Success()
}
func (this *LockPositionRecordService) FindByMemberIdAndCoinAndStatus(memberId int64, coin *entity.Coin, status *CommonStatus.CommonStatus) (result arraylist.List[entity.LockPositionRecord], err error) {
	return this.LockPositionRecordDao.FindByMemberIdAndCoinAndStatus(memberId, coin, status)
}
func NewLockPositionRecordService(lockPositionRecordDao *dao.LockPositionRecordDao, memberWalletDao *dao.MemberWalletDao, memberTransactionDao *dao.MemberTransactionDao) (ret *LockPositionRecordService) {
	ret = new(LockPositionRecordService)
	ret.LockPositionRecordDao = lockPositionRecordDao
	ret.MemberWalletDao = memberWalletDao
	ret.MemberTransactionDao = memberTransactionDao
	return
}

type LockPositionRecordService struct {
	LockPositionRecordDao dao.LockPositionRecordDao
	MemberWalletDao       dao.MemberWalletDao
	MemberTransactionDao  dao.MemberTransactionDao
	Base.BaseService
}
