package TransactionType

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	RECHARGE             = TransactionType{"充值", 1}
	WITHDRAW             = TransactionType{"提现", 2}
	TRANSFER_ACCOUNTS    = TransactionType{"转账", 3}
	EXCHANGE             = TransactionType{"币币交易", 4}
	OTC_BUY              = TransactionType{"法币买入", 5}
	OTC_SELL             = TransactionType{"法币卖出", 6}
	ACTIVITY_AWARD       = TransactionType{"活动奖励", 7}
	PROMOTION_AWARD      = TransactionType{"推广奖励", 8}
	DIVIDEND             = TransactionType{"分红", 9}
	VOTE                 = TransactionType{"投票", 10}
	ADMIN_RECHARGE       = TransactionType{"人工充值", 11}
	MATCH                = TransactionType{"配对", 12}
	DEPOSIT              = TransactionType{"缴纳商家认证保证金", 13}
	GET_BACK_DEPOSIT     = TransactionType{"退回商家认证保证金", 14}
	LEGAL_RECHARGE       = TransactionType{"法币充值", 15}
	ASSET_EXCHANGE       = TransactionType{"币币兑换", 16}
	CHANNEL_AWARD        = TransactionType{"渠道推广", 17}
	TRANSFER_INTO_LEVER  = TransactionType{"划转入杠杆钱包", 18}
	TRANSFER_OUT_LEVER   = TransactionType{"从杠杆钱包划转出", 19}
	MANUAL_AIRDROP       = TransactionType{"钱包空投", 20}
	LOCK_POSITION        = TransactionType{"锁仓", 21}
	UNLOCK_POSITION      = TransactionType{"解锁", 22}
	THIRD_PARTY_TRANSFER = TransactionType{"第三方转入", 23}
	THIRD_PARTY_TURN_OUT = TransactionType{"第三方转出", 24}
	COIN_TWO_OTC         = TransactionType{"币币转入法币", 25}
	OTC_TWO_COIN         = TransactionType{"法币转入币币", 26}
	LOAN_RECORD          = TransactionType{"借贷流水", 27}
	REPAYMENT_LOAN       = TransactionType{"还款流水", 28}
)

func (this *TransactionType) Ordinal() (result int) {
	return this.ordinal
}
func (this *TransactionType) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *TransactionType) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "充值"
		case 2:
			this.CnName = "提现"
		case 3:
			this.CnName = "转账"
		case 4:
			this.CnName = "币币交易"
		case 5:
			this.CnName = "法币买入"
		case 6:
			this.CnName = "法币卖出"
		case 7:
			this.CnName = "活动奖励"
		case 8:
			this.CnName = "推广奖励"
		case 9:
			this.CnName = "分红"
		case 10:
			this.CnName = "投票"
		case 11:
			this.CnName = "人工充值"
		case 12:
			this.CnName = "配对"
		case 13:
			this.CnName = "缴纳商家认证保证金"
		case 14:
			this.CnName = "退回商家认证保证金"
		case 15:
			this.CnName = "法币充值"
		case 16:
			this.CnName = "币币兑换"
		case 17:
			this.CnName = "渠道推广"
		case 18:
			this.CnName = "划转入杠杆钱包"
		case 19:
			this.CnName = "从杠杆钱包划转出"
		case 20:
			this.CnName = "钱包空投"
		case 21:
			this.CnName = "锁仓"
		case 22:
			this.CnName = "解锁"
		case 23:
			this.CnName = "第三方转入"
		case 24:
			this.CnName = "第三方转出"
		case 25:
			this.CnName = "币币转入法币"
		case 26:
			this.CnName = "法币转入币币"
		case 27:
			this.CnName = "借贷流水"
		case 28:
			this.CnName = "还款流水"
		}
	default:
		this = nil
	}
	return nil
}
func (this *TransactionType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *TransactionType) UnmarshalJSON(data []byte) (err error) {
	if data == nil || len(data) == 2 {
		return
	}
	this.ordinal, err = strconv.Atoi(string(data))
	if err != nil {
		return
	}
	switch this.ordinal {
	}
	switch this.ordinal {
	case 1:
		this.CnName = "充值"
	case 2:
		this.CnName = "提现"
	case 3:
		this.CnName = "转账"
	case 4:
		this.CnName = "币币交易"
	case 5:
		this.CnName = "法币买入"
	case 6:
		this.CnName = "法币卖出"
	case 7:
		this.CnName = "活动奖励"
	case 8:
		this.CnName = "推广奖励"
	case 9:
		this.CnName = "分红"
	case 10:
		this.CnName = "投票"
	case 11:
		this.CnName = "人工充值"
	case 12:
		this.CnName = "配对"
	case 13:
		this.CnName = "缴纳商家认证保证金"
	case 14:
		this.CnName = "退回商家认证保证金"
	case 15:
		this.CnName = "法币充值"
	case 16:
		this.CnName = "币币兑换"
	case 17:
		this.CnName = "渠道推广"
	case 18:
		this.CnName = "划转入杠杆钱包"
	case 19:
		this.CnName = "从杠杆钱包划转出"
	case 20:
		this.CnName = "钱包空投"
	case 21:
		this.CnName = "锁仓"
	case 22:
		this.CnName = "解锁"
	case 23:
		this.CnName = "第三方转入"
	case 24:
		this.CnName = "第三方转出"
	case 25:
		this.CnName = "币币转入法币"
	case 26:
		this.CnName = "法币转入币币"
	case 27:
		this.CnName = "借贷流水"
	case 28:
		this.CnName = "还款流水"
	}
	return
}
func (this *TransactionType) GetOrdinal() (result int) {
	return this.Ordinal()
}
func ValueOfOrdinal(ordinal int) (result TransactionType) {
	switch ordinal {
	case 0:
		return RECHARGE
	case 1:
		return WITHDRAW
	case 2:
		return TRANSFER_ACCOUNTS
	case 3:
		return EXCHANGE
	case 4:
		return OTC_BUY
	case 5:
		return OTC_SELL
	case 6:
		return ACTIVITY_AWARD
	case 7:
		return PROMOTION_AWARD
	case 8:
		return DIVIDEND
	case 9:
		return VOTE
	case 10:
		return ADMIN_RECHARGE
	}
	return
}

func (this *TransactionType) Equals(ordinal TransactionType) (result bool) {
	return *this == ordinal
}
func ParseOrdinal(ordinal TransactionType) (result int) {
	if RECHARGE.Equals(ordinal) {
		return 0
	} else if WITHDRAW.Equals(ordinal) {
		return 1
	} else if TRANSFER_ACCOUNTS.Equals(ordinal) {
		return 2
	} else if EXCHANGE.Equals(ordinal) {
		return 3
	} else if OTC_BUY.Equals(ordinal) {
		return 4
	} else if OTC_SELL.Equals(ordinal) {
		return 5
	} else if ACTIVITY_AWARD.Equals(ordinal) {
		return 6
	} else if PROMOTION_AWARD.Equals(ordinal) {
		return 7
	} else if DIVIDEND.Equals(ordinal) {
		return 8
	} else if VOTE.Equals(ordinal) {
		return 9
	} else if ADMIN_RECHARGE.Equals(ordinal) {
		return 10
	} else {
		return 20
	}
}

type TransactionType struct {
	CnName  string
	ordinal int
}
