package service

func (this *LegalWalletWithdrawService) SetDao(legalWalletWithdrawDao *dao.LegalWalletWithdrawDao) (err error) {
	super.SetDao(this.LegalWalletWithdrawDao)
}
func (this *LegalWalletWithdrawService) FindOne(id int64) (result *entity.LegalWalletWithdraw, err error) {
	return this.LegalWalletWithdrawDao.FindById(id).Get()
}
func (this *LegalWalletWithdrawService) Pass(legalWalletWithdraw *entity.LegalWalletWithdraw) (err error) {
	legalWalletWithdraw.SetStatus(WithdrawStatus.WAITING)
	legalWalletWithdraw.SetDealTime(time.Now())
	//处理时间
	this.LegalWalletWithdrawDao.Save(legalWalletWithdraw)
}
func (this *LegalWalletWithdrawService) FindDetailWeb(id int64, memberId int64) (result *entity.LegalWalletWithdraw, err error) {
	var and = QLegalWalletWithdraw.LegalWalletWithdraw.Id.Eq(id).And(QLegalWalletWithdraw.LegalWalletWithdraw.Member.Id.Eq(memberId))
	return this.LegalWalletWithdrawDao.FindOne(and).Get()
}
func (this *LegalWalletWithdrawService) Withdraw(wallet *entity.MemberWallet, legalWalletWithdraw *entity.LegalWalletWithdraw) (err error) {
	wallet.SetBalance(BigDecimalUtils.Sub(wallet.GetBalance(), legalWalletWithdraw.GetAmount()))
	wallet.SetFrozenBalance(BigDecimalUtils.Add(wallet.GetFrozenBalance(), legalWalletWithdraw.GetAmount()))
	this.LegalWalletWithdrawDao.Save(legalWalletWithdraw)
}
func (this *LegalWalletWithdrawService) NoPass(wallet *entity.MemberWallet, withdraw *entity.LegalWalletWithdraw) (err error) {
	wallet.SetFrozenBalance(BigDecimalUtils.Sub(wallet.GetFrozenBalance(), this.Withdraw.GetAmount()))
	//冻结金额减少
	wallet.SetBalance(BigDecimalUtils.Add(wallet.GetBalance(), this.Withdraw.GetAmount()))
	//本金增加
	this.Withdraw.SetStatus(WithdrawStatus.FAIL)
	//标记失败
	this.Withdraw.SetDealTime(time.Now())
	//处理时间
}
func (this *LegalWalletWithdrawService) Remit(paymentInstrument string, withdraw *entity.LegalWalletWithdraw, wallet *entity.MemberWallet) (err error) {
	this.Withdraw.SetPaymentInstrument(paymentInstrument)
	//支付凭证
	this.Withdraw.SetStatus(WithdrawStatus.SUCCESS)
	//标记成功
	this.Withdraw.SetRemitTime(time.Now())
	//打款时间
	wallet.SetFrozenBalance(BigDecimalUtils.Sub(wallet.GetFrozenBalance(), this.Withdraw.GetAmount()))
	//钱包冻结金额减少
}
func NewLegalWalletWithdrawService(legalWalletWithdrawDao *dao.LegalWalletWithdrawDao) (ret *LegalWalletWithdrawService) {
	ret = new(LegalWalletWithdrawService)
	ret.LegalWalletWithdrawDao = legalWalletWithdrawDao
	return
}

type LegalWalletWithdrawService struct {
	LegalWalletWithdrawDao dao.LegalWalletWithdrawDao
	Base.TopBaseService[entity.LegalWalletWithdraw, dao.LegalWalletWithdrawDao]
}
