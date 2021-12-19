package vo

type TransferType int

const (
	TRANSFERTYPE_apply TransferType = iota
	TRANSFERTYPE_success
	TRANSFERTYPE_failed
)

func (this TransferType) String() string {
	switch this {
	}
	return ""
}
