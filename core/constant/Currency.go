package constant

type Currency int

const CURRENCY_BTC Currency = iota

func (this Currency) String() string {
	switch this {
	case CURRENCY_BTC:
		return "比特币"
	}
	return ""
}
