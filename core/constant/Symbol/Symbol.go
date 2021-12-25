package Symbol

type Symbol int

const (
	USDT Symbol = iota
	BTC
	ETH
	GCC
	GCX
)

func (this Symbol) String() string {
	switch this {
	}
	return ""
}
