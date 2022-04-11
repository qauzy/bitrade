package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/math"
)

func (this *Coin) SetName(Name string) (result *Coin) {
	this.Name = Name
	return this
}
func (this *Coin) GetName() (Name string) {
	return this.Name
}
func (this *Coin) SetNameCn(NameCn string) (result *Coin) {
	this.NameCn = NameCn
	return this
}
func (this *Coin) GetNameCn() (NameCn string) {
	return this.NameCn
}
func (this *Coin) SetUnit(Unit string) (result *Coin) {
	this.Unit = Unit
	return this
}
func (this *Coin) GetUnit() (Unit string) {
	return this.Unit
}
func (this *Coin) SetStatus(Status CommonStatus.CommonStatus) (result *Coin) {
	this.Status = Status
	return this
}
func (this *Coin) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *Coin) SetMinTxFee(MinTxFee float64) (result *Coin) {
	this.MinTxFee = MinTxFee
	return this
}
func (this *Coin) GetMinTxFee() (MinTxFee float64) {
	return this.MinTxFee
}
func (this *Coin) SetCnyRate(CnyRate math.BigDecimal) (result *Coin) {
	this.CnyRate = CnyRate
	return this
}
func (this *Coin) GetCnyRate() (CnyRate math.BigDecimal) {
	return this.CnyRate
}
func (this *Coin) SetMaxTxFee(MaxTxFee float64) (result *Coin) {
	this.MaxTxFee = MaxTxFee
	return this
}
func (this *Coin) GetMaxTxFee() (MaxTxFee float64) {
	return this.MaxTxFee
}
func (this *Coin) SetUsdRate(UsdRate math.BigDecimal) (result *Coin) {
	this.UsdRate = UsdRate
	return this
}
func (this *Coin) GetUsdRate() (UsdRate math.BigDecimal) {
	return this.UsdRate
}
func (this *Coin) SetSgdRate(SgdRate math.BigDecimal) (result *Coin) {
	this.SgdRate = SgdRate
	return this
}
func (this *Coin) GetSgdRate() (SgdRate math.BigDecimal) {
	return this.SgdRate
}
func (this *Coin) SetEnableRpc(EnableRpc BooleanEnum.BooleanEnum) (result *Coin) {
	this.EnableRpc = EnableRpc
	return this
}
func (this *Coin) GetEnableRpc() (EnableRpc BooleanEnum.BooleanEnum) {
	return this.EnableRpc
}
func (this *Coin) SetSort(Sort int) (result *Coin) {
	this.Sort = Sort
	return this
}
func (this *Coin) GetSort() (Sort int) {
	return this.Sort
}
func (this *Coin) SetCanWithdraw(CanWithdraw BooleanEnum.BooleanEnum) (result *Coin) {
	this.CanWithdraw = CanWithdraw
	return this
}
func (this *Coin) GetCanWithdraw() (CanWithdraw BooleanEnum.BooleanEnum) {
	return this.CanWithdraw
}
func (this *Coin) SetCanRecharge(CanRecharge BooleanEnum.BooleanEnum) (result *Coin) {
	this.CanRecharge = CanRecharge
	return this
}
func (this *Coin) GetCanRecharge() (CanRecharge BooleanEnum.BooleanEnum) {
	return this.CanRecharge
}
func (this *Coin) SetCanTransfer(CanTransfer BooleanEnum.BooleanEnum) (result *Coin) {
	this.CanTransfer = CanTransfer
	return this
}
func (this *Coin) GetCanTransfer() (CanTransfer BooleanEnum.BooleanEnum) {
	return this.CanTransfer
}
func (this *Coin) SetCanAutoWithdraw(CanAutoWithdraw BooleanEnum.BooleanEnum) (result *Coin) {
	this.CanAutoWithdraw = CanAutoWithdraw
	return this
}
func (this *Coin) GetCanAutoWithdraw() (CanAutoWithdraw BooleanEnum.BooleanEnum) {
	return this.CanAutoWithdraw
}
func (this *Coin) SetWithdrawThreshold(WithdrawThreshold math.BigDecimal) (result *Coin) {
	this.WithdrawThreshold = WithdrawThreshold
	return this
}
func (this *Coin) GetWithdrawThreshold() (WithdrawThreshold math.BigDecimal) {
	return this.WithdrawThreshold
}
func (this *Coin) SetMinWithdrawAmount(MinWithdrawAmount math.BigDecimal) (result *Coin) {
	this.MinWithdrawAmount = MinWithdrawAmount
	return this
}
func (this *Coin) GetMinWithdrawAmount() (MinWithdrawAmount math.BigDecimal) {
	return this.MinWithdrawAmount
}
func (this *Coin) SetMaxWithdrawAmount(MaxWithdrawAmount math.BigDecimal) (result *Coin) {
	this.MaxWithdrawAmount = MaxWithdrawAmount
	return this
}
func (this *Coin) GetMaxWithdrawAmount() (MaxWithdrawAmount math.BigDecimal) {
	return this.MaxWithdrawAmount
}
func (this *Coin) SetIsPlatformCoin(IsPlatformCoin BooleanEnum.BooleanEnum) (result *Coin) {
	this.IsPlatformCoin = IsPlatformCoin
	return this
}
func (this *Coin) GetIsPlatformCoin() (IsPlatformCoin BooleanEnum.BooleanEnum) {
	return this.IsPlatformCoin
}
func (this *Coin) SetHasLegal(HasLegal bool) (result *Coin) {
	this.HasLegal = HasLegal
	return this
}
func (this *Coin) GetHasLegal() (HasLegal bool) {
	return this.HasLegal
}
func (this *Coin) SetAllBalance(AllBalance math.BigDecimal) (result *Coin) {
	this.AllBalance = AllBalance
	return this
}
func (this *Coin) GetAllBalance() (AllBalance math.BigDecimal) {
	return this.AllBalance
}
func (this *Coin) SetColdWalletAddress(ColdWalletAddress string) (result *Coin) {
	this.ColdWalletAddress = ColdWalletAddress
	return this
}
func (this *Coin) GetColdWalletAddress() (ColdWalletAddress string) {
	return this.ColdWalletAddress
}
func (this *Coin) SetHotAllBalance(HotAllBalance math.BigDecimal) (result *Coin) {
	this.HotAllBalance = HotAllBalance
	return this
}
func (this *Coin) GetHotAllBalance() (HotAllBalance math.BigDecimal) {
	return this.HotAllBalance
}
func (this *Coin) SetMinerFee(MinerFee math.BigDecimal) (result *Coin) {
	this.MinerFee = MinerFee
	return this
}
func (this *Coin) GetMinerFee() (MinerFee math.BigDecimal) {
	return this.MinerFee
}
func (this *Coin) SetWithdrawScale(WithdrawScale int) (result *Coin) {
	this.WithdrawScale = WithdrawScale
	return this
}
func (this *Coin) GetWithdrawScale() (WithdrawScale int) {
	return this.WithdrawScale
}
func (this *Coin) SetMinRechargeAmount(MinRechargeAmount math.BigDecimal) (result *Coin) {
	this.MinRechargeAmount = MinRechargeAmount
	return this
}
func (this *Coin) GetMinRechargeAmount() (MinRechargeAmount math.BigDecimal) {
	return this.MinRechargeAmount
}
func (this *Coin) SetMasterAddress(MasterAddress string) (result *Coin) {
	this.MasterAddress = MasterAddress
	return this
}
func (this *Coin) GetMasterAddress() (MasterAddress string) {
	return this.MasterAddress
}
func (this *Coin) SetMaxDailyWithdrawRate(MaxDailyWithdrawRate math.BigDecimal) (result *Coin) {
	this.MaxDailyWithdrawRate = MaxDailyWithdrawRate
	return this
}
func (this *Coin) GetMaxDailyWithdrawRate() (MaxDailyWithdrawRate math.BigDecimal) {
	return this.MaxDailyWithdrawRate
}
func NewCoin(name string, nameCn string, unit string, status CommonStatus.CommonStatus, minTxFee float64, cnyRate math.BigDecimal, maxTxFee float64, usdRate math.BigDecimal, sgdRate math.BigDecimal, enableRpc BooleanEnum.BooleanEnum, sort int, canWithdraw BooleanEnum.BooleanEnum, canRecharge BooleanEnum.BooleanEnum, canTransfer BooleanEnum.BooleanEnum, canAutoWithdraw BooleanEnum.BooleanEnum, withdrawThreshold math.BigDecimal, minWithdrawAmount math.BigDecimal, maxWithdrawAmount math.BigDecimal, isPlatformCoin BooleanEnum.BooleanEnum, hasLegal bool, allBalance math.BigDecimal, coldWalletAddress string, hotAllBalance math.BigDecimal, minerFee math.BigDecimal, withdrawScale int, minRechargeAmount math.BigDecimal, masterAddress string, maxDailyWithdrawRate math.BigDecimal) (ret *Coin) {
	ret = new(Coin)
	ret.Name = name
	ret.NameCn = nameCn
	ret.Unit = unit
	ret.Status = status
	ret.MinTxFee = minTxFee
	ret.CnyRate = cnyRate
	ret.MaxTxFee = maxTxFee
	ret.UsdRate = usdRate
	ret.SgdRate = sgdRate
	ret.EnableRpc = enableRpc
	ret.Sort = sort
	ret.CanWithdraw = canWithdraw
	ret.CanRecharge = canRecharge
	ret.CanTransfer = canTransfer
	ret.CanAutoWithdraw = canAutoWithdraw
	ret.WithdrawThreshold = withdrawThreshold
	ret.MinWithdrawAmount = minWithdrawAmount
	ret.MaxWithdrawAmount = maxWithdrawAmount
	ret.IsPlatformCoin = isPlatformCoin
	ret.HasLegal = hasLegal
	ret.AllBalance = allBalance
	ret.ColdWalletAddress = coldWalletAddress
	ret.HotAllBalance = hotAllBalance
	ret.MinerFee = minerFee
	ret.WithdrawScale = withdrawScale
	ret.MinRechargeAmount = minRechargeAmount
	ret.MasterAddress = masterAddress
	ret.MaxDailyWithdrawRate = maxDailyWithdrawRate
	return
}

