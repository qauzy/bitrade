package constant

var (
	COMMANDS_VERSION              int = 1
	SUBSCRIBE_SYMBOL_THUMB        int = 20001
	UNSUBSCRIBE_SYMBOL_THUMB      int = 20002
	PUSH_SYMBOL_THUMB             int = 20003
	SUBSCRIBE_EXCHANGE            int = 20021
	UNSUBSCRIBE_EXCHANGE          int = 20022
	PUSH_EXCHANGE_TRADE           int = 20023
	PUSH_EXCHANGE_PLATE           int = 20024
	PUSH_EXCHANGE_KLINE           int = 20025
	PUSH_EXCHANGE_ORDER_COMPLETED int = 20026
	PUSH_EXCHANGE_ORDER_CANCELED  int = 20027
	PUSH_EXCHANGE_ORDER_TRADE     int = 20028
	PUSH_EXCHANGE_DEPTH           int = 20029
	SUBSCRIBE_CHAT                int = 20031
	UNSUBSCRIBE_CHAT              int = 20032
	PUSH_CHAT                     int = 20033
	SEND_CHAT                     int = 20034
	SUBSCRIBE_GROUP_CHAT          int = 20035
	UNSUBSCRIBE_GROUP_CHAT        int = 20036
	SUBSCRIBE_APNS                int = 20037
	UNSUBSCRIBE_APNS              int = 20038
	PUSH_GROUP_CHAT               int = 20039
)

type NettyCommand struct {
}
