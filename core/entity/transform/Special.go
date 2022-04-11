package transform

import "github.com/qauzy/chocolate/lists/arraylist"

func (this *Special[E]) SetContext(Context arraylist.List[E]) (result *Special[E]) {
	this.Context = Context
	return this
}
func (this *Special[E]) GetContext() (Context arraylist.List[E]) {
	return this.Context
}
func NewSpecial[E any](context arraylist.List[E]) (ret *Special[E]) {
	ret = new(Special[E])
	ret.Context = context
	return
}

type Special[E any] struct {
	Context arraylist.List[E] `gorm:"column:context" json:"context"`
}