type Coin struct {
	Name                 string                    `gorm:"column:name" json:"name"`
	NameCn               string                    `gorm:"column:name_cn" json:"nameCn"`
	Unit                 string                    `gorm:"column:unit" json:"unit"`
	Status               CommonStatus.CommonStatus `gorm:"column:status" json:"status"`
	MinTxFee             float64                   `gorm:"column:min_tx_fee" json:"minTxFee"`
	CnyRate              math.BigDecimal           `gorm:"column:cny_rate" json:"cnyRate"`
	MaxTxFee             float64                   `gorm:"column:max_tx_fee" json:"maxTxFee"`
	UsdRate              math.BigDecimal           `gorm:"column:usd_rate" json:"usdRate"`
	SgdRate              math.BigDecimal           `gorm:"column:sgd_rate" json:"sgdRate"`
	EnableRpc            BooleanEnum.BooleanEnum   `gorm:"column:enable_rpc" json:"enableRpc"`
	Sort                 int                       `gorm:"column:sort" json:"sort"`
	CanWithdraw          BooleanEnum.BooleanEnum   `gorm:"column:can_withdraw" json:"canWithdraw"`
	CanRecharge          BooleanEnum.BooleanEnum   `gorm:"column:can_recharge" json:"canRecharge"`
	CanTransfer          BooleanEnum.BooleanEnum   `gorm:"column:can_transfer" json:"canTransfer"`
	CanAutoWithdraw      BooleanEnum.BooleanEnum   `gorm:"column:can_auto_withdraw" json:"canAutoWithdraw"`
	WithdrawThreshold    math.BigDecimal           `gorm:"column:withdraw_threshold" json:"withdrawThreshold"`
	MinWithdrawAmount    math.BigDecimal           `gorm:"column:min_withdraw_amount" json:"minWithdrawAmount"`
	MaxWithdrawAmount    math.BigDecimal           `gorm:"column:max_withdraw_amount" json:"maxWithdrawAmount"`
	IsPlatformCoin       BooleanEnum.BooleanEnum   `gorm:"column:is_platform_coin" json:"isPlatformCoin"`
	HasLegal             bool                      `gorm:"column:has_legal" json:"hasLegal"`
	AllBalance           math.BigDecimal           `gorm:"column:all_balance" json:"allBalance"`
	ColdWalletAddress    string                    `gorm:"column:cold_wallet_address" json:"coldWalletAddress"`
	HotAllBalance        math.BigDecimal           `gorm:"column:hot_all_balance" json:"hotAllBalance"`
	MinerFee             math.BigDecimal           `gorm:"column:miner_fee" json:"minerFee"`
	WithdrawScale        int                       `gorm:"column:withdraw_scale" json:"withdrawScale"`
	MinRechargeAmount    math.BigDecimal           `gorm:"column:min_recharge_amount" json:"minRechargeAmount"`
	MasterAddress        string                    `gorm:"column:master_address" json:"masterAddress"`
	MaxDailyWithdrawRate math.BigDecimal           `gorm:"column:max_daily_withdraw_rate" json:"maxDailyWithdrawRate"`
}
