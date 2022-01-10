package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/math"
)

func (this *Coin) SetName(name string) (result *Coin) {
	this.Name = name
	return this
}
func (this *Coin) GetName() (name string) {
	return this.Name
}
func (this *Coin) SetNameCn(nameCn string) (result *Coin) {
	this.NameCn = nameCn
	return this
}
func (this *Coin) GetNameCn() (nameCn string) {
	return this.NameCn
}
func (this *Coin) SetUnit(unit string) (result *Coin) {
	this.Unit = unit
	return this
}
func (this *Coin) GetUnit() (unit string) {
	return this.Unit
}
func (this *Coin) SetStatus(status CommonStatus.CommonStatus) (result *Coin) {
	this.Status = status
	return this
}
func (this *Coin) GetStatus() (status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *Coin) SetMinTxFee(minTxFee float64) (result *Coin) {
	this.MinTxFee = minTxFee
	return this
}
func (this *Coin) GetMinTxFee() (minTxFee float64) {
	return this.MinTxFee
}
func (this *Coin) SetCnyRate(cnyRate math.BigDecimal) (result *Coin) {
	this.CnyRate = cnyRate
	return this
}
func (this *Coin) GetCnyRate() (cnyRate math.BigDecimal) {
	return this.CnyRate
}
func (this *Coin) SetMaxTxFee(maxTxFee float64) (result *Coin) {
	this.MaxTxFee = maxTxFee
	return this
}
func (this *Coin) GetMaxTxFee() (maxTxFee float64) {
	return this.MaxTxFee
}
func (this *Coin) SetUsdRate(usdRate math.BigDecimal) (result *Coin) {
	this.UsdRate = usdRate
	return this
}
func (this *Coin) GetUsdRate() (usdRate math.BigDecimal) {
	return this.UsdRate
}
func (this *Coin) SetSgdRate(sgdRate math.BigDecimal) (result *Coin) {
	this.SgdRate = sgdRate
	return this
}
func (this *Coin) GetSgdRate() (sgdRate math.BigDecimal) {
	return this.SgdRate
}
func (this *Coin) SetEnableRpc(enableRpc BooleanEnum.BooleanEnum) (result *Coin) {
	this.EnableRpc = enableRpc
	return this
}
func (this *Coin) GetEnableRpc() (enableRpc BooleanEnum.BooleanEnum) {
	return this.EnableRpc
}
func (this *Coin) SetSort(sort int) (result *Coin) {
	this.Sort = sort
	return this
}
func (this *Coin) GetSort() (sort int) {
	return this.Sort
}
func (this *Coin) SetCanWithdraw(canWithdraw BooleanEnum.BooleanEnum) (result *Coin) {
	this.CanWithdraw = canWithdraw
	return this
}
func (this *Coin) GetCanWithdraw() (canWithdraw BooleanEnum.BooleanEnum) {
	return this.CanWithdraw
}
func (this *Coin) SetCanRecharge(canRecharge BooleanEnum.BooleanEnum) (result *Coin) {
	this.CanRecharge = canRecharge
	return this
}
func (this *Coin) GetCanRecharge() (canRecharge BooleanEnum.BooleanEnum) {
	return this.CanRecharge
}
func (this *Coin) SetCanTransfer(canTransfer BooleanEnum.BooleanEnum) (result *Coin) {
	this.CanTransfer = canTransfer
	return this
}
func (this *Coin) GetCanTransfer() (canTransfer BooleanEnum.BooleanEnum) {
	return this.CanTransfer
}
func (this *Coin) SetCanAutoWithdraw(canAutoWithdraw BooleanEnum.BooleanEnum) (result *Coin) {
	this.CanAutoWithdraw = canAutoWithdraw
	return this
}
func (this *Coin) GetCanAutoWithdraw() (canAutoWithdraw BooleanEnum.BooleanEnum) {
	return this.CanAutoWithdraw
}
func (this *Coin) SetWithdrawThreshold(withdrawThreshold math.BigDecimal) (result *Coin) {
	this.WithdrawThreshold = withdrawThreshold
	return this
}
func (this *Coin) GetWithdrawThreshold() (withdrawThreshold math.BigDecimal) {
	return this.WithdrawThreshold
}
func (this *Coin) SetMinWithdrawAmount(minWithdrawAmount math.BigDecimal) (result *Coin) {
	this.MinWithdrawAmount = minWithdrawAmount
	return this
}
func (this *Coin) GetMinWithdrawAmount() (minWithdrawAmount math.BigDecimal) {
	return this.MinWithdrawAmount
}
func (this *Coin) SetMaxWithdrawAmount(maxWithdrawAmount math.BigDecimal) (result *Coin) {
	this.MaxWithdrawAmount = maxWithdrawAmount
	return this
}
func (this *Coin) GetMaxWithdrawAmount() (maxWithdrawAmount math.BigDecimal) {
	return this.MaxWithdrawAmount
}
func (this *Coin) SetIsPlatformCoin(isPlatformCoin BooleanEnum.BooleanEnum) (result *Coin) {
	this.IsPlatformCoin = isPlatformCoin
	return this
}
func (this *Coin) GetIsPlatformCoin() (isPlatformCoin BooleanEnum.BooleanEnum) {
	return this.IsPlatformCoin
}
func (this *Coin) SetHasLegal(hasLegal bool) (result *Coin) {
	this.HasLegal = hasLegal
	return this
}
func (this *Coin) GetHasLegal() (hasLegal bool) {
	return this.HasLegal
}
func (this *Coin) SetAllBalance(allBalance math.BigDecimal) (result *Coin) {
	this.AllBalance = allBalance
	return this
}
func (this *Coin) GetAllBalance() (allBalance math.BigDecimal) {
	return this.AllBalance
}
func (this *Coin) SetColdWalletAddress(coldWalletAddress string) (result *Coin) {
	this.ColdWalletAddress = coldWalletAddress
	return this
}
func (this *Coin) GetColdWalletAddress() (coldWalletAddress string) {
	return this.ColdWalletAddress
}
func (this *Coin) SetHotAllBalance(hotAllBalance math.BigDecimal) (result *Coin) {
	this.HotAllBalance = hotAllBalance
	return this
}
func (this *Coin) GetHotAllBalance() (hotAllBalance math.BigDecimal) {
	return this.HotAllBalance
}
func (this *Coin) SetMinerFee(minerFee math.BigDecimal) (result *Coin) {
	this.MinerFee = minerFee
	return this
}
func (this *Coin) GetMinerFee() (minerFee math.BigDecimal) {
	return this.MinerFee
}
func (this *Coin) SetWithdrawScale(withdrawScale int) (result *Coin) {
	this.WithdrawScale = withdrawScale
	return this
}
func (this *Coin) GetWithdrawScale() (withdrawScale int) {
	return this.WithdrawScale
}
func (this *Coin) SetMinRechargeAmount(minRechargeAmount math.BigDecimal) (result *Coin) {
	this.MinRechargeAmount = minRechargeAmount
	return this
}
func (this *Coin) GetMinRechargeAmount() (minRechargeAmount math.BigDecimal) {
	return this.MinRechargeAmount
}
func (this *Coin) SetMasterAddress(masterAddress string) (result *Coin) {
	this.MasterAddress = masterAddress
	return this
}
func (this *Coin) GetMasterAddress() (masterAddress string) {
	return this.MasterAddress
}
func (this *Coin) SetMaxDailyWithdrawRate(maxDailyWithdrawRate math.BigDecimal) (result *Coin) {
	this.MaxDailyWithdrawRate = maxDailyWithdrawRate
	return this
}
func (this *Coin) GetMaxDailyWithdrawRate() (maxDailyWithdrawRate math.BigDecimal) {
	return this.MaxDailyWithdrawRate
}

type Coin struct {
	Name                 string
	NameCn               string
	Unit                 string
	Status               CommonStatus.CommonStatus
	MinTxFee             float64
	CnyRate              math.BigDecimal
	MaxTxFee             float64
	UsdRate              math.BigDecimal
	SgdRate              math.BigDecimal
	EnableRpc            BooleanEnum.BooleanEnum
	Sort                 int
	CanWithdraw          BooleanEnum.BooleanEnum
	CanRecharge          BooleanEnum.BooleanEnum
	CanTransfer          BooleanEnum.BooleanEnum
	CanAutoWithdraw      BooleanEnum.BooleanEnum
	WithdrawThreshold    math.BigDecimal
	MinWithdrawAmount    math.BigDecimal
	MaxWithdrawAmount    math.BigDecimal
	IsPlatformCoin       BooleanEnum.BooleanEnum
	HasLegal             bool
	AllBalance           math.BigDecimal
	ColdWalletAddress    string
	HotAllBalance        math.BigDecimal
	MinerFee             math.BigDecimal
	WithdrawScale        int
	MinRechargeAmount    math.BigDecimal
	MasterAddress        string
	MaxDailyWithdrawRate math.BigDecimal
}
