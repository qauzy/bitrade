package SignStatus

type SignStatus int

const (
	UNDERWAY SignStatus = iota
	FINISH
)

func (this SignStatus) String() string {
	switch this {
	}
	return ""
}
