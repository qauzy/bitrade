package Currency

type Currency int

const BTC Currency = iota

func (this Currency) String() string {
	switch this {
	case BTC:
		return "比特币"
	}
	return ""
}
