package entity

func (this *Coin) SetName(name string) {
	this.Name = name
}
func (this *Coin) GetName() (name string) {
	return this.Name
}
func (this *Coin) SetNameCn(nameCn string) {
	this.NameCn = nameCn
}
func (this *Coin) GetNameCn() (nameCn string) {
	return this.NameCn
}
func (this *Coin) SetUnit(unit string) {
	this.Unit = unit
}
func (this *Coin) GetUnit() (unit string) {
	return this.Unit
}
func (this *Coin) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *Coin) GetStatus() (status CommonStatus) {
	return this.Status
}
func (this *Coin) SetMinTxFee(minTxFee float64) {
	this.MinTxFee = minTxFee
}
func (this *Coin) GetMinTxFee() (minTxFee float64) {
	return this.MinTxFee
}
func (this *Coin) SetCnyRate(cnyRate BigDecimal) {
	this.CnyRate = cnyRate
}
func (this *Coin) GetCnyRate() (cnyRate BigDecimal) {
	return this.CnyRate
}
func (this *Coin) SetMaxTxFee(maxTxFee float64) {
	this.MaxTxFee = maxTxFee
}
func (this *Coin) GetMaxTxFee() (maxTxFee float64) {
	return this.MaxTxFee
}
func (this *Coin) SetUsdRate(usdRate BigDecimal) {
	this.UsdRate = usdRate
}
func (this *Coin) GetUsdRate() (usdRate BigDecimal) {
	return this.UsdRate
}
func (this *Coin) SetSgdRate(sgdRate BigDecimal) {
	this.SgdRate = sgdRate
}
func (this *Coin) GetSgdRate() (sgdRate BigDecimal) {
	return this.SgdRate
}
func (this *Coin) SetEnableRpc(enableRpc BooleanEnum) {
	this.EnableRpc = enableRpc
}
func (this *Coin) GetEnableRpc() (enableRpc BooleanEnum) {
	return this.EnableRpc
}
func (this *Coin) SetSort(sort int) {
	this.Sort = sort
}
func (this *Coin) GetSort() (sort int) {
	return this.Sort
}
func (this *Coin) SetCanWithdraw(canWithdraw BooleanEnum) {
	this.CanWithdraw = canWithdraw
}
func (this *Coin) GetCanWithdraw() (canWithdraw BooleanEnum) {
	return this.CanWithdraw
}
func (this *Coin) SetCanRecharge(canRecharge BooleanEnum) {
	this.CanRecharge = canRecharge
}
func (this *Coin) GetCanRecharge() (canRecharge BooleanEnum) {
	return this.CanRecharge
}
func (this *Coin) SetCanTransfer(canTransfer BooleanEnum) {
	this.CanTransfer = canTransfer
}
func (this *Coin) GetCanTransfer() (canTransfer BooleanEnum) {
	return this.CanTransfer
}
func (this *Coin) SetCanAutoWithdraw(canAutoWithdraw BooleanEnum) {
	this.CanAutoWithdraw = canAutoWithdraw
}
func (this *Coin) GetCanAutoWithdraw() (canAutoWithdraw BooleanEnum) {
	return this.CanAutoWithdraw
}
func (this *Coin) SetWithdrawThreshold(withdrawThreshold BigDecimal) {
	this.WithdrawThreshold = withdrawThreshold
}
func (this *Coin) GetWithdrawThreshold() (withdrawThreshold BigDecimal) {
	return this.WithdrawThreshold
}
func (this *Coin) SetMinWithdrawAmount(minWithdrawAmount BigDecimal) {
	this.MinWithdrawAmount = minWithdrawAmount
}
func (this *Coin) GetMinWithdrawAmount() (minWithdrawAmount BigDecimal) {
	return this.MinWithdrawAmount
}
func (this *Coin) SetMaxWithdrawAmount(maxWithdrawAmount BigDecimal) {
	this.MaxWithdrawAmount = maxWithdrawAmount
}
func (this *Coin) GetMaxWithdrawAmount() (maxWithdrawAmount BigDecimal) {
	return this.MaxWithdrawAmount
}
func (this *Coin) SetIsPlatformCoin(isPlatformCoin BooleanEnum) {
	this.IsPlatformCoin = isPlatformCoin
}
func (this *Coin) GetIsPlatformCoin() (isPlatformCoin BooleanEnum) {
	return this.IsPlatformCoin
}
func (this *Coin) SetHasLegal(hasLegal bool) {
	this.HasLegal = hasLegal
}
func (this *Coin) GetHasLegal() (hasLegal bool) {
	return this.HasLegal
}
func (this *Coin) SetAllBalance(allBalance BigDecimal) {
	this.AllBalance = allBalance
}
func (this *Coin) GetAllBalance() (allBalance BigDecimal) {
	return this.AllBalance
}
func (this *Coin) SetColdWalletAddress(coldWalletAddress string) {
	this.ColdWalletAddress = coldWalletAddress
}
func (this *Coin) GetColdWalletAddress() (coldWalletAddress string) {
	return this.ColdWalletAddress
}
func (this *Coin) SetHotAllBalance(hotAllBalance BigDecimal) {
	this.HotAllBalance = hotAllBalance
}
func (this *Coin) GetHotAllBalance() (hotAllBalance BigDecimal) {
	return this.HotAllBalance
}
func (this *Coin) SetMinerFee(minerFee BigDecimal) {
	this.MinerFee = minerFee
}
func (this *Coin) GetMinerFee() (minerFee BigDecimal) {
	return this.MinerFee
}
func (this *Coin) SetWithdrawScale(withdrawScale int) {
	this.WithdrawScale = withdrawScale
}
func (this *Coin) GetWithdrawScale() (withdrawScale int) {
	return this.WithdrawScale
}
func (this *Coin) SetMinRechargeAmount(minRechargeAmount BigDecimal) {
	this.MinRechargeAmount = minRechargeAmount
}
func (this *Coin) GetMinRechargeAmount() (minRechargeAmount BigDecimal) {
	return this.MinRechargeAmount
}
func (this *Coin) SetMasterAddress(masterAddress string) {
	this.MasterAddress = masterAddress
}
func (this *Coin) GetMasterAddress() (masterAddress string) {
	return this.MasterAddress
}
func (this *Coin) SetMaxDailyWithdrawRate(maxDailyWithdrawRate BigDecimal) {
	this.MaxDailyWithdrawRate = maxDailyWithdrawRate
}
func (this *Coin) GetMaxDailyWithdrawRate() (maxDailyWithdrawRate BigDecimal) {
	return this.MaxDailyWithdrawRate
}

type Coin struct {
	Name                 string
	NameCn               string
	Unit                 string
	Status               CommonStatus
	MinTxFee             float64
	CnyRate              BigDecimal
	MaxTxFee             float64
	UsdRate              BigDecimal
	SgdRate              BigDecimal
	EnableRpc            BooleanEnum
	Sort                 int
	CanWithdraw          BooleanEnum
	CanRecharge          BooleanEnum
	CanTransfer          BooleanEnum
	CanAutoWithdraw      BooleanEnum
	WithdrawThreshold    BigDecimal
	MinWithdrawAmount    BigDecimal
	MaxWithdrawAmount    BigDecimal
	IsPlatformCoin       BooleanEnum
	HasLegal             bool
	AllBalance           BigDecimal
	ColdWalletAddress    string
	HotAllBalance        BigDecimal
	MinerFee             BigDecimal
	WithdrawScale        int
	MinRechargeAmount    BigDecimal
	MasterAddress        string
	MaxDailyWithdrawRate BigDecimal
}
