package TransactionType

type TransactionType int

const (
	RECHARGE TransactionType = iota
	WITHDRAW
	TRANSFER_ACCOUNTS
	EXCHANGE
	OTC_BUY
	OTC_SELL
	ACTIVITY_AWARD
	PROMOTION_AWARD
	DIVIDEND
	VOTE
	ADMIN_RECHARGE
	MATCH
	DEPOSIT
	GET_BACK_DEPOSIT
	LEGAL_RECHARGE
	ASSET_EXCHANGE
	CHANNEL_AWARD
	TRANSFER_INTO_LEVER
	TRANSFER_OUT_LEVER
	MANUAL_AIRDROP
	LOCK_POSITION
	UNLOCK_POSITION
	THIRD_PARTY_TRANSFER
	THIRD_PARTY_TURN_OUT
	COIN_TWO_OTC
	OTC_TWO_COIN
	LOAN_RECORD
	REPAYMENT_LOAN
)

func (this TransactionType) String() string {
	switch this {
	case RECHARGE:
		return "充值"
	case WITHDRAW:
		return "提现"
	case TRANSFER_ACCOUNTS:
		return "转账"
	case EXCHANGE:
		return "币币交易"
	case OTC_BUY:
		return "法币买入"
	case OTC_SELL:
		return "法币卖出"
	case ACTIVITY_AWARD:
		return "活动奖励"
	case PROMOTION_AWARD:
		return "推广奖励"
	case DIVIDEND:
		return "分红"
	case VOTE:
		return "投票"
	case ADMIN_RECHARGE:
		return "人工充值"
	case MATCH:
		return "配对"
	case DEPOSIT:
		return "缴纳商家认证保证金"
	case GET_BACK_DEPOSIT:
		return "退回商家认证保证金"
	case LEGAL_RECHARGE:
		return "法币充值"
	case ASSET_EXCHANGE:
		return "币币兑换"
	case CHANNEL_AWARD:
		return "渠道推广"
	case TRANSFER_INTO_LEVER:
		return "划转入杠杆钱包"
	case TRANSFER_OUT_LEVER:
		return "从杠杆钱包划转出"
	case MANUAL_AIRDROP:
		return "钱包空投"
	case LOCK_POSITION:
		return "锁仓"
	case UNLOCK_POSITION:
		return "解锁"
	case THIRD_PARTY_TRANSFER:
		return "第三方转入"
	case THIRD_PARTY_TURN_OUT:
		return "第三方转出"
	case COIN_TWO_OTC:
		return "币币转入法币"
	case OTC_TWO_COIN:
		return "法币转入币币"
	case LOAN_RECORD:
		return "借贷流水"
	case REPAYMENT_LOAN:
		return "还款流水"
	}
	return ""
}
