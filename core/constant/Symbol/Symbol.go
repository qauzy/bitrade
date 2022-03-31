
package Symbol

var (
	USDT Symbol = iota
	BTC
	ETH
	GCC
	GCX
)

func (this *Symbol) Ordinal() (result int) {
	return this.ordinal
}

type Symbol struct {
	ordinal int
}

