package constant

type SignStatus int

const (
	SIGNSTATUS_UNDERWAY SignStatus = iota
	SIGNSTATUS_FINISH
)

func (this SignStatus) String() string {
	switch this {
	}
	return ""
}
