package constant

type Symbol int

const (
	SYMBOL_USDT Symbol = iota
	SYMBOL_BTC
	SYMBOL_ETH
	SYMBOL_GCC
	SYMBOL_GCX
)

func (this Symbol) String() string {
	switch this {
	}
	return ""
}
