package constant

var (
	SESSION_ADMIN                                  string = "ADMIN_MEMBER"
	SESSION_MEMBER                                 string = "API_MEMBER"
	PHONE_WITHDRAW_MONEY_CODE_PREFIX               string = "PHONE_WITHDRAW_MONEY_CODE_PREFIX_"
	PHONE_trade_CODE_PREFIX                        string = "PHONE_trade_CODE_PREFIX_"
	PHONE_REG_CODE_PREFIX                          string = "PHONE_REG_CODE_"
	PHONE_RESET_TRANS_CODE_PREFIX                  string = "PHONE_RESET_TRANS_CODE_"
	PHONE_BIND_CODE_PREFIX                         string = "PHONE_BIND_CODE_"
	PHONE_UPDATE_PASSWORD_PREFIX                   string = "PHONE_UPDATE_PASSWORD_"
	PHONE_ADD_ADDRESS_PREFIX                       string = "PHONE_ADD_ADDRESS_"
	EMAIL_BIND_CODE_PREFIX                         string = "EMAIL_BIND_CODE_"
	API_BIND_CODE_PREFIX                           string = "API_BIND_CODE_PREFIX_"
	EMAIL_UNTIE_CODE_PREFIX                        string = "EMAIL_UNTIE_CODE_"
	EMAIL_UPDATE_CODE_PREFIX                       string = "EMAIL_UPDATE_CODE_"
	ADD_ADDRESS_CODE_PREFIX                        string = "ADD_ADDRESS_CODE_"
	RESET_PASSWORD_CODE_PREFIX                     string = "RESET_PASSWORD_CODE_"
	PHONE_CHANGE_CODE_PREFIX                       string = "PHONE_CHANGE_CODE_"
	RESET_GOOGLE_CODE_PREFIX                       string = "RESET_PASSWORD_CODE_"
	ADMIN_LOGIN_PHONE_PREFIX                       string = "ADMIN_LOGIN_PHONE_PREFIX_"
	ADMIN_COIN_REVISE_PHONE_PREFIX                 string = "ADMIN_COIN_REVISE_PHONE_PREFIX_"
	ADMIN_COIN_TRANSFER_COLD_PREFIX                string = "ADMIN_COIN_TRANSFER_COLD_PREFIX_"
	ADMIN_EXCHANGE_COIN_SET_PREFIX                 string = "ADMIN_EXCHANGE_COIN_SET_PREFIX_"
	ANTI_ATTACK_                                   string = "ANTI_ATTACK_"
	ANTI_ROBOT_REGISTER                            string = "ANTI_ROBOT_REGISTER_"
	BHB_AMOUNT                                     string = "BHB_AMOUNT"
	BHB_AMOUNT_EXPIRE_TIME                         int    = 900
	NOTICE_DETAIL                                  string = "notice_detail_"
	NOTICE_DETAIL_EXPIRE_TIME                      int    = 300
	SYS_HELP                                       string = "SYS_HELP"
	SYS_HELP_EXPIRE_TIME                           int    = 300
	SYS_HELP_CATE                                  string = "SYS_HELP_CATE_"
	SYS_HELP_CATE_EXPIRE_TIME                      int    = 300
	SYS_HELP_DETAIL                                string = "SYS_HELP_DETAIL_"
	SYS_HELP_DETAIL_EXPIRE_TIME                    int    = 300
	SYS_HELP_TOP                                   string = "SYS_HELP_TOP_"
	SYS_HELP_TOP_EXPIRE_TIME                       int    = 300
	DATA_DICTIONARY_BOUND_KEY                      string = "data_dictionary_bound_key_"
	DATA_DICTIONARY_BOUND_EXPIRE_TIME              int    = 604800
	EXCHANGE_INIT_PLATE_SYMBOL_KEY                 string = "EXCHANGE_INIT_PLATE_SYMBOL_KEY_"
	EXCHANGE_INIT_PLATE_SYMBOL_EXPIRE_TIME         int    = 18000
	EXCHANGE_INIT_PLATE_ALL_SYMBOLS                string = "EXCHANGE_INIT_PLATE_ALL_SYMBOLS"
	USER_ADD_EXCHANGE_ORDER_TIME_LIMIT             string = "USER_ADD_EXCHANGE_ORDER_TIME_LIMIT_"
	USER_ADD_EXCHANGE_ORDER_TIME_LIMIT_EXPIRE_TIME int    = 20
	HANDLE_AIRDROP_LOCK                            string = "HANDLE_AIRDROP_LOCK_"
	LOGIN_LOCK                                     string = "LOGIN_LOCK_"
	INTEGRATION_GIVING_ONE_INVITE                  string = "integration_giving_one_invite"
	INTEGRATION_GIVING_TWO_INVITE                  string = "integration_giving_two_invite"
	INTEGRATION_GIVING_OTC_BUY_CNY_RATE            string = "integration_giving_otc_buy_cny_rate"
	INTEGRATION_GIVING_EXCHANGE_RECHARGE_USDT_RATE string = "integration_giving_exchange_recharge_usdt_rate"
	CUSTOMER_DAY_WITHDRAW_TOTAL_COUNT              string = "CUSTOMER_DAY_WITHDRAW_TOTAL_COUNT_"
	CUSTOMER_DAY_WITHDRAW_COVER_USD_AMOUNT         string = "CUSTOMER_DAY_WITHDRAW_COVER_USD_AMOUNT_"
	CUSTOMER_INTEGRATION_GRADE                     string = "CUSTOMER_INTEGRATION_GRADE_"
	DEFAULT_SYMBOL                                 string = "DEFAULT_SYMBOL"
)

type SysConstant struct {
}
