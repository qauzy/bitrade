
package SignStatus

var (
	UNDERWAY SignStatus = iota
	FINISH
)

func (this *SignStatus) Ordinal() (result int) {
	return this.ordinal
}

type SignStatus struct {
	ordinal int
}

