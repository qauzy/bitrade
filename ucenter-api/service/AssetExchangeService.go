package service

import (
	"bitrade/core/constant/TransactionType"
	"bitrade/core/service/Base"
	"bitrade/core/util"
	"bitrade/ucenter-api/dao"
	"bitrade/ucenter-api/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/math"
)

func (this *AssetExchangeService) FindAllByFromCoin(toUnit string) (result arraylist.List[entity.AssetExchangeCoin], err error) {
	return this.AssetExchangeDao.FindAllByToUnit(toUnit)
}
func (this *AssetExchangeService) FindOne(fromUnit string, toUnit string) (result *entity.AssetExchangeCoin, err error) {
	return this.AssetExchangeDao.FindByFromUnitAndToUnit(fromUnit, toUnit)
}
func (this *AssetExchangeService) Exchange(memberId int64, coin *entity.AssetExchangeCoin, amount math.BigDecimal) (result *util.MessageResult, err error) {
	var fromWallet, err = this.MemberWalletService.FindByCoinUnitAndMemberId(coin.GetFromUnit(), memberId)

	var toWallet, err = this.MemberWalletService.FindByCoinUnitAndMemberId(coin.GetToUnit(), memberId)
	if fromWallet == nil || toWallet == nil {
		return util.NewMessageResultV2(500, "钱包不存在"), nil
	}
	var toAmount = amount.Mul(coin.GetExchangeRate())
	if this.MemberWalletService.DeductBalance(fromWallet, amount) > 0 {
		//增加余额
		this.MemberWalletService.IncreaseBalance(toWallet.GetId(), toAmount)
		//增加入资金记录
		var transaction = new(entity.MemberTransaction)
		transaction.SetAmount(toAmount)
		transaction.SetSymbol(coin.GetToUnit())
		transaction.SetAddress("")
		transaction.SetMemberId(toWallet.GetMemberId())
		transaction.SetType(TransactionType.ASSET_EXCHANGE)
		transaction.SetFee(math.Zero)
		this.TransactionService.Save(transaction)
		//增加出资金记录
		var transaction2 *entity.MemberTransaction = new(entity.MemberTransaction)
		transaction2.SetAmount(amount.Neg())
		transaction2.SetSymbol(coin.GetFromUnit())
		transaction2.SetAddress("")
		transaction2.SetMemberId(fromWallet.GetMemberId())
		transaction2.SetType(TransactionType.ASSET_EXCHANGE)
		transaction2.SetFee(math.Zero)
		this.TransactionService.Save(transaction2)
		return util.NewMessageResultV2(0, "success"), nil
	} else {
		return util.NewMessageResultV2(500, "余额不足"), nil
	}
}

type AssetExchangeService struct {
	AssetExchangeDao    dao.AssetExchangeDao
	MemberWalletService *MemberWalletService
	TransactionService  *MemberTransactionService
	Base.BaseService
}
