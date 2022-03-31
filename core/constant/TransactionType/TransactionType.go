package TransactionType

var (
	RECHARGE             = TransactionType{"充值", 0}
	WITHDRAW             = TransactionType{"提现", 1}
	TRANSFER_ACCOUNTS    = TransactionType{"转账", 2}
	EXCHANGE             = TransactionType{"币币交易", 3}
	OTC_BUY              = TransactionType{"法币买入", 4}
	OTC_SELL             = TransactionType{"法币卖出", 5}
	ACTIVITY_AWARD       = TransactionType{"活动奖励", 6}
	PROMOTION_AWARD      = TransactionType{"推广奖励", 7}
	DIVIDEND             = TransactionType{"分红", 8}
	VOTE                 = TransactionType{"投票", 9}
	ADMIN_RECHARGE       = TransactionType{"人工充值", 10}
	MATCH                = TransactionType{"配对", 11}
	DEPOSIT              = TransactionType{"缴纳商家认证保证金", 12}
	GET_BACK_DEPOSIT     = TransactionType{"退回商家认证保证金", 13}
	LEGAL_RECHARGE       = TransactionType{"法币充值", 14}
	ASSET_EXCHANGE       = TransactionType{"币币兑换", 15}
	CHANNEL_AWARD        = TransactionType{"渠道推广", 16}
	TRANSFER_INTO_LEVER  = TransactionType{"划转入杠杆钱包", 17}
	TRANSFER_OUT_LEVER   = TransactionType{"从杠杆钱包划转出", 18}
	MANUAL_AIRDROP       = TransactionType{"钱包空投", 19}
	LOCK_POSITION        = TransactionType{"锁仓", 20}
	UNLOCK_POSITION      = TransactionType{"解锁", 21}
	THIRD_PARTY_TRANSFER = TransactionType{"第三方转入", 22}
	THIRD_PARTY_TURN_OUT = TransactionType{"第三方转出", 23}
	COIN_TWO_OTC         = TransactionType{"币币转入法币", 24}
	OTC_TWO_COIN         = TransactionType{"法币转入币币", 25}
	LOAN_RECORD          = TransactionType{"借贷流水", 26}
	REPAYMENT_LOAN       = TransactionType{"还款流水", 27}
)

func (this *TransactionType) Ordinal() (result int) {
	return this.ordinal
}
func (this *TransactionType) GetOrdinal() (result int) {
	return this.Ordinal()
}
func ValueOfOrdinal(ordinal int) (result TransactionType) {
	switch ordinal {
	case 0:
		return this.RECHARGE
	case 1:
		return this.WITHDRAW
	case 2:
		return this.TRANSFER_ACCOUNTS
	case 3:
		return this.EXCHANGE
	case 4:
		return this.OTC_BUY
	case 5:
		return this.OTC_SELL
	case 6:
		return this.ACTIVITY_AWARD
	case 7:
		return this.PROMOTION_AWARD
	case 8:
		return this.DIVIDEND
	case 9:
		return this.VOTE
	case 10:
		return this.ADMIN_RECHARGE
	}
	return nil
}
func ParseOrdinal(ordinal TransactionType) (result int) {
	if TransactionType.RECHARGE.Equals(ordinal) {
		return 0
	} else if TransactionType.WITHDRAW.Equals(ordinal) {
		return 1
	} else if TransactionType.TRANSFER_ACCOUNTS.Equals(ordinal) {
		return 2
	} else if TransactionType.EXCHANGE.Equals(ordinal) {
		return 3
	} else if TransactionType.OTC_BUY.Equals(ordinal) {
		return 4
	} else if TransactionType.OTC_SELL.Equals(ordinal) {
		return 5
	} else if TransactionType.ACTIVITY_AWARD.Equals(ordinal) {
		return 6
	} else if TransactionType.PROMOTION_AWARD.Equals(ordinal) {
		return 7
	} else if TransactionType.DIVIDEND.Equals(ordinal) {
		return 8
	} else if TransactionType.VOTE.Equals(ordinal) {
		return 9
	} else if TransactionType.ADMIN_RECHARGE.Equals(ordinal) {
		return 10
	} else {
		return 20
	}
}

type TransactionType struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
