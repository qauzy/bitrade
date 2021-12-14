package transform

func (this *Special) SetContext(context []E) {
	this.Context = context
}
func (this *Special) GetContext() (context []E) {
	return this.Context
}

type Special struct {
	Context []E
}
